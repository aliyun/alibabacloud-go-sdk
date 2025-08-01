// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaySlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewaySlbRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *ListGatewaySlbRequest
	GetGatewayUniqueId() *string
}

type ListGatewaySlbRequest struct {
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
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-1cef5440bf2d484db419fb264d4f****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s ListGatewaySlbRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaySlbRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaySlbRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewaySlbRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewaySlbRequest) SetAcceptLanguage(v string) *ListGatewaySlbRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewaySlbRequest) SetGatewayUniqueId(v string) *ListGatewaySlbRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewaySlbRequest) Validate() error {
	return dara.Validate(s)
}
