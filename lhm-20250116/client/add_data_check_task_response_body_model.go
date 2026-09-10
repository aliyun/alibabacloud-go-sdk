// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *AddDataCheckTaskResponseBody
	GetData() *int64
	SetErrCode(v string) *AddDataCheckTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddDataCheckTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *AddDataCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataCheckTaskResponseBody
	GetSuccess() *bool
}

type AddDataCheckTaskResponseBody struct {
	// The task ID of the newly created validation task. Use this ID for subsequent operations such as configuring table details, executing validations, or deleting the task.
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
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Use errCode and errMessage to identify the cause.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDataCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataCheckTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddDataCheckTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddDataCheckTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddDataCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataCheckTaskResponseBody) SetData(v int64) *AddDataCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *AddDataCheckTaskResponseBody) SetErrCode(v string) *AddDataCheckTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddDataCheckTaskResponseBody) SetErrMessage(v string) *AddDataCheckTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddDataCheckTaskResponseBody) SetRequestId(v string) *AddDataCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataCheckTaskResponseBody) SetSuccess(v bool) *AddDataCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
