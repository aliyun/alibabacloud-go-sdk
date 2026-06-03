// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditCountDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAuditCountDistributionResponseBody
	GetRequestId() *string
	SetTimeList(v []*GetAuditCountDistributionResponseBodyTimeList) *GetAuditCountDistributionResponseBody
	GetTimeList() []*GetAuditCountDistributionResponseBodyTimeList
}

type GetAuditCountDistributionResponseBody struct {
	// example:
	//
	// 482EF142-BFA5-43FF-B4B0-84A4B0763639
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetAuditCountDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetAuditCountDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditCountDistributionResponseBody) GetTimeList() []*GetAuditCountDistributionResponseBodyTimeList {
	return s.TimeList
}

func (s *GetAuditCountDistributionResponseBody) SetRequestId(v string) *GetAuditCountDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditCountDistributionResponseBody) SetTimeList(v []*GetAuditCountDistributionResponseBodyTimeList) *GetAuditCountDistributionResponseBody {
	s.TimeList = v
	return s
}

func (s *GetAuditCountDistributionResponseBody) Validate() error {
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

type GetAuditCountDistributionResponseBodyTimeList struct {
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 2019-06-06 01:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 1000
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// 2000
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 100000
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
}

func (s GetAuditCountDistributionResponseBodyTimeList) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionResponseBodyTimeList) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetAuditCountDistributionResponseBodyTimeList) GetEndDate() *string {
	return s.EndDate
}

func (s *GetAuditCountDistributionResponseBodyTimeList) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *GetAuditCountDistributionResponseBodyTimeList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *GetAuditCountDistributionResponseBodyTimeList) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetBeginDate(v string) *GetAuditCountDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetEndDate(v string) *GetAuditCountDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetRiskCount(v int64) *GetAuditCountDistributionResponseBodyTimeList {
	s.RiskCount = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetSessionCount(v int64) *GetAuditCountDistributionResponseBodyTimeList {
	s.SessionCount = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetSqlCount(v int64) *GetAuditCountDistributionResponseBodyTimeList {
	s.SqlCount = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) Validate() error {
	return dara.Validate(s)
}
