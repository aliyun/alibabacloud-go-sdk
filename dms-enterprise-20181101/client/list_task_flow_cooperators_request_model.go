// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowCooperatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ListTaskFlowCooperatorsRequest
	GetDagId() *int64
	SetTid(v int64) *ListTaskFlowCooperatorsRequest
	GetTid() *int64
}

type ListTaskFlowCooperatorsRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTaskFlowCooperatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowCooperatorsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskFlowCooperatorsRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListTaskFlowCooperatorsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTaskFlowCooperatorsRequest) SetDagId(v int64) *ListTaskFlowCooperatorsRequest {
	s.DagId = &v
	return s
}

func (s *ListTaskFlowCooperatorsRequest) SetTid(v int64) *ListTaskFlowCooperatorsRequest {
	s.Tid = &v
	return s
}

func (s *ListTaskFlowCooperatorsRequest) Validate() error {
	return dara.Validate(s)
}
