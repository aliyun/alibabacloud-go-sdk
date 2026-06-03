// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogTypeDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLogTypeDistributionResponseBody
	GetRequestId() *string
	SetTimeList(v []*GetLogTypeDistributionResponseBodyTimeList) *GetLogTypeDistributionResponseBody
	GetTimeList() []*GetLogTypeDistributionResponseBodyTimeList
}

type GetLogTypeDistributionResponseBody struct {
	// example:
	//
	// 8EC13467-A84C-401F-A4A0-AF893066FBA1
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetLogTypeDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetLogTypeDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogTypeDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogTypeDistributionResponseBody) GetTimeList() []*GetLogTypeDistributionResponseBodyTimeList {
	return s.TimeList
}

func (s *GetLogTypeDistributionResponseBody) SetRequestId(v string) *GetLogTypeDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogTypeDistributionResponseBody) SetTimeList(v []*GetLogTypeDistributionResponseBodyTimeList) *GetLogTypeDistributionResponseBody {
	s.TimeList = v
	return s
}

func (s *GetLogTypeDistributionResponseBody) Validate() error {
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

type GetLogTypeDistributionResponseBodyTimeList struct {
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 10000
	DeleteSqlCount *int64 `json:"DeleteSqlCount,omitempty" xml:"DeleteSqlCount,omitempty"`
	// example:
	//
	// 2019-06-06 01:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 500000
	ExecCostUS *string `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	// example:
	//
	// 10000
	InsertSqlCount *int64 `json:"InsertSqlCount,omitempty" xml:"InsertSqlCount,omitempty"`
	// example:
	//
	// 10000
	OtherSqlCount *int64 `json:"OtherSqlCount,omitempty" xml:"OtherSqlCount,omitempty"`
	// example:
	//
	// 10000
	SelectSqlCount *int64 `json:"SelectSqlCount,omitempty" xml:"SelectSqlCount,omitempty"`
	// example:
	//
	// 50000
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	// example:
	//
	// 10000
	UpdateSqlCount *int64 `json:"UpdateSqlCount,omitempty" xml:"UpdateSqlCount,omitempty"`
}

func (s GetLogTypeDistributionResponseBodyTimeList) String() string {
	return dara.Prettify(s)
}

func (s GetLogTypeDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetDeleteSqlCount() *int64 {
	return s.DeleteSqlCount
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetEndDate() *string {
	return s.EndDate
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetExecCostUS() *string {
	return s.ExecCostUS
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetInsertSqlCount() *int64 {
	return s.InsertSqlCount
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetOtherSqlCount() *int64 {
	return s.OtherSqlCount
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetSelectSqlCount() *int64 {
	return s.SelectSqlCount
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetLogTypeDistributionResponseBodyTimeList) GetUpdateSqlCount() *int64 {
	return s.UpdateSqlCount
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetBeginDate(v string) *GetLogTypeDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetDeleteSqlCount(v int64) *GetLogTypeDistributionResponseBodyTimeList {
	s.DeleteSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetEndDate(v string) *GetLogTypeDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetExecCostUS(v string) *GetLogTypeDistributionResponseBodyTimeList {
	s.ExecCostUS = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetInsertSqlCount(v int64) *GetLogTypeDistributionResponseBodyTimeList {
	s.InsertSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetOtherSqlCount(v int64) *GetLogTypeDistributionResponseBodyTimeList {
	s.OtherSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetSelectSqlCount(v int64) *GetLogTypeDistributionResponseBodyTimeList {
	s.SelectSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetSqlCount(v int64) *GetLogTypeDistributionResponseBodyTimeList {
	s.SqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetUpdateSqlCount(v int64) *GetLogTypeDistributionResponseBodyTimeList {
	s.UpdateSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) Validate() error {
	return dara.Validate(s)
}
