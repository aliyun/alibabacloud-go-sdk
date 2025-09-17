// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetResourceTypeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAcceptLanguage(v string) *GetResourceTypeHeaders
	GetXAcsAcceptLanguage() *string
}

type GetResourceTypeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The language selected for the returned product.
	//
	// zh_CH: Chinese (default)
	//
	// en_US: English
	//
	// example:
	//
	// zh_CH
	XAcsAcceptLanguage *string `json:"x-acs-accept-language,omitempty" xml:"x-acs-accept-language,omitempty"`
}

func (s GetResourceTypeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeHeaders) GoString() string {
	return s.String()
}

func (s *GetResourceTypeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetResourceTypeHeaders) GetXAcsAcceptLanguage() *string {
	return s.XAcsAcceptLanguage
}

func (s *GetResourceTypeHeaders) SetCommonHeaders(v map[string]*string) *GetResourceTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResourceTypeHeaders) SetXAcsAcceptLanguage(v string) *GetResourceTypeHeaders {
	s.XAcsAcceptLanguage = &v
	return s
}

func (s *GetResourceTypeHeaders) Validate() error {
	return dara.Validate(s)
}
