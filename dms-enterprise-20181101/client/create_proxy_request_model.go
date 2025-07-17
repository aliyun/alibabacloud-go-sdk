// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *CreateProxyRequest
	GetInstanceId() *int64
	SetPassword(v string) *CreateProxyRequest
	GetPassword() *string
	SetTid(v int64) *CreateProxyRequest
	GetTid() *int64
	SetUsername(v string) *CreateProxyRequest
	GetUsername() *string
}

type CreateProxyRequest struct {
	// The ID of the database instance. You can call the [ListInstances](https://www.alibabacloud.com/help/en/data-management-service/latest/listinstances) or [GetInstance](https://www.alibabacloud.com/help/en/data-management-service/latest/getinstance) operation to query the database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 183****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The username of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyRequest) GoString() string {
	return s.String()
}

func (s *CreateProxyRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CreateProxyRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateProxyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateProxyRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateProxyRequest) SetInstanceId(v int64) *CreateProxyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateProxyRequest) SetPassword(v string) *CreateProxyRequest {
	s.Password = &v
	return s
}

func (s *CreateProxyRequest) SetTid(v int64) *CreateProxyRequest {
	s.Tid = &v
	return s
}

func (s *CreateProxyRequest) SetUsername(v string) *CreateProxyRequest {
	s.Username = &v
	return s
}

func (s *CreateProxyRequest) Validate() error {
	return dara.Validate(s)
}
