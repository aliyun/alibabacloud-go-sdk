// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayElasticPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayElasticPolicyResponseBody
	GetCode() *string
	SetData(v *GetGatewayElasticPolicyResponseBodyData) *GetGatewayElasticPolicyResponseBody
	GetData() *GetGatewayElasticPolicyResponseBodyData
	SetMessage(v string) *GetGatewayElasticPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayElasticPolicyResponseBody
	GetRequestId() *string
}

type GetGatewayElasticPolicyResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetGatewayElasticPolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B7F02714-182D-55BC-AF0B-F454364445E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayElasticPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayElasticPolicyResponseBody) GetData() *GetGatewayElasticPolicyResponseBodyData {
	return s.Data
}

func (s *GetGatewayElasticPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayElasticPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayElasticPolicyResponseBody) SetCode(v string) *GetGatewayElasticPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBody) SetData(v *GetGatewayElasticPolicyResponseBodyData) *GetGatewayElasticPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayElasticPolicyResponseBody) SetMessage(v string) *GetGatewayElasticPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBody) SetRequestId(v string) *GetGatewayElasticPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayElasticPolicyResponseBodyData struct {
	ElasticPolicy *GetGatewayElasticPolicyResponseBodyDataElasticPolicy `json:"elasticPolicy,omitempty" xml:"elasticPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	ElasticStrategyId *string `json:"elasticStrategyId,omitempty" xml:"elasticStrategyId,omitempty"`
	// example:
	//
	// CronHPA
	ElasticType *string `json:"elasticType,omitempty" xml:"elasticType,omitempty"`
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s GetGatewayElasticPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponseBodyData) GetElasticPolicy() *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	return s.ElasticPolicy
}

func (s *GetGatewayElasticPolicyResponseBodyData) GetElasticStrategyId() *string {
	return s.ElasticStrategyId
}

func (s *GetGatewayElasticPolicyResponseBodyData) GetElasticType() *string {
	return s.ElasticType
}

func (s *GetGatewayElasticPolicyResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetGatewayElasticPolicyResponseBodyData) SetElasticPolicy(v *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) *GetGatewayElasticPolicyResponseBodyData {
	s.ElasticPolicy = v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyData) SetElasticStrategyId(v string) *GetGatewayElasticPolicyResponseBodyData {
	s.ElasticStrategyId = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyData) SetElasticType(v string) *GetGatewayElasticPolicyResponseBodyData {
	s.ElasticType = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyData) SetGatewayId(v string) *GetGatewayElasticPolicyResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyData) Validate() error {
	if s.ElasticPolicy != nil {
		if err := s.ElasticPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayElasticPolicyResponseBodyDataElasticPolicy struct {
	DisableScaleTimePolicyList []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList `json:"disableScaleTimePolicyList,omitempty" xml:"disableScaleTimePolicyList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	ElasticEnabled *bool `json:"elasticEnabled,omitempty" xml:"elasticEnabled,omitempty"`
	// example:
	//
	// AutoHPA
	ElasticType               *string                                                                          `json:"elasticType,omitempty" xml:"elasticType,omitempty"`
	EnableScaleTimePolicyList []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList `json:"enableScaleTimePolicyList,omitempty" xml:"enableScaleTimePolicyList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	LoadWarningThreshold *bool `json:"loadWarningThreshold,omitempty" xml:"loadWarningThreshold,omitempty"`
	// example:
	//
	// 10
	MaxUnits       *int32                                                                `json:"maxUnits,omitempty" xml:"maxUnits,omitempty"`
	TimePolicyList []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList `json:"timePolicyList,omitempty" xml:"timePolicyList,omitempty" type:"Repeated"`
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetDisableScaleTimePolicyList() []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList {
	return s.DisableScaleTimePolicyList
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetElasticEnabled() *bool {
	return s.ElasticEnabled
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetElasticType() *string {
	return s.ElasticType
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetEnableScaleTimePolicyList() []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList {
	return s.EnableScaleTimePolicyList
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetLoadWarningThreshold() *bool {
	return s.LoadWarningThreshold
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetMaxUnits() *int32 {
	return s.MaxUnits
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) GetTimePolicyList() []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList {
	return s.TimePolicyList
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetDisableScaleTimePolicyList(v []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.DisableScaleTimePolicyList = v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetElasticEnabled(v bool) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.ElasticEnabled = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetElasticType(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.ElasticType = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetEnableScaleTimePolicyList(v []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.EnableScaleTimePolicyList = v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetLoadWarningThreshold(v bool) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.LoadWarningThreshold = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetMaxUnits(v int32) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.MaxUnits = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) SetTimePolicyList(v []*GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) *GetGatewayElasticPolicyResponseBodyDataElasticPolicy {
	s.TimePolicyList = v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicy) Validate() error {
	if s.DisableScaleTimePolicyList != nil {
		for _, item := range s.DisableScaleTimePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EnableScaleTimePolicyList != nil {
		for _, item := range s.EnableScaleTimePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TimePolicyList != nil {
		for _, item := range s.TimePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList struct {
	// example:
	//
	// 00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) SetEndTime(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) SetStartTime(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyDisableScaleTimePolicyList) Validate() error {
	return dara.Validate(s)
}

type GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList struct {
	// example:
	//
	// 18:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 09:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) SetEndTime(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) SetStartTime(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyEnableScaleTimePolicyList) Validate() error {
	return dara.Validate(s)
}

type GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList struct {
	// example:
	//
	// 06:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 4
	Units *int64 `json:"units,omitempty" xml:"units,omitempty"`
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) GetUnits() *int64 {
	return s.Units
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) SetEndTime(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) SetStartTime(v string) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) SetUnits(v int64) *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList {
	s.Units = &v
	return s
}

func (s *GetGatewayElasticPolicyResponseBodyDataElasticPolicyTimePolicyList) Validate() error {
	return dara.Validate(s)
}
