// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v string) *CreateMonitorGroupRequest
	GetAuth() *string
	SetGroupId(v string) *CreateMonitorGroupRequest
	GetGroupId() *string
	SetRawMonitorGroupId(v int64) *CreateMonitorGroupRequest
	GetRawMonitorGroupId() *int64
	SetSecurityToken(v string) *CreateMonitorGroupRequest
	GetSecurityToken() *string
}

type CreateMonitorGroupRequest struct {
	// The caller authentication status of the API. Valid values: **ok**: The authentication is successful. **mismatch**: The request is redirected. **servicenotfound**: A request error occurred. **Unknown**: An unknown error occurred.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6735211ab9094c818f32f27bc545b6c8
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the monitoring group.
	//
	// example:
	//
	// 166636221
	RawMonitorGroupId *int64  `json:"RawMonitorGroupId,omitempty" xml:"RawMonitorGroupId,omitempty"`
	SecurityToken     *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateMonitorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupRequest) GetAuth() *string {
	return s.Auth
}

func (s *CreateMonitorGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateMonitorGroupRequest) GetRawMonitorGroupId() *int64 {
	return s.RawMonitorGroupId
}

func (s *CreateMonitorGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateMonitorGroupRequest) SetAuth(v string) *CreateMonitorGroupRequest {
	s.Auth = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetGroupId(v string) *CreateMonitorGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetRawMonitorGroupId(v int64) *CreateMonitorGroupRequest {
	s.RawMonitorGroupId = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetSecurityToken(v string) *CreateMonitorGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateMonitorGroupRequest) Validate() error {
	return dara.Validate(s)
}
