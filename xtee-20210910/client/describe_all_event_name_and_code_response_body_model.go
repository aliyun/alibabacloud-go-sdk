// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllEventNameAndCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAllEventNameAndCodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeAllEventNameAndCodeResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeAllEventNameAndCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAllEventNameAndCodeResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAllEventNameAndCodeResponseBodyResultObject) *DescribeAllEventNameAndCodeResponseBody
	GetResultObject() []*DescribeAllEventNameAndCodeResponseBodyResultObject
	SetSuccess(v bool) *DescribeAllEventNameAndCodeResponseBody
	GetSuccess() *bool
}

type DescribeAllEventNameAndCodeResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeAllEventNameAndCodeResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeAllEventNameAndCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEventNameAndCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllEventNameAndCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAllEventNameAndCodeResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeAllEventNameAndCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAllEventNameAndCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllEventNameAndCodeResponseBody) GetResultObject() []*DescribeAllEventNameAndCodeResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAllEventNameAndCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAllEventNameAndCodeResponseBody) SetCode(v string) *DescribeAllEventNameAndCodeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBody) SetHttpStatusCode(v string) *DescribeAllEventNameAndCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBody) SetMessage(v string) *DescribeAllEventNameAndCodeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBody) SetRequestId(v string) *DescribeAllEventNameAndCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBody) SetResultObject(v []*DescribeAllEventNameAndCodeResponseBodyResultObject) *DescribeAllEventNameAndCodeResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBody) SetSuccess(v bool) *DescribeAllEventNameAndCodeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAllEventNameAndCodeResponseBodyResultObject struct {
	// List of child fields.
	Children []*DescribeAllEventNameAndCodeResponseBodyResultObjectChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// Creation type
	//
	// example:
	//
	// MORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event type
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
}

func (s DescribeAllEventNameAndCodeResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEventNameAndCodeResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) GetChildren() []*DescribeAllEventNameAndCodeResponseBodyResultObjectChildren {
	return s.Children
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) GetEventType() *string {
	return s.EventType
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) SetChildren(v []*DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) *DescribeAllEventNameAndCodeResponseBodyResultObject {
	s.Children = v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) SetCreateType(v string) *DescribeAllEventNameAndCodeResponseBodyResultObject {
	s.CreateType = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) SetEventCode(v string) *DescribeAllEventNameAndCodeResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) SetEventName(v string) *DescribeAllEventNameAndCodeResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) SetEventType(v string) *DescribeAllEventNameAndCodeResponseBodyResultObject {
	s.EventType = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeAllEventNameAndCodeResponseBodyResultObjectChildren struct {
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aamexg3015
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险旁路
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event type
	//
	// example:
	//
	// BYPASS
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
}

func (s DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) GoString() string {
	return s.String()
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) GetEventType() *string {
	return s.EventType
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) SetCreateType(v string) *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren {
	s.CreateType = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) SetEventCode(v string) *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren {
	s.EventCode = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) SetEventName(v string) *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren {
	s.EventName = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) SetEventType(v string) *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren {
	s.EventType = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponseBodyResultObjectChildren) Validate() error {
	return dara.Validate(s)
}
