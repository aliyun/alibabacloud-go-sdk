// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonTicketFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCommonTicketFieldsResponseBody
	GetCode() *string
	SetData(v *ListCommonTicketFieldsResponseBodyData) *ListCommonTicketFieldsResponseBody
	GetData() *ListCommonTicketFieldsResponseBodyData
	SetHttpStatusCode(v int32) *ListCommonTicketFieldsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCommonTicketFieldsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCommonTicketFieldsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCommonTicketFieldsResponseBody
	GetRequestId() *string
}

type ListCommonTicketFieldsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCommonTicketFieldsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 30C7D235-DDCF-4C7F-A462-5E2598252C2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCommonTicketFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommonTicketFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonTicketFieldsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCommonTicketFieldsResponseBody) GetData() *ListCommonTicketFieldsResponseBodyData {
	return s.Data
}

func (s *ListCommonTicketFieldsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCommonTicketFieldsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCommonTicketFieldsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCommonTicketFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommonTicketFieldsResponseBody) SetCode(v string) *ListCommonTicketFieldsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBody) SetData(v *ListCommonTicketFieldsResponseBodyData) *ListCommonTicketFieldsResponseBody {
	s.Data = v
	return s
}

func (s *ListCommonTicketFieldsResponseBody) SetHttpStatusCode(v int32) *ListCommonTicketFieldsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBody) SetMessage(v string) *ListCommonTicketFieldsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBody) SetParams(v []*string) *ListCommonTicketFieldsResponseBody {
	s.Params = v
	return s
}

func (s *ListCommonTicketFieldsResponseBody) SetRequestId(v string) *ListCommonTicketFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCommonTicketFieldsResponseBodyData struct {
	// example:
	//
	// 1703517780627
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Schema ID。
	//
	// example:
	//
	// ticketing
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Properties map[string]*DataPropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// 1716211430928
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListCommonTicketFieldsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCommonTicketFieldsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCommonTicketFieldsResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListCommonTicketFieldsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListCommonTicketFieldsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCommonTicketFieldsResponseBodyData) GetProperties() map[string]*DataPropertiesValue {
	return s.Properties
}

func (s *ListCommonTicketFieldsResponseBodyData) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListCommonTicketFieldsResponseBodyData) SetCreatedTime(v string) *ListCommonTicketFieldsResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBodyData) SetId(v string) *ListCommonTicketFieldsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBodyData) SetInstanceId(v string) *ListCommonTicketFieldsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBodyData) SetProperties(v map[string]*DataPropertiesValue) *ListCommonTicketFieldsResponseBodyData {
	s.Properties = v
	return s
}

func (s *ListCommonTicketFieldsResponseBodyData) SetUpdatedTime(v string) *ListCommonTicketFieldsResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *ListCommonTicketFieldsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
