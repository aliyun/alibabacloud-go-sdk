// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetScheduledReportsResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *GetScheduledReportsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetScheduledReportsResponseBody
	GetPageSize() *int64
	SetReports(v []*GetScheduledReportsResponseBodyReports) *GetScheduledReportsResponseBody
	GetReports() []*GetScheduledReportsResponseBodyReports
	SetRequestId(v string) *GetScheduledReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScheduledReportsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetScheduledReportsResponseBody
	GetTotalCount() *int64
}

type GetScheduledReportsResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reports  []*GetScheduledReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
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

func (s GetScheduledReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledReportsResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledReportsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledReportsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetScheduledReportsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetScheduledReportsResponseBody) GetReports() []*GetScheduledReportsResponseBodyReports {
	return s.Reports
}

func (s *GetScheduledReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScheduledReportsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetScheduledReportsResponseBody) SetMessage(v string) *GetScheduledReportsResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledReportsResponseBody) SetPageNumber(v int64) *GetScheduledReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetScheduledReportsResponseBody) SetPageSize(v int64) *GetScheduledReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetScheduledReportsResponseBody) SetReports(v []*GetScheduledReportsResponseBodyReports) *GetScheduledReportsResponseBody {
	s.Reports = v
	return s
}

func (s *GetScheduledReportsResponseBody) SetRequestId(v string) *GetScheduledReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledReportsResponseBody) SetSuccess(v bool) *GetScheduledReportsResponseBody {
	s.Success = &v
	return s
}

func (s *GetScheduledReportsResponseBody) SetTotalCount(v int64) *GetScheduledReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetScheduledReportsResponseBody) Validate() error {
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

type GetScheduledReportsResponseBodyReports struct {
	// example:
	//
	// 2025-01-01T22:59:59Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2025-01-01T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2025-01-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 65f0053b-f933-49f5-bf65-4e4593e1****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetScheduledReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *GetScheduledReportsResponseBodyReports) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetScheduledReportsResponseBodyReports) GetEndTime() *string {
	return s.EndTime
}

func (s *GetScheduledReportsResponseBodyReports) GetStartTime() *string {
	return s.StartTime
}

func (s *GetScheduledReportsResponseBodyReports) GetStatus() *string {
	return s.Status
}

func (s *GetScheduledReportsResponseBodyReports) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScheduledReportsResponseBodyReports) SetCreatedTime(v string) *GetScheduledReportsResponseBodyReports {
	s.CreatedTime = &v
	return s
}

func (s *GetScheduledReportsResponseBodyReports) SetEndTime(v string) *GetScheduledReportsResponseBodyReports {
	s.EndTime = &v
	return s
}

func (s *GetScheduledReportsResponseBodyReports) SetStartTime(v string) *GetScheduledReportsResponseBodyReports {
	s.StartTime = &v
	return s
}

func (s *GetScheduledReportsResponseBodyReports) SetStatus(v string) *GetScheduledReportsResponseBodyReports {
	s.Status = &v
	return s
}

func (s *GetScheduledReportsResponseBodyReports) SetTaskId(v string) *GetScheduledReportsResponseBodyReports {
	s.TaskId = &v
	return s
}

func (s *GetScheduledReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}
