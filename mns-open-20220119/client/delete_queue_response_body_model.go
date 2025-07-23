// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteQueueResponseBody
	GetCode() *int64
	SetData(v *DeleteQueueResponseBodyData) *DeleteQueueResponseBody
	GetData() *DeleteQueueResponseBodyData
	SetMessage(v string) *DeleteQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQueueResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteQueueResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteQueueResponseBody
	GetSuccess() *bool
}

type DeleteQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteQueueResponseBody) GetData() *DeleteQueueResponseBodyData {
	return s.Data
}

func (s *DeleteQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQueueResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQueueResponseBody) SetCode(v int64) *DeleteQueueResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBody) SetData(v *DeleteQueueResponseBodyData) *DeleteQueueResponseBody {
	s.Data = v
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

func (s *DeleteQueueResponseBody) SetStatus(v string) *DeleteQueueResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteQueueResponseBody) SetSuccess(v bool) *DeleteQueueResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQueueResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteQueueResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *DeleteQueueResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteQueueResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQueueResponseBodyData) SetCode(v int64) *DeleteQueueResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBodyData) SetMessage(v string) *DeleteQueueResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteQueueResponseBodyData) SetSuccess(v bool) *DeleteQueueResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteQueueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
