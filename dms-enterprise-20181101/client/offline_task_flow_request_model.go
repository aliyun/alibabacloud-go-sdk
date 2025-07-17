// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineTaskFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *OfflineTaskFlowRequest
	GetDagId() *int64
	SetTid(v int64) *OfflineTaskFlowRequest
	GetTid() *int64
}

type OfflineTaskFlowRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s OfflineTaskFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *OfflineTaskFlowRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *OfflineTaskFlowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *OfflineTaskFlowRequest) SetDagId(v int64) *OfflineTaskFlowRequest {
	s.DagId = &v
	return s
}

func (s *OfflineTaskFlowRequest) SetTid(v int64) *OfflineTaskFlowRequest {
	s.Tid = &v
	return s
}

func (s *OfflineTaskFlowRequest) Validate() error {
	return dara.Validate(s)
}
