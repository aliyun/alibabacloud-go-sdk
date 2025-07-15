// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainWebSocketStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionValue(v string) *SetDomainWebSocketStatusRequest
	GetActionValue() *string
	SetDomainName(v string) *SetDomainWebSocketStatusRequest
	GetDomainName() *string
	SetGroupId(v string) *SetDomainWebSocketStatusRequest
	GetGroupId() *string
	SetSecurityToken(v string) *SetDomainWebSocketStatusRequest
	GetSecurityToken() *string
	SetWSSEnable(v string) *SetDomainWebSocketStatusRequest
	GetWSSEnable() *string
}

type SetDomainWebSocketStatusRequest struct {
	// The action.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	ActionValue *string `json:"ActionValue,omitempty" xml:"ActionValue,omitempty"`
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac.fluvet.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cf976e63b70c4993807e7bb9345d4695
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// If enable WSS.
	//
	// example:
	//
	// false
	WSSEnable *string `json:"WSSEnable,omitempty" xml:"WSSEnable,omitempty"`
}

func (s SetDomainWebSocketStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDomainWebSocketStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDomainWebSocketStatusRequest) GetActionValue() *string {
	return s.ActionValue
}

func (s *SetDomainWebSocketStatusRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDomainWebSocketStatusRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetDomainWebSocketStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetDomainWebSocketStatusRequest) GetWSSEnable() *string {
	return s.WSSEnable
}

func (s *SetDomainWebSocketStatusRequest) SetActionValue(v string) *SetDomainWebSocketStatusRequest {
	s.ActionValue = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetDomainName(v string) *SetDomainWebSocketStatusRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetGroupId(v string) *SetDomainWebSocketStatusRequest {
	s.GroupId = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetSecurityToken(v string) *SetDomainWebSocketStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetWSSEnable(v string) *SetDomainWebSocketStatusRequest {
	s.WSSEnable = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) Validate() error {
	return dara.Validate(s)
}
