// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSymbolicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *SubmitSymbolicResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *SubmitSymbolicResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *SubmitSymbolicResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitSymbolicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSymbolicResponseBody
	GetSuccess() *bool
}

type SubmitSymbolicResponseBody struct {
	// args
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 500
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B3AD0FE4-36EF-1641-90B1-77618166F2ff
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSymbolicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSymbolicResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSymbolicResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *SubmitSymbolicResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SubmitSymbolicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitSymbolicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSymbolicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSymbolicResponseBody) SetArgs(v map[string]interface{}) *SubmitSymbolicResponseBody {
	s.Args = v
	return s
}

func (s *SubmitSymbolicResponseBody) SetErrorCode(v int32) *SubmitSymbolicResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitSymbolicResponseBody) SetMessage(v string) *SubmitSymbolicResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSymbolicResponseBody) SetRequestId(v string) *SubmitSymbolicResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSymbolicResponseBody) SetSuccess(v bool) *SubmitSymbolicResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSymbolicResponseBody) Validate() error {
	return dara.Validate(s)
}
