// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactBuildTaskLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildTaskLogs(v []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) *ListArtifactBuildTaskLogResponseBody
	GetBuildTaskLogs() []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs
	SetCode(v string) *ListArtifactBuildTaskLogResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListArtifactBuildTaskLogResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListArtifactBuildTaskLogResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListArtifactBuildTaskLogResponseBody
	GetTotalCount() *int32
}

type ListArtifactBuildTaskLogResponseBody struct {
	// The log entries of the artifact build task.
	BuildTaskLogs []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs `json:"BuildTaskLogs,omitempty" xml:"BuildTaskLogs,omitempty" type:"Repeated"`
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of log entries.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactBuildTaskLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogResponseBody) GetBuildTaskLogs() []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs {
	return s.BuildTaskLogs
}

func (s *ListArtifactBuildTaskLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListArtifactBuildTaskLogResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListArtifactBuildTaskLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactBuildTaskLogResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListArtifactBuildTaskLogResponseBody) SetBuildTaskLogs(v []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) *ListArtifactBuildTaskLogResponseBody {
	s.BuildTaskLogs = v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetCode(v string) *ListArtifactBuildTaskLogResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetIsSuccess(v bool) *ListArtifactBuildTaskLogResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetRequestId(v string) *ListArtifactBuildTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetTotalCount(v int32) *ListArtifactBuildTaskLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListArtifactBuildTaskLogResponseBodyBuildTaskLogs struct {
	// The row number of the log entry.
	//
	// example:
	//
	// 1
	LineNumber *int32 `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	// The log data.
	//
	// example:
	//
	// Start Build
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) GetLineNumber() *int32 {
	return s.LineNumber
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) GetMessage() *string {
	return s.Message
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) SetLineNumber(v int32) *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs {
	s.LineNumber = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) SetMessage(v string) *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs {
	s.Message = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) Validate() error {
	return dara.Validate(s)
}
