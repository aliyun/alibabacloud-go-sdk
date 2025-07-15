// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDevicesResponseBody
	GetCode() *string
	SetData(v []*ListDevicesResponseBodyData) *ListDevicesResponseBody
	GetData() []*ListDevicesResponseBodyData
	SetHttpStatusCode(v int32) *ListDevicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDevicesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListDevicesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListDevicesResponseBody
	GetRequestId() *string
}

type ListDevicesResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDevicesResponseBody) GetData() []*ListDevicesResponseBodyData {
	return s.Data
}

func (s *ListDevicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDevicesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDevicesResponseBody) SetCode(v string) *ListDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDevicesResponseBody) SetData(v []*ListDevicesResponseBodyData) *ListDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicesResponseBody) SetHttpStatusCode(v int32) *ListDevicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDevicesResponseBody) SetMessage(v string) *ListDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicesResponseBody) SetParams(v []*string) *ListDevicesResponseBody {
	s.Params = v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyData struct {
	// example:
	//
	// d7b818c3-8d3a-732f-bc9e-1782wa16****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// sip:8032****@33.89.XX.XX:64189;transport=tcp;registering_acc=18_134_23_4
	Contact *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1609118499750
	Expires *int64 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDevicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *ListDevicesResponseBodyData) GetContact() *string {
	return s.Contact
}

func (s *ListDevicesResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListDevicesResponseBodyData) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListDevicesResponseBodyData) GetExpires() *int64 {
	return s.Expires
}

func (s *ListDevicesResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ListDevicesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDevicesResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListDevicesResponseBodyData) SetCallId(v string) *ListDevicesResponseBodyData {
	s.CallId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetContact(v string) *ListDevicesResponseBodyData {
	s.Contact = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceId(v string) *ListDevicesResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceType(v string) *ListDevicesResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetExpires(v int64) *ListDevicesResponseBodyData {
	s.Expires = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetExtension(v string) *ListDevicesResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetInstanceId(v string) *ListDevicesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetUserId(v string) *ListDevicesResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDevicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
