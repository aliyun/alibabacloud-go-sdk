// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCustomizeRuleCountResponseBody
	GetCode() *int32
	SetData(v *DescribeCustomizeRuleCountResponseBodyData) *DescribeCustomizeRuleCountResponseBody
	GetData() *DescribeCustomizeRuleCountResponseBodyData
	SetMessage(v string) *DescribeCustomizeRuleCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomizeRuleCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCustomizeRuleCountResponseBody
	GetSuccess() *bool
}

type DescribeCustomizeRuleCountResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCustomizeRuleCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCustomizeRuleCountResponseBody) GetData() *DescribeCustomizeRuleCountResponseBodyData {
	return s.Data
}

func (s *DescribeCustomizeRuleCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomizeRuleCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizeRuleCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomizeRuleCountResponseBody) SetCode(v int32) *DescribeCustomizeRuleCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetData(v *DescribeCustomizeRuleCountResponseBodyData) *DescribeCustomizeRuleCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetMessage(v string) *DescribeCustomizeRuleCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetRequestId(v string) *DescribeCustomizeRuleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomizeRuleCountResponseBodyData struct {
	// 同类聚合规则数。
	//
	// example:
	//
	// 3
	AggregationRuleNum *int32 `json:"AggregationRuleNum,omitempty" xml:"AggregationRuleNum,omitempty"`
	// 自定义规则数。
	//
	// example:
	//
	// 10
	CustomizeRuleNum *int32 `json:"CustomizeRuleNum,omitempty" xml:"CustomizeRuleNum,omitempty"`
	// 专家规则数。
	//
	// example:
	//
	// 7
	ExpertRuleNum *int32 `json:"ExpertRuleNum,omitempty" xml:"ExpertRuleNum,omitempty"`
	// 图计算规则数。
	//
	// example:
	//
	// 2
	GraphComputingRuleNum *int32 `json:"GraphComputingRuleNum,omitempty" xml:"GraphComputingRuleNum,omitempty"`
	// The number of rules that are used to identify high-risk threats.
	//
	// example:
	//
	// 12
	HighRuleNum *int32 `json:"HighRuleNum,omitempty" xml:"HighRuleNum,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 20
	InUseRuleNum *int32 `json:"InUseRuleNum,omitempty" xml:"InUseRuleNum,omitempty"`
	// The number of rules that are used to identify low-risk threats.
	//
	// example:
	//
	// 3
	LowRuleNum *int32 `json:"LowRuleNum,omitempty" xml:"LowRuleNum,omitempty"`
	// The number of rules that are used to identify medium-risk threats.
	//
	// example:
	//
	// 5
	MediumRuleNum *int32 `json:"MediumRuleNum,omitempty" xml:"MediumRuleNum,omitempty"`
	// 预定义规则数。
	//
	// example:
	//
	// 10
	PredefinedRuleNum *int32 `json:"PredefinedRuleNum,omitempty" xml:"PredefinedRuleNum,omitempty"`
	// 告警透传规则数。
	//
	// example:
	//
	// 3
	SingleAlertRuleNum *int32 `json:"SingleAlertRuleNum,omitempty" xml:"SingleAlertRuleNum,omitempty"`
	// 总规则数。
	//
	// example:
	//
	// 10
	TotalRuleNum *int32 `json:"TotalRuleNum,omitempty" xml:"TotalRuleNum,omitempty"`
	// 不产生事件规则数。
	//
	// example:
	//
	// 3
	UnEventRuleNum *int32 `json:"UnEventRuleNum,omitempty" xml:"UnEventRuleNum,omitempty"`
}

func (s DescribeCustomizeRuleCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetAggregationRuleNum() *int32 {
	return s.AggregationRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetCustomizeRuleNum() *int32 {
	return s.CustomizeRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetExpertRuleNum() *int32 {
	return s.ExpertRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetGraphComputingRuleNum() *int32 {
	return s.GraphComputingRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetHighRuleNum() *int32 {
	return s.HighRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetInUseRuleNum() *int32 {
	return s.InUseRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetLowRuleNum() *int32 {
	return s.LowRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetMediumRuleNum() *int32 {
	return s.MediumRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetPredefinedRuleNum() *int32 {
	return s.PredefinedRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetSingleAlertRuleNum() *int32 {
	return s.SingleAlertRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetTotalRuleNum() *int32 {
	return s.TotalRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) GetUnEventRuleNum() *int32 {
	return s.UnEventRuleNum
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetAggregationRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.AggregationRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetCustomizeRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.CustomizeRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetExpertRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.ExpertRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetGraphComputingRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.GraphComputingRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetHighRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.HighRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetInUseRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.InUseRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetLowRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.LowRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetMediumRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.MediumRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetPredefinedRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.PredefinedRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetSingleAlertRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.SingleAlertRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetTotalRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.TotalRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetUnEventRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.UnEventRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
