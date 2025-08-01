// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListSecurityGroupRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *ListSecurityGroupRequest
	GetGatewayUniqueId() *string
}

type ListSecurityGroupRequest struct {
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
	// The ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-e98e40675aaf49bda082137d158e1585
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s ListSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListSecurityGroupRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListSecurityGroupRequest) SetAcceptLanguage(v string) *ListSecurityGroupRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListSecurityGroupRequest) SetGatewayUniqueId(v string) *ListSecurityGroupRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
