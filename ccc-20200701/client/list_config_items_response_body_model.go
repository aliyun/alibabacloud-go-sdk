// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConfigItemsResponseBody
	GetCode() *string
	SetData(v []*ListConfigItemsResponseBodyData) *ListConfigItemsResponseBody
	GetData() []*ListConfigItemsResponseBodyData
	SetHttpStatusCode(v int32) *ListConfigItemsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListConfigItemsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListConfigItemsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListConfigItemsResponseBody
	GetRequestId() *string
}

type ListConfigItemsResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListConfigItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListConfigItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConfigItemsResponseBody) GetData() []*ListConfigItemsResponseBodyData {
	return s.Data
}

func (s *ListConfigItemsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConfigItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConfigItemsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListConfigItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigItemsResponseBody) SetCode(v string) *ListConfigItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetData(v []*ListConfigItemsResponseBodyData) *ListConfigItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListConfigItemsResponseBody) SetHttpStatusCode(v int32) *ListConfigItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetMessage(v string) *ListConfigItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetParams(v []*string) *ListConfigItemsResponseBody {
	s.Params = v
	return s
}

func (s *ListConfigItemsResponseBody) SetRequestId(v string) *ListConfigItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConfigItemsResponseBodyData struct {
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// config-item
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ccc-test
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConfigItemsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConfigItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConfigItemsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConfigItemsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListConfigItemsResponseBodyData) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListConfigItemsResponseBodyData) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListConfigItemsResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *ListConfigItemsResponseBodyData) SetInstanceId(v string) *ListConfigItemsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetName(v string) *ListConfigItemsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetObjectId(v string) *ListConfigItemsResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetObjectType(v string) *ListConfigItemsResponseBodyData {
	s.ObjectType = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetValue(v string) *ListConfigItemsResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
