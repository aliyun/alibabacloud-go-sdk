// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhmAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetLhmAgentStatusResponseBody
	GetData() *string
	SetErrCode(v string) *GetLhmAgentStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetLhmAgentStatusResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetLhmAgentStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLhmAgentStatusResponseBody
	GetSuccess() *bool
}

type GetLhmAgentStatusResponseBody struct {
	// The business data returned by the operation as a string. The content varies depending on the operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// The request ID, which is used to locate and troubleshoot issues with the call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetLhmAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLhmAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetLhmAgentStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetLhmAgentStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetLhmAgentStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetLhmAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLhmAgentStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLhmAgentStatusResponseBody) SetData(v string) *GetLhmAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetLhmAgentStatusResponseBody) SetErrCode(v string) *GetLhmAgentStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetLhmAgentStatusResponseBody) SetErrMessage(v string) *GetLhmAgentStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetLhmAgentStatusResponseBody) SetRequestId(v string) *GetLhmAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLhmAgentStatusResponseBody) SetSuccess(v bool) *GetLhmAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetLhmAgentStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
