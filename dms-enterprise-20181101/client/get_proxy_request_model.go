// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProxyId(v int64) *GetProxyRequest
	GetProxyId() *int64
	SetTid(v int64) *GetProxyRequest
	GetTid() *int64
}

type GetProxyRequest struct {
	// The ID of the secure access proxy. You can call the [ListProxies](https://help.aliyun.com/document_detail/295371.html) operation to query the ID of the secure access proxy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4**
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProxyRequest) GoString() string {
	return s.String()
}

func (s *GetProxyRequest) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *GetProxyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetProxyRequest) SetProxyId(v int64) *GetProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *GetProxyRequest) SetTid(v int64) *GetProxyRequest {
	s.Tid = &v
	return s
}

func (s *GetProxyRequest) Validate() error {
	return dara.Validate(s)
}
