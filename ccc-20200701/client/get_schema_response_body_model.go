// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSchemaResponseBody
	GetCode() *string
	SetData(v *GetSchemaResponseBodyData) *GetSchemaResponseBody
	GetData() *GetSchemaResponseBodyData
	SetHttpStatusCode(v int32) *GetSchemaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSchemaResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetSchemaResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetSchemaResponseBody
	GetRequestId() *string
}

type GetSchemaResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSchemaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Service abnormal, the instance 0418 is ceased.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 19D09CCC-F298-4124-849A-AFA217819011
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSchemaResponseBody) GetData() *GetSchemaResponseBodyData {
	return s.Data
}

func (s *GetSchemaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSchemaResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSchemaResponseBody) SetCode(v string) *GetSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *GetSchemaResponseBody) SetData(v *GetSchemaResponseBodyData) *GetSchemaResponseBody {
	s.Data = v
	return s
}

func (s *GetSchemaResponseBody) SetHttpStatusCode(v int32) *GetSchemaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSchemaResponseBody) SetMessage(v string) *GetSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *GetSchemaResponseBody) SetParams(v []*string) *GetSchemaResponseBody {
	s.Params = v
	return s
}

func (s *GetSchemaResponseBody) SetRequestId(v string) *GetSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSchemaResponseBodyData struct {
	// example:
	//
	// 2021-07-14 10:48:43.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// false
	Deleted     *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// schema id
	//
	// example:
	//
	// profile
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 5e0964fd-951c-4e45-b518-d09d4d2db8ca
	InstanceId *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Properties map[string]*DataPropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// 2021-07-14 10:48:43.0
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetSchemaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSchemaResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetSchemaResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *GetSchemaResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetSchemaResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetSchemaResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSchemaResponseBodyData) GetProperties() map[string]*DataPropertiesValue {
	return s.Properties
}

func (s *GetSchemaResponseBodyData) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *GetSchemaResponseBodyData) SetCreatedTime(v string) *GetSchemaResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetSchemaResponseBodyData) SetDeleted(v bool) *GetSchemaResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *GetSchemaResponseBodyData) SetDescription(v string) *GetSchemaResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSchemaResponseBodyData) SetId(v string) *GetSchemaResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSchemaResponseBodyData) SetInstanceId(v string) *GetSchemaResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSchemaResponseBodyData) SetProperties(v map[string]*DataPropertiesValue) *GetSchemaResponseBodyData {
	s.Properties = v
	return s
}

func (s *GetSchemaResponseBodyData) SetUpdatedTime(v string) *GetSchemaResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetSchemaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
