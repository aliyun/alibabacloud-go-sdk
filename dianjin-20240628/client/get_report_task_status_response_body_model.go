// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetReportTaskStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetReportTaskStatusResponseBody
	GetErrorMessage() *string
	SetGmtCreate(v string) *GetReportTaskStatusResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetReportTaskStatusResponseBody
	GetGmtModified() *string
	SetOutRequestNo(v string) *GetReportTaskStatusResponseBody
	GetOutRequestNo() *string
	SetTaskStatus(v string) *GetReportTaskStatusResponseBody
	GetTaskStatus() *string
}

type GetReportTaskStatusResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	OutRequestNo *string `json:"outRequestNo,omitempty" xml:"outRequestNo,omitempty"`
	TaskStatus   *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetReportTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportTaskStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetReportTaskStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetReportTaskStatusResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetReportTaskStatusResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetReportTaskStatusResponseBody) GetOutRequestNo() *string {
	return s.OutRequestNo
}

func (s *GetReportTaskStatusResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetReportTaskStatusResponseBody) SetErrorCode(v string) *GetReportTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetReportTaskStatusResponseBody) SetErrorMessage(v string) *GetReportTaskStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetReportTaskStatusResponseBody) SetGmtCreate(v string) *GetReportTaskStatusResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetReportTaskStatusResponseBody) SetGmtModified(v string) *GetReportTaskStatusResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetReportTaskStatusResponseBody) SetOutRequestNo(v string) *GetReportTaskStatusResponseBody {
	s.OutRequestNo = &v
	return s
}

func (s *GetReportTaskStatusResponseBody) SetTaskStatus(v string) *GetReportTaskStatusResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetReportTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
