// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSLARulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateSLARulesRequest
	GetDagId() *int64
	SetSlaRuleList(v []*UpdateSLARulesRequestSlaRuleList) *UpdateSLARulesRequest
	GetSlaRuleList() []*UpdateSLARulesRequestSlaRuleList
	SetTid(v int64) *UpdateSLARulesRequest
	GetTid() *int64
}

type UpdateSLARulesRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The list of SLA rules.
	SlaRuleList []*UpdateSLARulesRequestSlaRuleList `json:"SlaRuleList,omitempty" xml:"SlaRuleList,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// > :To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateSLARulesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSLARulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateSLARulesRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateSLARulesRequest) GetSlaRuleList() []*UpdateSLARulesRequestSlaRuleList {
	return s.SlaRuleList
}

func (s *UpdateSLARulesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateSLARulesRequest) SetDagId(v int64) *UpdateSLARulesRequest {
	s.DagId = &v
	return s
}

func (s *UpdateSLARulesRequest) SetSlaRuleList(v []*UpdateSLARulesRequestSlaRuleList) *UpdateSLARulesRequest {
	s.SlaRuleList = v
	return s
}

func (s *UpdateSLARulesRequest) SetTid(v int64) *UpdateSLARulesRequest {
	s.Tid = &v
	return s
}

func (s *UpdateSLARulesRequest) Validate() error {
	if s.SlaRuleList != nil {
		for _, item := range s.SlaRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSLARulesRequestSlaRuleList struct {
	// The ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The timeout period. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1080
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
	// The ID of the task node.
	//
	// example:
	//
	// 0
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The rule type. Valid values:
	//
	// 	- **0**: SLA rules for task flows
	//
	// 	- **1**: SLA rules for nodes
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateSLARulesRequestSlaRuleList) String() string {
	return dara.Prettify(s)
}

func (s UpdateSLARulesRequestSlaRuleList) GoString() string {
	return s.String()
}

func (s *UpdateSLARulesRequestSlaRuleList) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateSLARulesRequestSlaRuleList) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *UpdateSLARulesRequestSlaRuleList) GetNodeId() *int64 {
	return s.NodeId
}

func (s *UpdateSLARulesRequestSlaRuleList) GetType() *int32 {
	return s.Type
}

func (s *UpdateSLARulesRequestSlaRuleList) SetDagId(v int64) *UpdateSLARulesRequestSlaRuleList {
	s.DagId = &v
	return s
}

func (s *UpdateSLARulesRequestSlaRuleList) SetIntervalMinutes(v int32) *UpdateSLARulesRequestSlaRuleList {
	s.IntervalMinutes = &v
	return s
}

func (s *UpdateSLARulesRequestSlaRuleList) SetNodeId(v int64) *UpdateSLARulesRequestSlaRuleList {
	s.NodeId = &v
	return s
}

func (s *UpdateSLARulesRequestSlaRuleList) SetType(v int32) *UpdateSLARulesRequestSlaRuleList {
	s.Type = &v
	return s
}

func (s *UpdateSLARulesRequestSlaRuleList) Validate() error {
	return dara.Validate(s)
}
