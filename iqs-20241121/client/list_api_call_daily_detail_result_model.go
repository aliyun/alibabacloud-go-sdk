// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiCallDailyDetailResult interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v []*ListApiCallDailyDetailResultDetails) *ListApiCallDailyDetailResult
	GetDetails() []*ListApiCallDailyDetailResultDetails
}

type ListApiCallDailyDetailResult struct {
	Details []*ListApiCallDailyDetailResultDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s ListApiCallDailyDetailResult) String() string {
	return dara.Prettify(s)
}

func (s ListApiCallDailyDetailResult) GoString() string {
	return s.String()
}

func (s *ListApiCallDailyDetailResult) GetDetails() []*ListApiCallDailyDetailResultDetails {
	return s.Details
}

func (s *ListApiCallDailyDetailResult) SetDetails(v []*ListApiCallDailyDetailResultDetails) *ListApiCallDailyDetailResult {
	s.Details = v
	return s
}

func (s *ListApiCallDailyDetailResult) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiCallDailyDetailResultDetails struct {
	ApiName       *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	AvgLatency    *int32  `json:"avgLatency,omitempty" xml:"avgLatency,omitempty"`
	BillDate      *string `json:"billDate,omitempty" xml:"billDate,omitempty"`
	EngineType    *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	FailedCount   *int32  `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	P90Latency    *int32  `json:"p90Latency,omitempty" xml:"p90Latency,omitempty"`
	P95Latency    *int32  `json:"p95Latency,omitempty" xml:"p95Latency,omitempty"`
	RequestMaxQps *int32  `json:"requestMaxQps,omitempty" xml:"requestMaxQps,omitempty"`
	SubAccountId  *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	SuccessCount  *int32  `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s ListApiCallDailyDetailResultDetails) String() string {
	return dara.Prettify(s)
}

func (s ListApiCallDailyDetailResultDetails) GoString() string {
	return s.String()
}

func (s *ListApiCallDailyDetailResultDetails) GetApiName() *string {
	return s.ApiName
}

func (s *ListApiCallDailyDetailResultDetails) GetAvgLatency() *int32 {
	return s.AvgLatency
}

func (s *ListApiCallDailyDetailResultDetails) GetBillDate() *string {
	return s.BillDate
}

func (s *ListApiCallDailyDetailResultDetails) GetEngineType() *string {
	return s.EngineType
}

func (s *ListApiCallDailyDetailResultDetails) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *ListApiCallDailyDetailResultDetails) GetP90Latency() *int32 {
	return s.P90Latency
}

func (s *ListApiCallDailyDetailResultDetails) GetP95Latency() *int32 {
	return s.P95Latency
}

func (s *ListApiCallDailyDetailResultDetails) GetRequestMaxQps() *int32 {
	return s.RequestMaxQps
}

func (s *ListApiCallDailyDetailResultDetails) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *ListApiCallDailyDetailResultDetails) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *ListApiCallDailyDetailResultDetails) SetApiName(v string) *ListApiCallDailyDetailResultDetails {
	s.ApiName = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetAvgLatency(v int32) *ListApiCallDailyDetailResultDetails {
	s.AvgLatency = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetBillDate(v string) *ListApiCallDailyDetailResultDetails {
	s.BillDate = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetEngineType(v string) *ListApiCallDailyDetailResultDetails {
	s.EngineType = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetFailedCount(v int32) *ListApiCallDailyDetailResultDetails {
	s.FailedCount = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetP90Latency(v int32) *ListApiCallDailyDetailResultDetails {
	s.P90Latency = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetP95Latency(v int32) *ListApiCallDailyDetailResultDetails {
	s.P95Latency = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetRequestMaxQps(v int32) *ListApiCallDailyDetailResultDetails {
	s.RequestMaxQps = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetSubAccountId(v string) *ListApiCallDailyDetailResultDetails {
	s.SubAccountId = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) SetSuccessCount(v int32) *ListApiCallDailyDetailResultDetails {
	s.SuccessCount = &v
	return s
}

func (s *ListApiCallDailyDetailResultDetails) Validate() error {
	return dara.Validate(s)
}
