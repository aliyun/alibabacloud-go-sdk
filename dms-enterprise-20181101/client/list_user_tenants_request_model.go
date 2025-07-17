// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserTenantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *ListUserTenantsRequest
	GetTid() *int64
}

type ListUserTenantsRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListUserTenantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserTenantsRequest) GoString() string {
	return s.String()
}

func (s *ListUserTenantsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListUserTenantsRequest) SetTid(v int64) *ListUserTenantsRequest {
	s.Tid = &v
	return s
}

func (s *ListUserTenantsRequest) Validate() error {
	return dara.Validate(s)
}
