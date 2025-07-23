// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UnsubscribeResponseBody
	GetCode() *int64
	SetData(v *UnsubscribeResponseBodyData) *UnsubscribeResponseBody
	GetData() *UnsubscribeResponseBodyData
	SetMessage(v string) *UnsubscribeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnsubscribeResponseBody
	GetRequestId() *string
	SetStatus(v string) *UnsubscribeResponseBody
	GetStatus() *string
	SetSuccess(v bool) *UnsubscribeResponseBody
	GetSuccess() *bool
}

type UnsubscribeResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UnsubscribeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UnsubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UnsubscribeResponseBody) GetData() *UnsubscribeResponseBodyData {
	return s.Data
}

func (s *UnsubscribeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnsubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnsubscribeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UnsubscribeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnsubscribeResponseBody) SetCode(v int64) *UnsubscribeResponseBody {
	s.Code = &v
	return s
}

func (s *UnsubscribeResponseBody) SetData(v *UnsubscribeResponseBodyData) *UnsubscribeResponseBody {
	s.Data = v
	return s
}

func (s *UnsubscribeResponseBody) SetMessage(v string) *UnsubscribeResponseBody {
	s.Message = &v
	return s
}

func (s *UnsubscribeResponseBody) SetRequestId(v string) *UnsubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeResponseBody) SetStatus(v string) *UnsubscribeResponseBody {
	s.Status = &v
	return s
}

func (s *UnsubscribeResponseBody) SetSuccess(v bool) *UnsubscribeResponseBody {
	s.Success = &v
	return s
}

func (s *UnsubscribeResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnsubscribeResponseBodyData struct {
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

func (s UnsubscribeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *UnsubscribeResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UnsubscribeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UnsubscribeResponseBodyData) SetCode(v int64) *UnsubscribeResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnsubscribeResponseBodyData) SetMessage(v string) *UnsubscribeResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnsubscribeResponseBodyData) SetSuccess(v bool) *UnsubscribeResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnsubscribeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
