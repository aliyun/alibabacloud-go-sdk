// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserActiveTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *GetUserActiveTenantRequest
	GetTid() *int64
}

type GetUserActiveTenantRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetUserActiveTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserActiveTenantRequest) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetUserActiveTenantRequest) SetTid(v int64) *GetUserActiveTenantRequest {
	s.Tid = &v
	return s
}

func (s *GetUserActiveTenantRequest) Validate() error {
	return dara.Validate(s)
}
