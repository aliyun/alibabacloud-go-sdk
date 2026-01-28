// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceId(v string) *DeleteAuthorizationResourceRequest
	GetAuthorizationResourceId() *string
	SetAuthorizationRuleId(v string) *DeleteAuthorizationResourceRequest
	GetAuthorizationRuleId() *string
	SetInstanceId(v string) *DeleteAuthorizationResourceRequest
	GetInstanceId() *string
}

type DeleteAuthorizationResourceRequest struct {
	// 授权资源标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arres_01kgh3jvt7pk093rv6giu0c0qxxxx
	AuthorizationResourceId *string `json:"AuthorizationResourceId,omitempty" xml:"AuthorizationResourceId,omitempty"`
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteAuthorizationResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationResourceRequest) GetAuthorizationResourceId() *string {
	return s.AuthorizationResourceId
}

func (s *DeleteAuthorizationResourceRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *DeleteAuthorizationResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAuthorizationResourceRequest) SetAuthorizationResourceId(v string) *DeleteAuthorizationResourceRequest {
	s.AuthorizationResourceId = &v
	return s
}

func (s *DeleteAuthorizationResourceRequest) SetAuthorizationRuleId(v string) *DeleteAuthorizationResourceRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *DeleteAuthorizationResourceRequest) SetInstanceId(v string) *DeleteAuthorizationResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAuthorizationResourceRequest) Validate() error {
	return dara.Validate(s)
}
