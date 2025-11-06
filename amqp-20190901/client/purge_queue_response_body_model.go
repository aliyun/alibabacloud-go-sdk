// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PurgeQueueResponseBody
	GetCode() *int32
	SetData(v string) *PurgeQueueResponseBody
	GetData() *string
	SetMessage(v string) *PurgeQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *PurgeQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PurgeQueueResponseBody
	GetSuccess() *bool
}

type PurgeQueueResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PurgeQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurgeQueueResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeQueueResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PurgeQueueResponseBody) GetData() *string {
	return s.Data
}

func (s *PurgeQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PurgeQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurgeQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PurgeQueueResponseBody) SetCode(v int32) *PurgeQueueResponseBody {
	s.Code = &v
	return s
}

func (s *PurgeQueueResponseBody) SetData(v string) *PurgeQueueResponseBody {
	s.Data = &v
	return s
}

func (s *PurgeQueueResponseBody) SetMessage(v string) *PurgeQueueResponseBody {
	s.Message = &v
	return s
}

func (s *PurgeQueueResponseBody) SetRequestId(v string) *PurgeQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurgeQueueResponseBody) SetSuccess(v bool) *PurgeQueueResponseBody {
	s.Success = &v
	return s
}

func (s *PurgeQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
