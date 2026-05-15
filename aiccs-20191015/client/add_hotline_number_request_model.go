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
	// 05710000****
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
	OutboundAllDepart *bool                                       `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
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
	// example:
	//
	// 123456
	DepartmentId *int64   `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	GroupIdList  []*int64 `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
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
