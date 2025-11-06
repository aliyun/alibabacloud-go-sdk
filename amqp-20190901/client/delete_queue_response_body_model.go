// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteQueueResponseBody
	GetCode() *int32
	SetData(v string) *DeleteQueueResponseBody
	GetData() *string
	SetMessage(v string) *DeleteQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQueueResponseBody
	GetSuccess() *bool
}

type DeleteQueueResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteQueueResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQueueResponseBody) SetCode(v int32) *DeleteQueueResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBody) SetData(v string) *DeleteQueueResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQueueResponseBody) SetMessage(v string) *DeleteQueueResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueueResponseBody) SetSuccess(v bool) *DeleteQueueResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
