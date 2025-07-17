// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSLARulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ListSLARulesRequest
	GetDagId() *int64
	SetTid(v int64) *ListSLARulesRequest
	GetTid() *int64
}

type ListSLARulesRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSLARulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSLARulesRequest) GoString() string {
	return s.String()
}

func (s *ListSLARulesRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListSLARulesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSLARulesRequest) SetDagId(v int64) *ListSLARulesRequest {
	s.DagId = &v
	return s
}

func (s *ListSLARulesRequest) SetTid(v int64) *ListSLARulesRequest {
	s.Tid = &v
	return s
}

func (s *ListSLARulesRequest) Validate() error {
	return dara.Validate(s)
}
