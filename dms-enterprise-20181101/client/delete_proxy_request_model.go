// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProxyId(v int64) *DeleteProxyRequest
	GetProxyId() *int64
	SetTid(v int64) *DeleteProxyRequest
	GetTid() *int64
}

type DeleteProxyRequest struct {
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
}

func (s DeleteProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProxyRequest) GoString() string {
	return s.String()
}

func (s *DeleteProxyRequest) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *DeleteProxyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteProxyRequest) SetProxyId(v int64) *DeleteProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *DeleteProxyRequest) SetTid(v int64) *DeleteProxyRequest {
	s.Tid = &v
	return s
}

func (s *DeleteProxyRequest) Validate() error {
	return dara.Validate(s)
}
