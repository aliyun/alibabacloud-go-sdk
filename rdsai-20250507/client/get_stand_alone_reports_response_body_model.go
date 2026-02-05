// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandAloneReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetStandAloneReportsResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *GetStandAloneReportsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetStandAloneReportsResponseBody
	GetPageSize() *int64
	SetReports(v []*GetStandAloneReportsResponseBodyReports) *GetStandAloneReportsResponseBody
	GetReports() []*GetStandAloneReportsResponseBodyReports
	SetRequestId(v string) *GetStandAloneReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStandAloneReportsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetStandAloneReportsResponseBody
	GetTotalCount() *int64
}

type GetStandAloneReportsResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reports  []*GetStandAloneReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetStandAloneReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandAloneReportsResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandAloneReportsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandAloneReportsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetStandAloneReportsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetStandAloneReportsResponseBody) GetReports() []*GetStandAloneReportsResponseBodyReports {
	return s.Reports
}

func (s *GetStandAloneReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandAloneReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandAloneReportsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetStandAloneReportsResponseBody) SetMessage(v string) *GetStandAloneReportsResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandAloneReportsResponseBody) SetPageNumber(v int64) *GetStandAloneReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetStandAloneReportsResponseBody) SetPageSize(v int64) *GetStandAloneReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetStandAloneReportsResponseBody) SetReports(v []*GetStandAloneReportsResponseBodyReports) *GetStandAloneReportsResponseBody {
	s.Reports = v
	return s
}

func (s *GetStandAloneReportsResponseBody) SetRequestId(v string) *GetStandAloneReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandAloneReportsResponseBody) SetSuccess(v bool) *GetStandAloneReportsResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandAloneReportsResponseBody) SetTotalCount(v int64) *GetStandAloneReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetStandAloneReportsResponseBody) Validate() error {
	if s.Reports != nil {
		for _, item := range s.Reports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandAloneReportsResponseBodyReports struct {
	// example:
	//
	// 2026-01-22T08:20:31Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2026-01-23T08:20:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2026-01-23T08:00:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0f19210c-7bb8-4e38-a099-f94152df****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetStandAloneReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s GetStandAloneReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *GetStandAloneReportsResponseBodyReports) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetStandAloneReportsResponseBodyReports) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStandAloneReportsResponseBodyReports) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStandAloneReportsResponseBodyReports) GetStatus() *string {
	return s.Status
}

func (s *GetStandAloneReportsResponseBodyReports) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStandAloneReportsResponseBodyReports) SetCreatedTime(v string) *GetStandAloneReportsResponseBodyReports {
	s.CreatedTime = &v
	return s
}

func (s *GetStandAloneReportsResponseBodyReports) SetEndTime(v string) *GetStandAloneReportsResponseBodyReports {
	s.EndTime = &v
	return s
}

func (s *GetStandAloneReportsResponseBodyReports) SetStartTime(v string) *GetStandAloneReportsResponseBodyReports {
	s.StartTime = &v
	return s
}

func (s *GetStandAloneReportsResponseBodyReports) SetStatus(v string) *GetStandAloneReportsResponseBodyReports {
	s.Status = &v
	return s
}

func (s *GetStandAloneReportsResponseBodyReports) SetTaskId(v string) *GetStandAloneReportsResponseBodyReports {
	s.TaskId = &v
	return s
}

func (s *GetStandAloneReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}
