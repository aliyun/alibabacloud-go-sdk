// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSenderStatisticsByTagNameAndBatchIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SenderStatisticsByTagNameAndBatchIDResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *SenderStatisticsByTagNameAndBatchIDResponseBody
	GetTotalCount() *int32
	SetData(v *SenderStatisticsByTagNameAndBatchIDResponseBodyData) *SenderStatisticsByTagNameAndBatchIDResponseBody
	GetData() *SenderStatisticsByTagNameAndBatchIDResponseBodyData
}

type SenderStatisticsByTagNameAndBatchIDResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Data records
	Data *SenderStatisticsByTagNameAndBatchIDResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) GetData() *SenderStatisticsByTagNameAndBatchIDResponseBodyData {
	return s.Data
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetRequestId(v string) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetTotalCount(v int32) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetData(v *SenderStatisticsByTagNameAndBatchIDResponseBodyData) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.Data = v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SenderStatisticsByTagNameAndBatchIDResponseBodyData struct {
	Stat []*SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat `json:"stat,omitempty" xml:"stat,omitempty" type:"Repeated"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyData) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyData) GetStat() []*SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	return s.Stat
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyData) SetStat(v []*SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) *SenderStatisticsByTagNameAndBatchIDResponseBodyData {
	s.Stat = v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyData) Validate() error {
	if s.Stat != nil {
		for _, item := range s.Stat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat struct {
	// Creation time
	//
	// example:
	//
	// 2021-07-02
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Failure count
	//
	// example:
	//
	// 0
	FaildCount *string `json:"faildCount,omitempty" xml:"faildCount,omitempty"`
	// Request count
	//
	// example:
	//
	// 4
	RequestCount *string `json:"requestCount,omitempty" xml:"requestCount,omitempty"`
	// Success rate
	//
	// example:
	//
	// 100.00%
	SucceededPercent *string `json:"succeededPercent,omitempty" xml:"succeededPercent,omitempty"`
	// Success count
	//
	// example:
	//
	// 4
	SuccessCount *string `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// Invalid count
	//
	// example:
	//
	// 0
	UnavailableCount *string `json:"unavailableCount,omitempty" xml:"unavailableCount,omitempty"`
	// Unavailability rate
	//
	// example:
	//
	// 0%
	UnavailablePercent *string `json:"unavailablePercent,omitempty" xml:"unavailablePercent,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetFaildCount() *string {
	return s.FaildCount
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetRequestCount() *string {
	return s.RequestCount
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetSucceededPercent() *string {
	return s.SucceededPercent
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetSuccessCount() *string {
	return s.SuccessCount
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetUnavailableCount() *string {
	return s.UnavailableCount
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GetUnavailablePercent() *string {
	return s.UnavailablePercent
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetCreateTime(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetFaildCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.FaildCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetRequestCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.RequestCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetSucceededPercent(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.SucceededPercent = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetSuccessCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.SuccessCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetUnavailableCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.UnavailableCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetUnavailablePercent(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.UnavailablePercent = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) Validate() error {
	return dara.Validate(s)
}
