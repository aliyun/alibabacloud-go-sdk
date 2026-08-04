// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayAndPauseControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PlayAndPauseControlResponseBody
	GetCode() *int32
	SetMessage(v string) *PlayAndPauseControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *PlayAndPauseControlResponseBody
	GetRequestId() *string
	SetResult(v bool) *PlayAndPauseControlResponseBody
	GetResult() *bool
	SetSuccess(v string) *PlayAndPauseControlResponseBody
	GetSuccess() *string
}

type PlayAndPauseControlResponseBody struct {
	// Return code of the invocation
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information. In typical scenarios, this provides a brief description of a failed invocation to help the caller troubleshoot the issue.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10002398812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Actual return result from the service
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the invocation succeeded. The value true means success, and false means failure. If the invocation fails, check the Message field for details.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PlayAndPauseControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlResponseBody) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PlayAndPauseControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PlayAndPauseControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PlayAndPauseControlResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PlayAndPauseControlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PlayAndPauseControlResponseBody) SetCode(v int32) *PlayAndPauseControlResponseBody {
	s.Code = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetMessage(v string) *PlayAndPauseControlResponseBody {
	s.Message = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetRequestId(v string) *PlayAndPauseControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetResult(v bool) *PlayAndPauseControlResponseBody {
	s.Result = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) SetSuccess(v string) *PlayAndPauseControlResponseBody {
	s.Success = &v
	return s
}

func (s *PlayAndPauseControlResponseBody) Validate() error {
	return dara.Validate(s)
}
