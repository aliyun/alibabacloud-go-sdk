// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRecordLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordLogs(v []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs) *ListRepoBuildRecordLogResponseBody
	GetBuildRecordLogs() []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs
	SetCode(v string) *ListRepoBuildRecordLogResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoBuildRecordLogResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoBuildRecordLogResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoBuildRecordLogResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoBuildRecordLogResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRepoBuildRecordLogResponseBody
	GetTotalCount() *string
}

type ListRepoBuildRecordLogResponseBody struct {
	// The log content of the image building record.
	BuildRecordLogs []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs `json:"BuildRecordLogs,omitempty" xml:"BuildRecordLogs,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1000
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoBuildRecordLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogResponseBody) GetBuildRecordLogs() []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	return s.BuildRecordLogs
}

func (s *ListRepoBuildRecordLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoBuildRecordLogResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoBuildRecordLogResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoBuildRecordLogResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoBuildRecordLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoBuildRecordLogResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRepoBuildRecordLogResponseBody) SetBuildRecordLogs(v []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs) *ListRepoBuildRecordLogResponseBody {
	s.BuildRecordLogs = v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetCode(v string) *ListRepoBuildRecordLogResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetIsSuccess(v bool) *ListRepoBuildRecordLogResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetPageNo(v int32) *ListRepoBuildRecordLogResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetPageSize(v int32) *ListRepoBuildRecordLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetRequestId(v string) *ListRepoBuildRecordLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetTotalCount(v string) *ListRepoBuildRecordLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) Validate() error {
	if s.BuildRecordLogs != nil {
		for _, item := range s.BuildRecordLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepoBuildRecordLogResponseBodyBuildRecordLogs struct {
	// The stage of the building that is recorded in the log entry.
	//
	// example:
	//
	// GIT_CLONE
	BuildStage *string `json:"BuildStage,omitempty" xml:"BuildStage,omitempty"`
	// The line number of the log entry.
	//
	// example:
	//
	// 2
	LineNumber *int32 `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	// The content of the log.
	//
	// example:
	//
	// fetch stage begin
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListRepoBuildRecordLogResponseBodyBuildRecordLogs) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordLogResponseBodyBuildRecordLogs) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) GetBuildStage() *string {
	return s.BuildStage
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) GetLineNumber() *int32 {
	return s.LineNumber
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) GetMessage() *string {
	return s.Message
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) SetBuildStage(v string) *ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	s.BuildStage = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) SetLineNumber(v int32) *ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	s.LineNumber = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) SetMessage(v string) *ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	s.Message = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) Validate() error {
	return dara.Validate(s)
}
