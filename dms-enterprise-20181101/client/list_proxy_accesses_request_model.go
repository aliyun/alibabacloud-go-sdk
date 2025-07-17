// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxyAccessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProxyId(v int64) *ListProxyAccessesRequest
	GetProxyId() *int64
	SetTid(v int64) *ListProxyAccessesRequest
	GetTid() *int64
}

type ListProxyAccessesRequest struct {
	// The ID of the secure access proxy.
	//
	// >  You can call the [ListProxies](https://www.alibabacloud.com/help/en/data-management-service/latest/listproxies) operation to query the ID of the secure access proxy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The ID of the tenant.
	//
	// >  You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) operation to query the ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListProxyAccessesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProxyAccessesRequest) GoString() string {
	return s.String()
}

func (s *ListProxyAccessesRequest) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *ListProxyAccessesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListProxyAccessesRequest) SetProxyId(v int64) *ListProxyAccessesRequest {
	s.ProxyId = &v
	return s
}

func (s *ListProxyAccessesRequest) SetTid(v int64) *ListProxyAccessesRequest {
	s.Tid = &v
	return s
}

func (s *ListProxyAccessesRequest) Validate() error {
	return dara.Validate(s)
}
