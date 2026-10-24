// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRuleStrategyStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetClientRuleStrategyStatisticsResponseBodyData) *GetClientRuleStrategyStatisticsResponseBody
	GetData() *GetClientRuleStrategyStatisticsResponseBodyData
	SetRequestId(v string) *GetClientRuleStrategyStatisticsResponseBody
	GetRequestId() *string
}

type GetClientRuleStrategyStatisticsResponseBody struct {
	// The statistics of proactive defense policies.
	Data *GetClientRuleStrategyStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 01A157FA-F957-3A5A-A6B8-5BBB8ADB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientRuleStrategyStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientRuleStrategyStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientRuleStrategyStatisticsResponseBody) GetData() *GetClientRuleStrategyStatisticsResponseBodyData {
	return s.Data
}

func (s *GetClientRuleStrategyStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientRuleStrategyStatisticsResponseBody) SetData(v *GetClientRuleStrategyStatisticsResponseBodyData) *GetClientRuleStrategyStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponseBody) SetRequestId(v string) *GetClientRuleStrategyStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClientRuleStrategyStatisticsResponseBodyData struct {
	// The number of custom rule policies for proactive defense.
	//
	// example:
	//
	// 2
	CustomRuleCount *int32 `json:"CustomRuleCount,omitempty" xml:"CustomRuleCount,omitempty"`
	// The number of system rule policies for proactive defense.
	//
	// example:
	//
	// 10
	SystemRuleCount *int32 `json:"SystemRuleCount,omitempty" xml:"SystemRuleCount,omitempty"`
}

func (s GetClientRuleStrategyStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClientRuleStrategyStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClientRuleStrategyStatisticsResponseBodyData) GetCustomRuleCount() *int32 {
	return s.CustomRuleCount
}

func (s *GetClientRuleStrategyStatisticsResponseBodyData) GetSystemRuleCount() *int32 {
	return s.SystemRuleCount
}

func (s *GetClientRuleStrategyStatisticsResponseBodyData) SetCustomRuleCount(v int32) *GetClientRuleStrategyStatisticsResponseBodyData {
	s.CustomRuleCount = &v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponseBodyData) SetSystemRuleCount(v int32) *GetClientRuleStrategyStatisticsResponseBodyData {
	s.SystemRuleCount = &v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
