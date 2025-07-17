// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeLhDagOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ChangeLhDagOwnerRequest
	GetDagId() *int64
	SetOwnerUserId(v int64) *ChangeLhDagOwnerRequest
	GetOwnerUserId() *int64
	SetTid(v int64) *ChangeLhDagOwnerRequest
	GetTid() *int64
}

type ChangeLhDagOwnerRequest struct {
	// The ID of the task flow. You can call the [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the user to be specified as the new owner of the task flow. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50****
	OwnerUserId *int64 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ChangeLhDagOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeLhDagOwnerRequest) GoString() string {
	return s.String()
}

func (s *ChangeLhDagOwnerRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ChangeLhDagOwnerRequest) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *ChangeLhDagOwnerRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ChangeLhDagOwnerRequest) SetDagId(v int64) *ChangeLhDagOwnerRequest {
	s.DagId = &v
	return s
}

func (s *ChangeLhDagOwnerRequest) SetOwnerUserId(v int64) *ChangeLhDagOwnerRequest {
	s.OwnerUserId = &v
	return s
}

func (s *ChangeLhDagOwnerRequest) SetTid(v int64) *ChangeLhDagOwnerRequest {
	s.Tid = &v
	return s
}

func (s *ChangeLhDagOwnerRequest) Validate() error {
	return dara.Validate(s)
}
