// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGatewayRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryGatewayRegionRequest
	GetAcceptLanguage() *string
}

type QueryGatewayRegionRequest struct {
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

func (s QueryGatewayRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGatewayRegionRequest) GoString() string {
	return s.String()
}

func (s *QueryGatewayRegionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryGatewayRegionRequest) SetAcceptLanguage(v string) *QueryGatewayRegionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryGatewayRegionRequest) Validate() error {
	return dara.Validate(s)
}
