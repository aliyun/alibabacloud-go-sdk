// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhmDWResourceGroupStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetLhmDWResourceGroupStatusResponseBody
	GetData() *string
	SetErrCode(v string) *GetLhmDWResourceGroupStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetLhmDWResourceGroupStatusResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetLhmDWResourceGroupStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLhmDWResourceGroupStatusResponseBody
	GetSuccess() *bool
}

type GetLhmDWResourceGroupStatusResponseBody struct {
	// The business data returned by the operation as a string. The specific content varies by operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID. You can use this ID to locate and troubleshoot issues related to this call.
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

func (s GetLhmDWResourceGroupStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLhmDWResourceGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetLhmDWResourceGroupStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetLhmDWResourceGroupStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetLhmDWResourceGroupStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetLhmDWResourceGroupStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLhmDWResourceGroupStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLhmDWResourceGroupStatusResponseBody) SetData(v string) *GetLhmDWResourceGroupStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponseBody) SetErrCode(v string) *GetLhmDWResourceGroupStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponseBody) SetErrMessage(v string) *GetLhmDWResourceGroupStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponseBody) SetRequestId(v string) *GetLhmDWResourceGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponseBody) SetSuccess(v bool) *GetLhmDWResourceGroupStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
