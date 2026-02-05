// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScheduledTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScheduledTaskResponseBody
	GetSuccess() *bool
}

type DeleteScheduledTaskResponseBody struct {
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

func (s DeleteScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduledTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScheduledTaskResponseBody) SetMessage(v string) *DeleteScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) SetRequestId(v string) *DeleteScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) SetSuccess(v bool) *DeleteScheduledTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
