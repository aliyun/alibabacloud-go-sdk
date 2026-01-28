// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceId(v string) *GetAuthorizationResourceRequest
	GetAuthorizationResourceId() *string
	SetAuthorizationRuleId(v string) *GetAuthorizationResourceRequest
	GetAuthorizationRuleId() *string
	SetInstanceId(v string) *GetAuthorizationResourceRequest
	GetInstanceId() *string
}

type GetAuthorizationResourceRequest struct {
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

func (s GetAuthorizationResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationResourceRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResourceRequest) GetAuthorizationResourceId() *string {
	return s.AuthorizationResourceId
}

func (s *GetAuthorizationResourceRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *GetAuthorizationResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationResourceRequest) SetAuthorizationResourceId(v string) *GetAuthorizationResourceRequest {
	s.AuthorizationResourceId = &v
	return s
}

func (s *GetAuthorizationResourceRequest) SetAuthorizationRuleId(v string) *GetAuthorizationResourceRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *GetAuthorizationResourceRequest) SetInstanceId(v string) *GetAuthorizationResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationResourceRequest) Validate() error {
	return dara.Validate(s)
}
