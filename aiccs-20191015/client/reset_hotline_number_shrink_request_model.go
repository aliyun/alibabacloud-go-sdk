// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHotlineNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ResetHotlineNumberShrinkRequest
	GetDescription() *string
	SetEnableInbound(v bool) *ResetHotlineNumberShrinkRequest
	GetEnableInbound() *bool
	SetEnableInboundEvaluation(v bool) *ResetHotlineNumberShrinkRequest
	GetEnableInboundEvaluation() *bool
	SetEnableOutbound(v bool) *ResetHotlineNumberShrinkRequest
	GetEnableOutbound() *bool
	SetEnableOutboundEvaluation(v bool) *ResetHotlineNumberShrinkRequest
	GetEnableOutboundEvaluation() *bool
	SetEvaluationLevel(v int32) *ResetHotlineNumberShrinkRequest
	GetEvaluationLevel() *int32
	SetHotlineNumber(v string) *ResetHotlineNumberShrinkRequest
	GetHotlineNumber() *string
	SetInboundFlowId(v int64) *ResetHotlineNumberShrinkRequest
	GetInboundFlowId() *int64
	SetInstanceId(v string) *ResetHotlineNumberShrinkRequest
	GetInstanceId() *string
	SetOutboundAllDepart(v bool) *ResetHotlineNumberShrinkRequest
	GetOutboundAllDepart() *bool
	SetOutboundRangeListShrink(v string) *ResetHotlineNumberShrinkRequest
	GetOutboundRangeListShrink() *string
}

type ResetHotlineNumberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableInbound *bool `json:"EnableInbound,omitempty" xml:"EnableInbound,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// example:
	//
	// 2
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0571********
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// example:
	//
	// 123456
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	OutboundAllDepart       *bool   `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	OutboundRangeListShrink *string `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty"`
}

func (s ResetHotlineNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetHotlineNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ResetHotlineNumberShrinkRequest) GetEnableInbound() *bool {
	return s.EnableInbound
}

func (s *ResetHotlineNumberShrinkRequest) GetEnableInboundEvaluation() *bool {
	return s.EnableInboundEvaluation
}

func (s *ResetHotlineNumberShrinkRequest) GetEnableOutbound() *bool {
	return s.EnableOutbound
}

func (s *ResetHotlineNumberShrinkRequest) GetEnableOutboundEvaluation() *bool {
	return s.EnableOutboundEvaluation
}

func (s *ResetHotlineNumberShrinkRequest) GetEvaluationLevel() *int32 {
	return s.EvaluationLevel
}

func (s *ResetHotlineNumberShrinkRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *ResetHotlineNumberShrinkRequest) GetInboundFlowId() *int64 {
	return s.InboundFlowId
}

func (s *ResetHotlineNumberShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetHotlineNumberShrinkRequest) GetOutboundAllDepart() *bool {
	return s.OutboundAllDepart
}

func (s *ResetHotlineNumberShrinkRequest) GetOutboundRangeListShrink() *string {
	return s.OutboundRangeListShrink
}

func (s *ResetHotlineNumberShrinkRequest) SetDescription(v string) *ResetHotlineNumberShrinkRequest {
	s.Description = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableInbound(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableInbound = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableInboundEvaluation(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableOutbound(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableOutbound = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableOutboundEvaluation(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEvaluationLevel(v int32) *ResetHotlineNumberShrinkRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetHotlineNumber(v string) *ResetHotlineNumberShrinkRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetInboundFlowId(v int64) *ResetHotlineNumberShrinkRequest {
	s.InboundFlowId = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetInstanceId(v string) *ResetHotlineNumberShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetOutboundAllDepart(v bool) *ResetHotlineNumberShrinkRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetOutboundRangeListShrink(v string) *ResetHotlineNumberShrinkRequest {
	s.OutboundRangeListShrink = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
