// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateTopicResponseBody
	GetCode() *int64
	SetData(v *CreateTopicResponseBodyData) *CreateTopicResponseBody
	GetData() *CreateTopicResponseBodyData
	SetMessage(v string) *CreateTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTopicResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateTopicResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CreateTopicResponseBody
	GetSuccess() *bool
}

type CreateTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateTopicResponseBody) GetData() *CreateTopicResponseBodyData {
	return s.Data
}

func (s *CreateTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTopicResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTopicResponseBody) SetCode(v int64) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetData(v *CreateTopicResponseBodyData) *CreateTopicResponseBody {
	s.Data = v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetStatus(v string) *CreateTopicResponseBody {
	s.Status = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTopicResponseBodyData struct {
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

func (s CreateTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *CreateTopicResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateTopicResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTopicResponseBodyData) SetCode(v int64) *CreateTopicResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBodyData) SetMessage(v string) *CreateTopicResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBodyData) SetSuccess(v bool) *CreateTopicResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreateTopicResponseBodyData) Validate() error {
	return dara.Validate(s)
}
