// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListResourceTypesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAcceptLanguage(v string) *ListResourceTypesHeaders
	GetXAcsAcceptLanguage() *string
}

type ListResourceTypesHeaders struct {
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

func (s ListResourceTypesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesHeaders) GoString() string {
	return s.String()
}

func (s *ListResourceTypesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListResourceTypesHeaders) GetXAcsAcceptLanguage() *string {
	return s.XAcsAcceptLanguage
}

func (s *ListResourceTypesHeaders) SetCommonHeaders(v map[string]*string) *ListResourceTypesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResourceTypesHeaders) SetXAcsAcceptLanguage(v string) *ListResourceTypesHeaders {
	s.XAcsAcceptLanguage = &v
	return s
}

func (s *ListResourceTypesHeaders) Validate() error {
	return dara.Validate(s)
}
