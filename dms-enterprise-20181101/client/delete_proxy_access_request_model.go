// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProxyAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProxyAccessId(v int64) *DeleteProxyAccessRequest
	GetProxyAccessId() *int64
	SetTid(v int64) *DeleteProxyAccessRequest
	GetTid() *int64
}

type DeleteProxyAccessRequest struct {
	// The ID of the security protection authorization. After the security protection agent authorizes the target user, the system automatically generates a security protection authorization ID. The ID is globally unique. You can call the [ListProxyAccesses](https://www.alibabacloud.com/help/en/data-management-service/latest/listproxyaccesses) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProxyAccessId *int64 `json:"ProxyAccessId,omitempty" xml:"ProxyAccessId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) or [ListUserTenants](https://www.alibabacloud.com/help/en/data-management-service/latest/listusertenants) operation to obtain this parameter.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteProxyAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProxyAccessRequest) GoString() string {
	return s.String()
}

func (s *DeleteProxyAccessRequest) GetProxyAccessId() *int64 {
	return s.ProxyAccessId
}

func (s *DeleteProxyAccessRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteProxyAccessRequest) SetProxyAccessId(v int64) *DeleteProxyAccessRequest {
	s.ProxyAccessId = &v
	return s
}

func (s *DeleteProxyAccessRequest) SetTid(v int64) *DeleteProxyAccessRequest {
	s.Tid = &v
	return s
}

func (s *DeleteProxyAccessRequest) Validate() error {
	return dara.Validate(s)
}
