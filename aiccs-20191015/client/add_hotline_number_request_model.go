// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotlineNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddHotlineNumberRequest
	GetDescription() *string
	SetEnableInbound(v bool) *AddHotlineNumberRequest
	GetEnableInbound() *bool
	SetEnableInboundEvaluation(v bool) *AddHotlineNumberRequest
	GetEnableInboundEvaluation() *bool
	SetEnableOutbound(v bool) *AddHotlineNumberRequest
	GetEnableOutbound() *bool
	SetEnableOutboundEvaluation(v bool) *AddHotlineNumberRequest
	GetEnableOutboundEvaluation() *bool
	SetEvaluationLevel(v int32) *AddHotlineNumberRequest
	GetEvaluationLevel() *int32
	SetHotlineNumber(v string) *AddHotlineNumberRequest
	GetHotlineNumber() *string
	SetInboundFlowId(v int64) *AddHotlineNumberRequest
	GetInboundFlowId() *int64
	SetInstanceId(v string) *AddHotlineNumberRequest
	GetInstanceId() *string
	SetOutboundAllDepart(v bool) *AddHotlineNumberRequest
	GetOutboundAllDepart() *bool
	SetOutboundRangeList(v []*AddHotlineNumberRequestOutboundRangeList) *AddHotlineNumberRequest
	GetOutboundRangeList() []*AddHotlineNumberRequestOutboundRangeList
}

type AddHotlineNumberRequest struct {
	// Description of the number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 热线号码
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the number is used for inbound calls.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableInbound *bool `json:"EnableInbound,omitempty" xml:"EnableInbound,omitempty"`
	// Whether inbound satisfaction evaluation is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// Whether it is used for outbound calls.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// Indicates whether outbound call satisfaction evaluation is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// Satisfaction Level. Valid values:
	//
	// - **2**: Two-level (Satisfied, Not satisfied)
	//
	// - **3**: Three-level (Satisfied, Neutral, Not satisfied)
	//
	// - **4**: Four-level (Very satisfied, Satisfied, Neutral, Not satisfied)
	//
	// - **5**: Five-level (Very satisfied, Satisfied, Neutral, Not satisfied, Very poor)
	//
	// example:
	//
	// 2
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// Hotline number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05710000****
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The IVR flow ID for inbound calls. You can obtain it on the SaaS Workbench > Channel Integration > IVR Flow Management page.
	//
	// example:
	//
	// 123456
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID. You can obtain it in the Intelligent Contact Center console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Whether outbound calls apply to all departments under the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	OutboundAllDepart *bool `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	// Outbound call effective scope.
	OutboundRangeList []*AddHotlineNumberRequestOutboundRangeList `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty" type:"Repeated"`
}

func (s AddHotlineNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberRequest) GetDescription() *string {
	return s.Description
}

func (s *AddHotlineNumberRequest) GetEnableInbound() *bool {
	return s.EnableInbound
}

func (s *AddHotlineNumberRequest) GetEnableInboundEvaluation() *bool {
	return s.EnableInboundEvaluation
}

func (s *AddHotlineNumberRequest) GetEnableOutbound() *bool {
	return s.EnableOutbound
}

func (s *AddHotlineNumberRequest) GetEnableOutboundEvaluation() *bool {
	return s.EnableOutboundEvaluation
}

func (s *AddHotlineNumberRequest) GetEvaluationLevel() *int32 {
	return s.EvaluationLevel
}

func (s *AddHotlineNumberRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *AddHotlineNumberRequest) GetInboundFlowId() *int64 {
	return s.InboundFlowId
}

func (s *AddHotlineNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHotlineNumberRequest) GetOutboundAllDepart() *bool {
	return s.OutboundAllDepart
}

func (s *AddHotlineNumberRequest) GetOutboundRangeList() []*AddHotlineNumberRequestOutboundRangeList {
	return s.OutboundRangeList
}

func (s *AddHotlineNumberRequest) SetDescription(v string) *AddHotlineNumberRequest {
	s.Description = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableInbound(v bool) *AddHotlineNumberRequest {
	s.EnableInbound = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableInboundEvaluation(v bool) *AddHotlineNumberRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableOutbound(v bool) *AddHotlineNumberRequest {
	s.EnableOutbound = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableOutboundEvaluation(v bool) *AddHotlineNumberRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEvaluationLevel(v int32) *AddHotlineNumberRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *AddHotlineNumberRequest) SetHotlineNumber(v string) *AddHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *AddHotlineNumberRequest) SetInboundFlowId(v int64) *AddHotlineNumberRequest {
	s.InboundFlowId = &v
	return s
}

func (s *AddHotlineNumberRequest) SetInstanceId(v string) *AddHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHotlineNumberRequest) SetOutboundAllDepart(v bool) *AddHotlineNumberRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *AddHotlineNumberRequest) SetOutboundRangeList(v []*AddHotlineNumberRequestOutboundRangeList) *AddHotlineNumberRequest {
	s.OutboundRangeList = v
	return s
}

func (s *AddHotlineNumberRequest) Validate() error {
	if s.OutboundRangeList != nil {
		for _, item := range s.OutboundRangeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddHotlineNumberRequestOutboundRangeList struct {
	// Effective department ID.
	//
	// example:
	//
	// 123456
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// List of effective skill groups.
	//
	// > If the skill group list is empty, the setting applies to the entire department. Otherwise, it applies only to the specified skill groups under the department.
	GroupIdList []*int64 `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
}

func (s AddHotlineNumberRequestOutboundRangeList) String() string {
	return dara.Prettify(s)
}

func (s AddHotlineNumberRequestOutboundRangeList) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberRequestOutboundRangeList) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *AddHotlineNumberRequestOutboundRangeList) GetGroupIdList() []*int64 {
	return s.GroupIdList
}

func (s *AddHotlineNumberRequestOutboundRangeList) SetDepartmentId(v int64) *AddHotlineNumberRequestOutboundRangeList {
	s.DepartmentId = &v
	return s
}

func (s *AddHotlineNumberRequestOutboundRangeList) SetGroupIdList(v []*int64) *AddHotlineNumberRequestOutboundRangeList {
	s.GroupIdList = v
	return s
}

func (s *AddHotlineNumberRequestOutboundRangeList) Validate() error {
	return dara.Validate(s)
}
