// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopProjectTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopProjectTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopProjectTaskResponseBody
	GetSuccess() *bool
}

type StopProjectTaskResponseBody struct {
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopProjectTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopProjectTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopProjectTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopProjectTaskResponseBody) SetRequestId(v string) *StopProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopProjectTaskResponseBody) SetSuccess(v bool) *StopProjectTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopProjectTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
