// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProxyAccessId(v int64) *GetProxyAccessRequest
	GetProxyAccessId() *int64
	SetTid(v int64) *GetProxyAccessRequest
	GetTid() *int64
}

type GetProxyAccessRequest struct {
	// The ID that Data Management (DMS) generates after the user is authorized to enable the secure access proxy feature for an instance. The ID is unique in DMS. You can call the [ListProxyAccesses](https://help.aliyun.com/document_detail/295386.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProxyAccessId *int64 `json:"ProxyAccessId,omitempty" xml:"ProxyAccessId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetProxyAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProxyAccessRequest) GoString() string {
	return s.String()
}

func (s *GetProxyAccessRequest) GetProxyAccessId() *int64 {
	return s.ProxyAccessId
}

func (s *GetProxyAccessRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetProxyAccessRequest) SetProxyAccessId(v int64) *GetProxyAccessRequest {
	s.ProxyAccessId = &v
	return s
}

func (s *GetProxyAccessRequest) SetTid(v int64) *GetProxyAccessRequest {
	s.Tid = &v
	return s
}

func (s *GetProxyAccessRequest) Validate() error {
	return dara.Validate(s)
}
