// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyzedRequestCount(v int64) *DescribeAiEngineOverviewResponseBody
	GetAnalyzedRequestCount() *int64
	SetEventCount(v int64) *DescribeAiEngineOverviewResponseBody
	GetEventCount() *int64
	SetHandleStatusCounts(v *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) *DescribeAiEngineOverviewResponseBody
	GetHandleStatusCounts() *DescribeAiEngineOverviewResponseBodyHandleStatusCounts
	SetProtectedGroupCount(v int32) *DescribeAiEngineOverviewResponseBody
	GetProtectedGroupCount() *int32
	SetProtectedResourceCount(v int32) *DescribeAiEngineOverviewResponseBody
	GetProtectedResourceCount() *int32
	SetRequestId(v string) *DescribeAiEngineOverviewResponseBody
	GetRequestId() *string
	SetRiskCounts(v *DescribeAiEngineOverviewResponseBodyRiskCounts) *DescribeAiEngineOverviewResponseBody
	GetRiskCounts() *DescribeAiEngineOverviewResponseBodyRiskCounts
}

type DescribeAiEngineOverviewResponseBody struct {
	// The number of analyzed requests.
	//
	// example:
	//
	// 1000000
	AnalyzedRequestCount *int64 `json:"AnalyzedRequestCount,omitempty" xml:"AnalyzedRequestCount,omitempty"`
	// The number of events within the specified time range.
	//
	// example:
	//
	// 1
	EventCount *int64 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The number of events within the specified time range, categorized by handling status.
	HandleStatusCounts *DescribeAiEngineOverviewResponseBodyHandleStatusCounts `json:"HandleStatusCounts,omitempty" xml:"HandleStatusCounts,omitempty" type:"Struct"`
	// The number of protection groups for which AI Engine is enabled.
	//
	// example:
	//
	// 2
	ProtectedGroupCount *int32 `json:"ProtectedGroupCount,omitempty" xml:"ProtectedGroupCount,omitempty"`
	// The number of protected objects for which AI Engine is enabled. This count does not include protected objects that belong to a protection group.
	//
	// example:
	//
	// 9
	ProtectedResourceCount *int32 `json:"ProtectedResourceCount,omitempty" xml:"ProtectedResourceCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01A1CEE****-019A-8B44-255****F9C2C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of events within the specified time range, categorized by risk level.
	RiskCounts *DescribeAiEngineOverviewResponseBodyRiskCounts `json:"RiskCounts,omitempty" xml:"RiskCounts,omitempty" type:"Struct"`
}

func (s DescribeAiEngineOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineOverviewResponseBody) GetAnalyzedRequestCount() *int64 {
	return s.AnalyzedRequestCount
}

func (s *DescribeAiEngineOverviewResponseBody) GetEventCount() *int64 {
	return s.EventCount
}

func (s *DescribeAiEngineOverviewResponseBody) GetHandleStatusCounts() *DescribeAiEngineOverviewResponseBodyHandleStatusCounts {
	return s.HandleStatusCounts
}

func (s *DescribeAiEngineOverviewResponseBody) GetProtectedGroupCount() *int32 {
	return s.ProtectedGroupCount
}

func (s *DescribeAiEngineOverviewResponseBody) GetProtectedResourceCount() *int32 {
	return s.ProtectedResourceCount
}

func (s *DescribeAiEngineOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineOverviewResponseBody) GetRiskCounts() *DescribeAiEngineOverviewResponseBodyRiskCounts {
	return s.RiskCounts
}

func (s *DescribeAiEngineOverviewResponseBody) SetAnalyzedRequestCount(v int64) *DescribeAiEngineOverviewResponseBody {
	s.AnalyzedRequestCount = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) SetEventCount(v int64) *DescribeAiEngineOverviewResponseBody {
	s.EventCount = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) SetHandleStatusCounts(v *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) *DescribeAiEngineOverviewResponseBody {
	s.HandleStatusCounts = v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) SetProtectedGroupCount(v int32) *DescribeAiEngineOverviewResponseBody {
	s.ProtectedGroupCount = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) SetProtectedResourceCount(v int32) *DescribeAiEngineOverviewResponseBody {
	s.ProtectedResourceCount = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) SetRequestId(v string) *DescribeAiEngineOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) SetRiskCounts(v *DescribeAiEngineOverviewResponseBodyRiskCounts) *DescribeAiEngineOverviewResponseBody {
	s.RiskCounts = v
	return s
}

func (s *DescribeAiEngineOverviewResponseBody) Validate() error {
	if s.HandleStatusCounts != nil {
		if err := s.HandleStatusCounts.Validate(); err != nil {
			return err
		}
	}
	if s.RiskCounts != nil {
		if err := s.RiskCounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAiEngineOverviewResponseBodyHandleStatusCounts struct {
	// The number of handled events.
	//
	// example:
	//
	// 0
	Handled *int64 `json:"Handled,omitempty" xml:"Handled,omitempty"`
	// The number of events being handled.
	//
	// example:
	//
	// 0
	Handling *int64 `json:"Handling,omitempty" xml:"Handling,omitempty"`
	// The number of unhandled events.
	//
	// example:
	//
	// 1
	Unhandled *int64 `json:"Unhandled,omitempty" xml:"Unhandled,omitempty"`
}

func (s DescribeAiEngineOverviewResponseBodyHandleStatusCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineOverviewResponseBodyHandleStatusCounts) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) GetHandled() *int64 {
	return s.Handled
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) GetHandling() *int64 {
	return s.Handling
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) GetUnhandled() *int64 {
	return s.Unhandled
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) SetHandled(v int64) *DescribeAiEngineOverviewResponseBodyHandleStatusCounts {
	s.Handled = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) SetHandling(v int64) *DescribeAiEngineOverviewResponseBodyHandleStatusCounts {
	s.Handling = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) SetUnhandled(v int64) *DescribeAiEngineOverviewResponseBodyHandleStatusCounts {
	s.Unhandled = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyHandleStatusCounts) Validate() error {
	return dara.Validate(s)
}

type DescribeAiEngineOverviewResponseBodyRiskCounts struct {
	// The number of critical events.
	//
	// example:
	//
	// 0
	Critical *int64 `json:"Critical,omitempty" xml:"Critical,omitempty"`
	// The number of high-risk events.
	//
	// example:
	//
	// 1
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of low-risk events.
	//
	// example:
	//
	// 0
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-risk events.
	//
	// example:
	//
	// 0
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s DescribeAiEngineOverviewResponseBodyRiskCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineOverviewResponseBodyRiskCounts) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) GetCritical() *int64 {
	return s.Critical
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) GetHigh() *int64 {
	return s.High
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) GetLow() *int64 {
	return s.Low
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) GetMedium() *int64 {
	return s.Medium
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) SetCritical(v int64) *DescribeAiEngineOverviewResponseBodyRiskCounts {
	s.Critical = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) SetHigh(v int64) *DescribeAiEngineOverviewResponseBodyRiskCounts {
	s.High = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) SetLow(v int64) *DescribeAiEngineOverviewResponseBodyRiskCounts {
	s.Low = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) SetMedium(v int64) *DescribeAiEngineOverviewResponseBodyRiskCounts {
	s.Medium = &v
	return s
}

func (s *DescribeAiEngineOverviewResponseBodyRiskCounts) Validate() error {
	return dara.Validate(s)
}
