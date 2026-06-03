// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskLevelDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRiskLevelDistributionResponseBody
	GetRequestId() *string
	SetTimeList(v []*GetRiskLevelDistributionResponseBodyTimeList) *GetRiskLevelDistributionResponseBody
	GetTimeList() []*GetRiskLevelDistributionResponseBodyTimeList
}

type GetRiskLevelDistributionResponseBody struct {
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetRiskLevelDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetRiskLevelDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRiskLevelDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRiskLevelDistributionResponseBody) GetTimeList() []*GetRiskLevelDistributionResponseBodyTimeList {
	return s.TimeList
}

func (s *GetRiskLevelDistributionResponseBody) SetRequestId(v string) *GetRiskLevelDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBody) SetTimeList(v []*GetRiskLevelDistributionResponseBodyTimeList) *GetRiskLevelDistributionResponseBody {
	s.TimeList = v
	return s
}

func (s *GetRiskLevelDistributionResponseBody) Validate() error {
	if s.TimeList != nil {
		for _, item := range s.TimeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRiskLevelDistributionResponseBodyTimeList struct {
	// example:
	//
	// 2019-06-06 01:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 2019-06-06 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 200
	HighRiskCount *int64 `json:"HighRiskCount,omitempty" xml:"HighRiskCount,omitempty"`
	// example:
	//
	// 500
	LowRiskCount *int64 `json:"LowRiskCount,omitempty" xml:"LowRiskCount,omitempty"`
	// example:
	//
	// 300
	MiddleRiskCount *int64 `json:"MiddleRiskCount,omitempty" xml:"MiddleRiskCount,omitempty"`
	// example:
	//
	// 1000
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
}

func (s GetRiskLevelDistributionResponseBodyTimeList) String() string {
	return dara.Prettify(s)
}

func (s GetRiskLevelDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) GetEndDate() *string {
	return s.EndDate
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) GetHighRiskCount() *int64 {
	return s.HighRiskCount
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) GetLowRiskCount() *int64 {
	return s.LowRiskCount
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) GetMiddleRiskCount() *int64 {
	return s.MiddleRiskCount
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetBeginDate(v string) *GetRiskLevelDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetEndDate(v string) *GetRiskLevelDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetHighRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.HighRiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetLowRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.LowRiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetMiddleRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.MiddleRiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.RiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) Validate() error {
	return dara.Validate(s)
}
