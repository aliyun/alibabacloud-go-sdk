// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebugExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *StartDebugExecutionResponseBody
	GetExecutionId() *string
	SetRequestId(v string) *StartDebugExecutionResponseBody
	GetRequestId() *string
}

type StartDebugExecutionResponseBody struct {
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// example:
	//
	// A5320F1D-92D9-44BB-A416-5FC525ED6D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDebugExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDebugExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartDebugExecutionResponseBody) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *StartDebugExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDebugExecutionResponseBody) SetExecutionId(v string) *StartDebugExecutionResponseBody {
	s.ExecutionId = &v
	return s
}

func (s *StartDebugExecutionResponseBody) SetRequestId(v string) *StartDebugExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDebugExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
