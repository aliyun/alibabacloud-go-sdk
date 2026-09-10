// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationWorkflowSubmitStartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *GetBwmMigrationWorkflowSubmitStartResponseBody
	GetData() *int64
	SetErrCode(v string) *GetBwmMigrationWorkflowSubmitStartResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetBwmMigrationWorkflowSubmitStartResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetBwmMigrationWorkflowSubmitStartResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBwmMigrationWorkflowSubmitStartResponseBody
	GetSuccess() *bool
}

type GetBwmMigrationWorkflowSubmitStartResponseBody struct {
	// The response data.
	//
	// example:
	//
	// 100
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, troubleshoot the issue based on errCode and errMessage.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBwmMigrationWorkflowSubmitStartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationWorkflowSubmitStartResponseBody) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) GetData() *int64 {
	return s.Data
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) SetData(v int64) *GetBwmMigrationWorkflowSubmitStartResponseBody {
	s.Data = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) SetErrCode(v string) *GetBwmMigrationWorkflowSubmitStartResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) SetErrMessage(v string) *GetBwmMigrationWorkflowSubmitStartResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) SetRequestId(v string) *GetBwmMigrationWorkflowSubmitStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) SetSuccess(v bool) *GetBwmMigrationWorkflowSubmitStartResponseBody {
	s.Success = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponseBody) Validate() error {
	return dara.Validate(s)
}
