// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualRunMailTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ManualRunMailTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *ManualRunMailTaskResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ManualRunMailTaskResponseBody
	GetSuccess() *bool
}

type ManualRunMailTaskResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: The request succeeded
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ManualRunMailTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManualRunMailTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ManualRunMailTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManualRunMailTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ManualRunMailTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ManualRunMailTaskResponseBody) SetRequestId(v string) *ManualRunMailTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManualRunMailTaskResponseBody) SetResult(v bool) *ManualRunMailTaskResponseBody {
	s.Result = &v
	return s
}

func (s *ManualRunMailTaskResponseBody) SetSuccess(v bool) *ManualRunMailTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ManualRunMailTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
