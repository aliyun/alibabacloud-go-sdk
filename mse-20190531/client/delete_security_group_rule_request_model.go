// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteSecurityGroupRuleRequest
	GetAcceptLanguage() *string
	SetCascadingDelete(v bool) *DeleteSecurityGroupRuleRequest
	GetCascadingDelete() *bool
	SetGatewayUniqueId(v string) *DeleteSecurityGroupRuleRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *DeleteSecurityGroupRuleRequest
	GetId() *int64
}

type DeleteSecurityGroupRuleRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage  *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	CascadingDelete *bool   `json:"CascadingDelete,omitempty" xml:"CascadingDelete,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-83b0ddb569434f82b9fe8e4c60c40f7c
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The destination ID.
	//
	// example:
	//
	// 93
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteSecurityGroupRuleRequest) GetCascadingDelete() *bool {
	return s.CascadingDelete
}

func (s *DeleteSecurityGroupRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteSecurityGroupRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSecurityGroupRuleRequest) SetAcceptLanguage(v string) *DeleteSecurityGroupRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteSecurityGroupRuleRequest) SetCascadingDelete(v bool) *DeleteSecurityGroupRuleRequest {
	s.CascadingDelete = &v
	return s
}

func (s *DeleteSecurityGroupRuleRequest) SetGatewayUniqueId(v string) *DeleteSecurityGroupRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteSecurityGroupRuleRequest) SetId(v int64) *DeleteSecurityGroupRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
