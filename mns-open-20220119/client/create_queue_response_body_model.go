// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateQueueResponseBody
	GetCode() *int64
	SetData(v *CreateQueueResponseBodyData) *CreateQueueResponseBody
	GetData() *CreateQueueResponseBodyData
	SetMessage(v string) *CreateQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateQueueResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateQueueResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CreateQueueResponseBody
	GetSuccess() *bool
}

type CreateQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 06273500-249F-5863-121D-74D51123E62C
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

func (s CreateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateQueueResponseBody) GetData() *CreateQueueResponseBodyData {
	return s.Data
}

func (s *CreateQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueueResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQueueResponseBody) SetCode(v int64) *CreateQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBody) SetData(v *CreateQueueResponseBodyData) *CreateQueueResponseBody {
	s.Data = v
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

func (s *CreateQueueResponseBody) SetStatus(v string) *CreateQueueResponseBody {
	s.Status = &v
	return s
}

func (s *CreateQueueResponseBody) SetSuccess(v bool) *CreateQueueResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQueueResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateQueueResponseBodyData struct {
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

func (s CreateQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *CreateQueueResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateQueueResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQueueResponseBodyData) SetCode(v int64) *CreateQueueResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBodyData) SetMessage(v string) *CreateQueueResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateQueueResponseBodyData) SetSuccess(v bool) *CreateQueueResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreateQueueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
