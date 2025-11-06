// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateQueueResponseBody
	GetCode() *int32
	SetData(v bool) *CreateQueueResponseBody
	GetData() *bool
	SetMessage(v string) *CreateQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQueueResponseBody
	GetSuccess() *bool
}

type CreateQueueResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateQueueResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQueueResponseBody) SetCode(v int32) *CreateQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBody) SetData(v bool) *CreateQueueResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQueueResponseBody) SetMessage(v string) *CreateQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueueResponseBody) SetSuccess(v bool) *CreateQueueResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
