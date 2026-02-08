// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeConcurrencyReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRealtimeConcurrencyReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetRealtimeConcurrencyReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRealtimeConcurrencyReportResponseBody
	GetMessage() *string
	SetReportDate(v int64) *GetRealtimeConcurrencyReportResponseBody
	GetReportDate() *int64
	SetReports(v *GetRealtimeConcurrencyReportResponseBodyReports) *GetRealtimeConcurrencyReportResponseBody
	GetReports() *GetRealtimeConcurrencyReportResponseBodyReports
	SetRequestId(v string) *GetRealtimeConcurrencyReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRealtimeConcurrencyReportResponseBody
	GetSuccess() *bool
}

type GetRealtimeConcurrencyReportResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Report generation time
	//
	// example:
	//
	// 1743474900488
	ReportDate *int64 `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// Report details.
	Reports *GetRealtimeConcurrencyReportResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRealtimeConcurrencyReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeConcurrencyReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetReportDate() *int64 {
	return s.ReportDate
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetReports() *GetRealtimeConcurrencyReportResponseBodyReports {
	return s.Reports
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealtimeConcurrencyReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetCode(v string) *GetRealtimeConcurrencyReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetHttpStatusCode(v int32) *GetRealtimeConcurrencyReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetMessage(v string) *GetRealtimeConcurrencyReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetReportDate(v int64) *GetRealtimeConcurrencyReportResponseBody {
	s.ReportDate = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetReports(v *GetRealtimeConcurrencyReportResponseBodyReports) *GetRealtimeConcurrencyReportResponseBody {
	s.Reports = v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetRequestId(v string) *GetRealtimeConcurrencyReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) SetSuccess(v bool) *GetRealtimeConcurrencyReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBody) Validate() error {
	if s.Reports != nil {
		if err := s.Reports.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRealtimeConcurrencyReportResponseBodyReports struct {
	// Collection of report data
	List []*GetRealtimeConcurrencyReportResponseBodyReportsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetRealtimeConcurrencyReportResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeConcurrencyReportResponseBodyReports) GoString() string {
	return s.String()
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) GetList() []*GetRealtimeConcurrencyReportResponseBodyReportsList {
	return s.List
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) SetList(v []*GetRealtimeConcurrencyReportResponseBodyReportsList) *GetRealtimeConcurrencyReportResponseBodyReports {
	s.List = v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) SetPageNumber(v int32) *GetRealtimeConcurrencyReportResponseBodyReports {
	s.PageNumber = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) SetPageSize(v int32) *GetRealtimeConcurrencyReportResponseBodyReports {
	s.PageSize = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) SetTotalCount(v int32) *GetRealtimeConcurrencyReportResponseBodyReports {
	s.TotalCount = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReports) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRealtimeConcurrencyReportResponseBodyReportsList struct {
	// Instance ID
	//
	// example:
	//
	// 85bf7efa-a07c-498a-850e-99a5849b8589
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name
	//
	// example:
	//
	// 智能外呼场景
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Job ID
	//
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job Name.
	//
	// example:
	//
	// 第一个作业组
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// Configured maximum concurrency limit
	//
	// example:
	//
	// 2
	MaxConcurrencyLimit *int64 `json:"MaxConcurrencyLimit,omitempty" xml:"MaxConcurrencyLimit,omitempty"`
	// Minimum Concurrency Limit of the job
	//
	// example:
	//
	// 0
	MinConcurrencyLimit *int64 `json:"MinConcurrencyLimit,omitempty" xml:"MinConcurrencyLimit,omitempty"`
	// Occupied concurrent value
	//
	// example:
	//
	// 1
	OccupiedConcurrencyCount *int64 `json:"OccupiedConcurrencyCount,omitempty" xml:"OccupiedConcurrencyCount,omitempty"`
	// Report generation time
	//
	// example:
	//
	// 1743474900488
	ReportDate *int64 `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
}

func (s GetRealtimeConcurrencyReportResponseBodyReportsList) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeConcurrencyReportResponseBodyReportsList) GoString() string {
	return s.String()
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetMaxConcurrencyLimit() *int64 {
	return s.MaxConcurrencyLimit
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetMinConcurrencyLimit() *int64 {
	return s.MinConcurrencyLimit
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetOccupiedConcurrencyCount() *int64 {
	return s.OccupiedConcurrencyCount
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) GetReportDate() *int64 {
	return s.ReportDate
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetInstanceId(v string) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetInstanceName(v string) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.InstanceName = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetJobGroupId(v string) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.JobGroupId = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetJobGroupName(v string) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.JobGroupName = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetMaxConcurrencyLimit(v int64) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.MaxConcurrencyLimit = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetMinConcurrencyLimit(v int64) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.MinConcurrencyLimit = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetOccupiedConcurrencyCount(v int64) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.OccupiedConcurrencyCount = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) SetReportDate(v int64) *GetRealtimeConcurrencyReportResponseBodyReportsList {
	s.ReportDate = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponseBodyReportsList) Validate() error {
	return dara.Validate(s)
}
