package charsetx

import (
	"bytes"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

var (
	// Alias for charsets.
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

// Supported returns whether charset `charset` is supported.
func Supported(charset string) bool {
	return getEncoding(charset) != nil
}

// Convert converts `src` charset encoding from `srcCharset` to `dstCharset`,
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	// Converting `src` to UTF-8.
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if err != nil {
				return "", err
			}
			src = string(tmp)
		} else {
			return dst, fmt.Errorf(`unsupported srcCharset "%s"`, srcCharset)
		}
	}
	// Do the converting from UTF-8 to `dstCharset`.
	if dstCharset != "UTF-8" {
		if e := getEncoding(dstCharset); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if err != nil {
				return "", err
			}
			dst = string(tmp)
		} else {
			return dst, err
		}
	} else {
		dst = src
	}
	return dst, nil
}

// ToUTF8 converts `src` charset encoding from `srcCharset` to UTF-8 ,
// and returns the converted string.
func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

// UTF8To converts `src` charset encoding from UTF-8 to `dstCharset`,
// and returns the converted string.
func UTF8To(dstCharset string, src string) (dst string, err error) {
	return Convert(dstCharset, "UTF-8", src)
}

// getEncoding returns the encoding.Encoding interface object for `charset`.
// It returns nil if `charset` is not supported.
func getEncoding(charset string) encoding.Encoding {
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}
	enc, err := ianaindex.MIB.Encoding(charset)
	if err != nil {
		logger.Error(err)
	}
	return enc
}
