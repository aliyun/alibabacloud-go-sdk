// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndepAccount(v string) *CreateProxyAccessRequest
	GetIndepAccount() *string
	SetIndepPassword(v string) *CreateProxyAccessRequest
	GetIndepPassword() *string
	SetProxyId(v int64) *CreateProxyAccessRequest
	GetProxyId() *int64
	SetTid(v int64) *CreateProxyAccessRequest
	GetTid() *int64
	SetUserId(v int64) *CreateProxyAccessRequest
	GetUserId() *int64
}

type CreateProxyAccessRequest struct {
	// The database account.
	//
	// example:
	//
	// xxx
	IndepAccount *string `json:"IndepAccount,omitempty" xml:"IndepAccount,omitempty"`
	// The password that is used to log on to the database.
	//
	// example:
	//
	// xxx
	IndepPassword *string `json:"IndepPassword,omitempty" xml:"IndepPassword,omitempty"`
	// The ID of the security protection agent. You can call the [ListProxies](https://www.alibabacloud.com/help/en/data-management-service/latest/listproxies) or [GetProxy](https://www.alibabacloud.com/help/en/data-management-service/latest/getproxy) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) or [ListUserTenants](https://www.alibabacloud.com/help/en/data-management-service/latest/listusertenants) operation to obtain this parameter.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the user. You can call the [ListUsers](https://www.alibabacloud.com/help/en/data-management-service/latest/listusers) or [GetUser](https://www.alibabacloud.com/help/en/data-management-service/latest/getuser) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProxyAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyAccessRequest) GoString() string {
	return s.String()
}

func (s *CreateProxyAccessRequest) GetIndepAccount() *string {
	return s.IndepAccount
}

func (s *CreateProxyAccessRequest) GetIndepPassword() *string {
	return s.IndepPassword
}

func (s *CreateProxyAccessRequest) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *CreateProxyAccessRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateProxyAccessRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateProxyAccessRequest) SetIndepAccount(v string) *CreateProxyAccessRequest {
	s.IndepAccount = &v
	return s
}

func (s *CreateProxyAccessRequest) SetIndepPassword(v string) *CreateProxyAccessRequest {
	s.IndepPassword = &v
	return s
}

func (s *CreateProxyAccessRequest) SetProxyId(v int64) *CreateProxyAccessRequest {
	s.ProxyId = &v
	return s
}

func (s *CreateProxyAccessRequest) SetTid(v int64) *CreateProxyAccessRequest {
	s.Tid = &v
	return s
}

func (s *CreateProxyAccessRequest) SetUserId(v int64) *CreateProxyAccessRequest {
	s.UserId = &v
	return s
}

func (s *CreateProxyAccessRequest) Validate() error {
	return dara.Validate(s)
}
