// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProgressControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ProgressControlResponseBody
	GetCode() *int32
	SetMessage(v string) *ProgressControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *ProgressControlResponseBody
	GetRequestId() *string
	SetResult(v bool) *ProgressControlResponseBody
	GetResult() *bool
	SetSuccess(v string) *ProgressControlResponseBody
	GetSuccess() *string
}

type ProgressControlResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10002398812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ProgressControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlResponseBody) GoString() string {
	return s.String()
}

func (s *ProgressControlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ProgressControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProgressControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProgressControlResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ProgressControlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ProgressControlResponseBody) SetCode(v int32) *ProgressControlResponseBody {
	s.Code = &v
	return s
}

func (s *ProgressControlResponseBody) SetMessage(v string) *ProgressControlResponseBody {
	s.Message = &v
	return s
}

func (s *ProgressControlResponseBody) SetRequestId(v string) *ProgressControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProgressControlResponseBody) SetResult(v bool) *ProgressControlResponseBody {
	s.Result = &v
	return s
}

func (s *ProgressControlResponseBody) SetSuccess(v string) *ProgressControlResponseBody {
	s.Success = &v
	return s
}

func (s *ProgressControlResponseBody) Validate() error {
	return dara.Validate(s)
}
