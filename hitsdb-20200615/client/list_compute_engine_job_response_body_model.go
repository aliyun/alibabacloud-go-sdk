// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeEngineJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListComputeEngineJobResponseBody
	GetAccessDeniedDetail() *string
	SetJobList(v []*ListComputeEngineJobResponseBodyJobList) *ListComputeEngineJobResponseBody
	GetJobList() []*ListComputeEngineJobResponseBodyJobList
	SetPageNumber(v int32) *ListComputeEngineJobResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComputeEngineJobResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListComputeEngineJobResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListComputeEngineJobResponseBody
	GetTotal() *int32
}

type ListComputeEngineJobResponseBody struct {
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	JobList            []*ListComputeEngineJobResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	PageNumber         *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total              *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListComputeEngineJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeEngineJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeEngineJobResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListComputeEngineJobResponseBody) GetJobList() []*ListComputeEngineJobResponseBodyJobList {
	return s.JobList
}

func (s *ListComputeEngineJobResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComputeEngineJobResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeEngineJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeEngineJobResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListComputeEngineJobResponseBody) SetAccessDeniedDetail(v string) *ListComputeEngineJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListComputeEngineJobResponseBody) SetJobList(v []*ListComputeEngineJobResponseBodyJobList) *ListComputeEngineJobResponseBody {
	s.JobList = v
	return s
}

func (s *ListComputeEngineJobResponseBody) SetPageNumber(v int32) *ListComputeEngineJobResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListComputeEngineJobResponseBody) SetPageSize(v int32) *ListComputeEngineJobResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListComputeEngineJobResponseBody) SetRequestId(v string) *ListComputeEngineJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeEngineJobResponseBody) SetTotal(v int32) *ListComputeEngineJobResponseBody {
	s.Total = &v
	return s
}

func (s *ListComputeEngineJobResponseBody) Validate() error {
	if s.JobList != nil {
		for _, item := range s.JobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeEngineJobResponseBodyJobList struct {
	AppName    *string                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoint   *string                `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ExtraInfo  map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	FinishTime *int64                 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	JobId      *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason     *string                `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartedAt  *string                `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	State      *string                `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListComputeEngineJobResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeEngineJobResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListComputeEngineJobResponseBodyJobList) GetAppName() *string {
	return s.AppName
}

func (s *ListComputeEngineJobResponseBodyJobList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListComputeEngineJobResponseBodyJobList) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListComputeEngineJobResponseBodyJobList) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *ListComputeEngineJobResponseBodyJobList) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListComputeEngineJobResponseBodyJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListComputeEngineJobResponseBodyJobList) GetMessage() *string {
	return s.Message
}

func (s *ListComputeEngineJobResponseBodyJobList) GetReason() *string {
	return s.Reason
}

func (s *ListComputeEngineJobResponseBodyJobList) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListComputeEngineJobResponseBodyJobList) GetState() *string {
	return s.State
}

func (s *ListComputeEngineJobResponseBodyJobList) SetAppName(v string) *ListComputeEngineJobResponseBodyJobList {
	s.AppName = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetCreateTime(v int64) *ListComputeEngineJobResponseBodyJobList {
	s.CreateTime = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetEndpoint(v string) *ListComputeEngineJobResponseBodyJobList {
	s.Endpoint = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetExtraInfo(v map[string]interface{}) *ListComputeEngineJobResponseBodyJobList {
	s.ExtraInfo = v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetFinishTime(v int64) *ListComputeEngineJobResponseBodyJobList {
	s.FinishTime = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetJobId(v string) *ListComputeEngineJobResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetMessage(v string) *ListComputeEngineJobResponseBodyJobList {
	s.Message = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetReason(v string) *ListComputeEngineJobResponseBodyJobList {
	s.Reason = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetStartedAt(v string) *ListComputeEngineJobResponseBodyJobList {
	s.StartedAt = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) SetState(v string) *ListComputeEngineJobResponseBodyJobList {
	s.State = &v
	return s
}

func (s *ListComputeEngineJobResponseBodyJobList) Validate() error {
	return dara.Validate(s)
}
