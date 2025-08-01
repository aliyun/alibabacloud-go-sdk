// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayZoneRequest
	GetAcceptLanguage() *string
}

type ListGatewayZoneRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListGatewayZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayZoneRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayZoneRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayZoneRequest) SetAcceptLanguage(v string) *ListGatewayZoneRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayZoneRequest) Validate() error {
	return dara.Validate(s)
}
