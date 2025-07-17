// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSLARulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateSLARulesShrinkRequest
	GetDagId() *int64
	SetSlaRuleListShrink(v string) *UpdateSLARulesShrinkRequest
	GetSlaRuleListShrink() *string
	SetTid(v int64) *UpdateSLARulesShrinkRequest
	GetTid() *int64
}

type UpdateSLARulesShrinkRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The list of SLA rules.
	SlaRuleListShrink *string `json:"SlaRuleList,omitempty" xml:"SlaRuleList,omitempty"`
	// The ID of the tenant.
	//
	// > :To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateSLARulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSLARulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSLARulesShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateSLARulesShrinkRequest) GetSlaRuleListShrink() *string {
	return s.SlaRuleListShrink
}

func (s *UpdateSLARulesShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateSLARulesShrinkRequest) SetDagId(v int64) *UpdateSLARulesShrinkRequest {
	s.DagId = &v
	return s
}

func (s *UpdateSLARulesShrinkRequest) SetSlaRuleListShrink(v string) *UpdateSLARulesShrinkRequest {
	s.SlaRuleListShrink = &v
	return s
}

func (s *UpdateSLARulesShrinkRequest) SetTid(v int64) *UpdateSLARulesShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateSLARulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
