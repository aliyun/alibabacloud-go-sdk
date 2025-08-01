// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGatewayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryGatewayTypeRequest
	GetAcceptLanguage() *string
}

type QueryGatewayTypeRequest struct {
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

func (s QueryGatewayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGatewayTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryGatewayTypeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryGatewayTypeRequest) SetAcceptLanguage(v string) *QueryGatewayTypeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryGatewayTypeRequest) Validate() error {
	return dara.Validate(s)
}
