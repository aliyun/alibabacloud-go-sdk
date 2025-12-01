// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendCustomEventResponseBody
	GetCode() *string
	SetData(v *SendCustomEventResponseBodyData) *SendCustomEventResponseBody
	GetData() *SendCustomEventResponseBodyData
	SetHttpStatusCode(v int32) *SendCustomEventResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendCustomEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendCustomEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendCustomEventResponseBody
	GetSuccess() *bool
}

type SendCustomEventResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Interface return data.
	Data *SendCustomEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message. When the request is successful, it returns a success message; when the request fails, it returns the reason for the failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 606EB377-155D-5AEB-AC4F-F013444A4C45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - true: Call succeeded.
	//
	// - false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCustomEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendCustomEventResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendCustomEventResponseBody) GetData() *SendCustomEventResponseBodyData {
	return s.Data
}

func (s *SendCustomEventResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendCustomEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendCustomEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendCustomEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendCustomEventResponseBody) SetCode(v string) *SendCustomEventResponseBody {
	s.Code = &v
	return s
}

func (s *SendCustomEventResponseBody) SetData(v *SendCustomEventResponseBodyData) *SendCustomEventResponseBody {
	s.Data = v
	return s
}

func (s *SendCustomEventResponseBody) SetHttpStatusCode(v int32) *SendCustomEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendCustomEventResponseBody) SetMessage(v string) *SendCustomEventResponseBody {
	s.Message = &v
	return s
}

func (s *SendCustomEventResponseBody) SetRequestId(v string) *SendCustomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomEventResponseBody) SetSuccess(v bool) *SendCustomEventResponseBody {
	s.Success = &v
	return s
}

func (s *SendCustomEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendCustomEventResponseBodyData struct {
	// Service UID.
	//
	// example:
	//
	// 1601097845544644
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Customer name.
	//
	// example:
	//
	// Tianjin Ruipengsheng Technology Development Co., Ltd
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// Alert ID.
	//
	// example:
	//
	// c0dc71d1-8a1d-4043-9767-f6c420e34901-81bd
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// Alert type.
	//
	// example:
	//
	// SUSP_CUSTOM_WAF
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Work order ID.
	//
	// example:
	//
	// 1914348
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Owner ID.
	//
	// example:
	//
	// 352675
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Owner name.
	//
	// example:
	//
	// Le Ya
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Work order name.
	//
	// example:
	//
	// 22端口禁止任意IP访问
	WorkTaskName *string `json:"WorkTaskName,omitempty" xml:"WorkTaskName,omitempty"`
}

func (s SendCustomEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendCustomEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *SendCustomEventResponseBodyData) GetCustomerName() *string {
	return s.CustomerName
}

func (s *SendCustomEventResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *SendCustomEventResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *SendCustomEventResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SendCustomEventResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SendCustomEventResponseBodyData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *SendCustomEventResponseBodyData) GetWorkTaskName() *string {
	return s.WorkTaskName
}

func (s *SendCustomEventResponseBodyData) SetCustomerId(v string) *SendCustomEventResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetCustomerName(v string) *SendCustomEventResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetEventId(v string) *SendCustomEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetEventType(v string) *SendCustomEventResponseBodyData {
	s.EventType = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetId(v int64) *SendCustomEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetOwnerId(v string) *SendCustomEventResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetOwnerName(v string) *SendCustomEventResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetWorkTaskName(v string) *SendCustomEventResponseBodyData {
	s.WorkTaskName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
