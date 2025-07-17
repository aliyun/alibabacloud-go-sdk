// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *ListTaskFlowRequest
	GetTid() *int64
}

type ListTaskFlowRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTaskFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *ListTaskFlowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTaskFlowRequest) SetTid(v int64) *ListTaskFlowRequest {
	s.Tid = &v
	return s
}

func (s *ListTaskFlowRequest) Validate() error {
	return dara.Validate(s)
}
