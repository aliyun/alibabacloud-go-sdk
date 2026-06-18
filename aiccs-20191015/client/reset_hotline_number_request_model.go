// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHotlineNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ResetHotlineNumberRequest
	GetDescription() *string
	SetEnableInbound(v bool) *ResetHotlineNumberRequest
	GetEnableInbound() *bool
	SetEnableInboundEvaluation(v bool) *ResetHotlineNumberRequest
	GetEnableInboundEvaluation() *bool
	SetEnableOutbound(v bool) *ResetHotlineNumberRequest
	GetEnableOutbound() *bool
	SetEnableOutboundEvaluation(v bool) *ResetHotlineNumberRequest
	GetEnableOutboundEvaluation() *bool
	SetEvaluationLevel(v int32) *ResetHotlineNumberRequest
	GetEvaluationLevel() *int32
	SetHotlineNumber(v string) *ResetHotlineNumberRequest
	GetHotlineNumber() *string
	SetInboundFlowId(v int64) *ResetHotlineNumberRequest
	GetInboundFlowId() *int64
	SetInstanceId(v string) *ResetHotlineNumberRequest
	GetInstanceId() *string
	SetOutboundAllDepart(v bool) *ResetHotlineNumberRequest
	GetOutboundAllDepart() *bool
	SetOutboundRangeList(v []*ResetHotlineNumberRequestOutboundRangeList) *ResetHotlineNumberRequest
	GetOutboundRangeList() []*ResetHotlineNumberRequestOutboundRangeList
}

type ResetHotlineNumberRequest struct {
	// Number description.
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
	// Whether inbound call satisfaction evaluation is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// Whether used for outbound calls.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// Indicates whether outbound satisfaction evaluation is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// Satisfaction level. Valid values:
	//
	// - **2**: Two-level (Satisfied, Not Satisfied)
	//
	// - **3**: Three-level (Satisfied, Neutral, Not Satisfied)
	//
	// - **4**: Four-level (Very Satisfied, Satisfied, Neutral, Not Satisfied)
	//
	// - **5**: Five-level (Very Satisfied, Satisfied, Neutral, Not Satisfied, Very Poor)
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
	// 0571********
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The IVR flow ID for inbound calls.
	//
	// example:
	//
	// 123456
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID. You can obtain it from the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether outbound calls apply to all departments.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	OutboundAllDepart *bool `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	// Outbound call effective scope.
	OutboundRangeList []*ResetHotlineNumberRequestOutboundRangeList `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty" type:"Repeated"`
}

func (s ResetHotlineNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberRequest) GetDescription() *string {
	return s.Description
}

func (s *ResetHotlineNumberRequest) GetEnableInbound() *bool {
	return s.EnableInbound
}

func (s *ResetHotlineNumberRequest) GetEnableInboundEvaluation() *bool {
	return s.EnableInboundEvaluation
}

func (s *ResetHotlineNumberRequest) GetEnableOutbound() *bool {
	return s.EnableOutbound
}

func (s *ResetHotlineNumberRequest) GetEnableOutboundEvaluation() *bool {
	return s.EnableOutboundEvaluation
}

func (s *ResetHotlineNumberRequest) GetEvaluationLevel() *int32 {
	return s.EvaluationLevel
}

func (s *ResetHotlineNumberRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *ResetHotlineNumberRequest) GetInboundFlowId() *int64 {
	return s.InboundFlowId
}

func (s *ResetHotlineNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetHotlineNumberRequest) GetOutboundAllDepart() *bool {
	return s.OutboundAllDepart
}

func (s *ResetHotlineNumberRequest) GetOutboundRangeList() []*ResetHotlineNumberRequestOutboundRangeList {
	return s.OutboundRangeList
}

func (s *ResetHotlineNumberRequest) SetDescription(v string) *ResetHotlineNumberRequest {
	s.Description = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableInbound(v bool) *ResetHotlineNumberRequest {
	s.EnableInbound = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableInboundEvaluation(v bool) *ResetHotlineNumberRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableOutbound(v bool) *ResetHotlineNumberRequest {
	s.EnableOutbound = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableOutboundEvaluation(v bool) *ResetHotlineNumberRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEvaluationLevel(v int32) *ResetHotlineNumberRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetHotlineNumber(v string) *ResetHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetInboundFlowId(v int64) *ResetHotlineNumberRequest {
	s.InboundFlowId = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetInstanceId(v string) *ResetHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetOutboundAllDepart(v bool) *ResetHotlineNumberRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetOutboundRangeList(v []*ResetHotlineNumberRequestOutboundRangeList) *ResetHotlineNumberRequest {
	s.OutboundRangeList = v
	return s
}

func (s *ResetHotlineNumberRequest) Validate() error {
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

type ResetHotlineNumberRequestOutboundRangeList struct {
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

func (s ResetHotlineNumberRequestOutboundRangeList) String() string {
	return dara.Prettify(s)
}

func (s ResetHotlineNumberRequestOutboundRangeList) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberRequestOutboundRangeList) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *ResetHotlineNumberRequestOutboundRangeList) GetGroupIdList() []*int64 {
	return s.GroupIdList
}

func (s *ResetHotlineNumberRequestOutboundRangeList) SetDepartmentId(v int64) *ResetHotlineNumberRequestOutboundRangeList {
	s.DepartmentId = &v
	return s
}

func (s *ResetHotlineNumberRequestOutboundRangeList) SetGroupIdList(v []*int64) *ResetHotlineNumberRequestOutboundRangeList {
	s.GroupIdList = v
	return s
}

func (s *ResetHotlineNumberRequestOutboundRangeList) Validate() error {
	return dara.Validate(s)
}
