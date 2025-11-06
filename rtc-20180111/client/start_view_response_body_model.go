// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartViewResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartViewResponseBody
	GetTaskId() *string
}

type StartViewResponseBody struct {
	// example:
	//
	// FA2F9DE9-8EAD-580E-87DF-A3D25EE87C37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartViewResponseBody) GoString() string {
	return s.String()
}

func (s *StartViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartViewResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartViewResponseBody) SetRequestId(v string) *StartViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartViewResponseBody) SetTaskId(v string) *StartViewResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartViewResponseBody) Validate() error {
	return dara.Validate(s)
}
