// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBusinessLocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryBusinessLocationsRequest
	GetAcceptLanguage() *string
}

type QueryBusinessLocationsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s QueryBusinessLocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBusinessLocationsRequest) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryBusinessLocationsRequest) SetAcceptLanguage(v string) *QueryBusinessLocationsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryBusinessLocationsRequest) Validate() error {
	return dara.Validate(s)
}
