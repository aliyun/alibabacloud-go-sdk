// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEngineJobLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetComputeEngineJobLogResponseBody
	GetAccessDeniedDetail() *string
	SetJobId(v string) *GetComputeEngineJobLogResponseBody
	GetJobId() *string
	SetLogs(v []*string) *GetComputeEngineJobLogResponseBody
	GetLogs() []*string
	SetPageNumber(v int32) *GetComputeEngineJobLogResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetComputeEngineJobLogResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetComputeEngineJobLogResponseBody
	GetRequestId() *string
}

type GetComputeEngineJobLogResponseBody struct {
	AccessDeniedDetail *string   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	JobId              *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Logs               []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	PageNumber         *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetComputeEngineJobLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEngineJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeEngineJobLogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetComputeEngineJobLogResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetComputeEngineJobLogResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *GetComputeEngineJobLogResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetComputeEngineJobLogResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetComputeEngineJobLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeEngineJobLogResponseBody) SetAccessDeniedDetail(v string) *GetComputeEngineJobLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetComputeEngineJobLogResponseBody) SetJobId(v string) *GetComputeEngineJobLogResponseBody {
	s.JobId = &v
	return s
}

func (s *GetComputeEngineJobLogResponseBody) SetLogs(v []*string) *GetComputeEngineJobLogResponseBody {
	s.Logs = v
	return s
}

func (s *GetComputeEngineJobLogResponseBody) SetPageNumber(v int32) *GetComputeEngineJobLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetComputeEngineJobLogResponseBody) SetPageSize(v int32) *GetComputeEngineJobLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetComputeEngineJobLogResponseBody) SetRequestId(v string) *GetComputeEngineJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeEngineJobLogResponseBody) Validate() error {
	return dara.Validate(s)
}
