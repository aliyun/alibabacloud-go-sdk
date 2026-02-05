// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyScheduledTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyScheduledTaskResponseBody
	GetSuccess() *bool
}

type ModifyScheduledTaskResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScheduledTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScheduledTaskResponseBody) SetMessage(v string) *ModifyScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetSuccess(v bool) *ModifyScheduledTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
