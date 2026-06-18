// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotlineNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddHotlineNumberShrinkRequest
	GetDescription() *string
	SetEnableInbound(v bool) *AddHotlineNumberShrinkRequest
	GetEnableInbound() *bool
	SetEnableInboundEvaluation(v bool) *AddHotlineNumberShrinkRequest
	GetEnableInboundEvaluation() *bool
	SetEnableOutbound(v bool) *AddHotlineNumberShrinkRequest
	GetEnableOutbound() *bool
	SetEnableOutboundEvaluation(v bool) *AddHotlineNumberShrinkRequest
	GetEnableOutboundEvaluation() *bool
	SetEvaluationLevel(v int32) *AddHotlineNumberShrinkRequest
	GetEvaluationLevel() *int32
	SetHotlineNumber(v string) *AddHotlineNumberShrinkRequest
	GetHotlineNumber() *string
	SetInboundFlowId(v int64) *AddHotlineNumberShrinkRequest
	GetInboundFlowId() *int64
	SetInstanceId(v string) *AddHotlineNumberShrinkRequest
	GetInstanceId() *string
	SetOutboundAllDepart(v bool) *AddHotlineNumberShrinkRequest
	GetOutboundAllDepart() *bool
	SetOutboundRangeListShrink(v string) *AddHotlineNumberShrinkRequest
	GetOutboundRangeListShrink() *string
}

type AddHotlineNumberShrinkRequest struct {
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
	OutboundRangeListShrink *string `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty"`
}

func (s AddHotlineNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHotlineNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddHotlineNumberShrinkRequest) GetEnableInbound() *bool {
	return s.EnableInbound
}

func (s *AddHotlineNumberShrinkRequest) GetEnableInboundEvaluation() *bool {
	return s.EnableInboundEvaluation
}

func (s *AddHotlineNumberShrinkRequest) GetEnableOutbound() *bool {
	return s.EnableOutbound
}

func (s *AddHotlineNumberShrinkRequest) GetEnableOutboundEvaluation() *bool {
	return s.EnableOutboundEvaluation
}

func (s *AddHotlineNumberShrinkRequest) GetEvaluationLevel() *int32 {
	return s.EvaluationLevel
}

func (s *AddHotlineNumberShrinkRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *AddHotlineNumberShrinkRequest) GetInboundFlowId() *int64 {
	return s.InboundFlowId
}

func (s *AddHotlineNumberShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHotlineNumberShrinkRequest) GetOutboundAllDepart() *bool {
	return s.OutboundAllDepart
}

func (s *AddHotlineNumberShrinkRequest) GetOutboundRangeListShrink() *string {
	return s.OutboundRangeListShrink
}

func (s *AddHotlineNumberShrinkRequest) SetDescription(v string) *AddHotlineNumberShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableInbound(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableInbound = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableInboundEvaluation(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableOutbound(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableOutbound = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableOutboundEvaluation(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEvaluationLevel(v int32) *AddHotlineNumberShrinkRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetHotlineNumber(v string) *AddHotlineNumberShrinkRequest {
	s.HotlineNumber = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetInboundFlowId(v int64) *AddHotlineNumberShrinkRequest {
	s.InboundFlowId = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetInstanceId(v string) *AddHotlineNumberShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetOutboundAllDepart(v bool) *AddHotlineNumberShrinkRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetOutboundRangeListShrink(v string) *AddHotlineNumberShrinkRequest {
	s.OutboundRangeListShrink = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
