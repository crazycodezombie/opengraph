package main

import (
	"context"
	"fmt"
	"github.com/otiai10/opengraph/v2"
	"github.com/otiai10/opengraph/v2/http_fetchers"
	"net/url"
	"strings"
	"testing"
	"time"
)

func BaseURL(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	return strings.ToLower(result), nil
}

func extractOgpData(url string, pageLoad bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	intent := opengraph.Intent{Context: ctx}
	if pageLoad {
		intent.HTTPFetcher = http_fetchers.NewPageLoadHTTPFetcher(10)
	}

	ogp, err := opengraph.Fetch(url, intent)
	if err != nil {
		return err
	}

	if ogp.URL == "" {
		if ogp.Title == "" || ogp.Description == "" {
			return nil
		}

		ogp.URL, err = BaseURL(url)
		if err != nil {
			return nil
		}
	}
	//
	//result := OgpData{
	//	URL: strings.TrimSpace(ogp.URL),
	//}
	//
	//if ogp.Title != "" {
	//	result.Title = proto.String(strings.TrimSpace(ogp.Title))
	//}
	//
	//if ogp.Description != "" {
	//	result.Description = proto.String(strings.TrimSpace(ogp.Description))
	//}
	//
	//if len(ogp.Image) > 0 && ogp.Image[0].URL != "" {
	//	if err = driver.enrichOgpDataWithImageInfo(&result, ogp.Image[0].URL); err != nil {
	//		logrus.Warnf("unable to fetch ogp image info for url %v: %v. contimue without image", url, err)
	//	}
	//} else {
	//	logrus.Infof("no ogp image for url %v", url)
	//}

	return nil
}

func TestOgp(t *testing.T) {
	if err := extractOgpData("http://www.balyan.co.il", true); err != nil {
		panic("nil")
	}
}
