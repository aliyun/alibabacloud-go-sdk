// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListProductsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAcceptLanguage(v string) *ListProductsHeaders
	GetXAcsAcceptLanguage() *string
}

type ListProductsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// Select the language in which the response is returned.
	//
	// zh_CH: Chinese (default).
	//
	// en_US: English.
	//
	// example:
	//
	// zh_CH
	XAcsAcceptLanguage *string `json:"x-acs-accept-language,omitempty" xml:"x-acs-accept-language,omitempty"`
}

func (s ListProductsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListProductsHeaders) GoString() string {
	return s.String()
}

func (s *ListProductsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListProductsHeaders) GetXAcsAcceptLanguage() *string {
	return s.XAcsAcceptLanguage
}

func (s *ListProductsHeaders) SetCommonHeaders(v map[string]*string) *ListProductsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListProductsHeaders) SetXAcsAcceptLanguage(v string) *ListProductsHeaders {
	s.XAcsAcceptLanguage = &v
	return s
}

func (s *ListProductsHeaders) Validate() error {
	return dara.Validate(s)
}
