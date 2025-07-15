// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDocumentsResponseBody
	GetCode() *string
	SetData(v *ListDocumentsResponseBodyData) *ListDocumentsResponseBody
	GetData() *ListDocumentsResponseBodyData
	SetHttpStatusCode(v int32) *ListDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDocumentsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListDocumentsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListDocumentsResponseBody
	GetRequestId() *string
}

type ListDocumentsResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDocumentsResponseBody) GetData() *ListDocumentsResponseBodyData {
	return s.Data
}

func (s *ListDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDocumentsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentsResponseBody) SetCode(v string) *ListDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDocumentsResponseBody) SetData(v *ListDocumentsResponseBodyData) *ListDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDocumentsResponseBody) SetHttpStatusCode(v int32) *ListDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDocumentsResponseBody) SetMessage(v string) *ListDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocumentsResponseBody) SetParams(v []*string) *ListDocumentsResponseBody {
	s.Params = v
	return s
}

func (s *ListDocumentsResponseBody) SetRequestId(v string) *ListDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDocumentsResponseBodyData struct {
	Documents []map[string]interface{} `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// 54d1a616d95a4a01ba58967a9115b649
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// scheme
	Schema *ListDocumentsResponseBodyDataSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBodyData) GetDocuments() []map[string]interface{} {
	return s.Documents
}

func (s *ListDocumentsResponseBodyData) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDocumentsResponseBodyData) GetSchema() *ListDocumentsResponseBodyDataSchema {
	return s.Schema
}

func (s *ListDocumentsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDocumentsResponseBodyData) SetDocuments(v []map[string]interface{}) *ListDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListDocumentsResponseBodyData) SetNextPageToken(v string) *ListDocumentsResponseBodyData {
	s.NextPageToken = &v
	return s
}

func (s *ListDocumentsResponseBodyData) SetSchema(v *ListDocumentsResponseBodyDataSchema) *ListDocumentsResponseBodyData {
	s.Schema = v
	return s
}

func (s *ListDocumentsResponseBodyData) SetTotalCount(v int64) *ListDocumentsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDocumentsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDocumentsResponseBodyDataSchema struct {
	// example:
	//
	// 2020-10-14T09:53:53Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// schame id
	//
	// example:
	//
	// profile
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 87b12784-8ce2-40b6-b21f-c49cb3b5501e
	InstanceId *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Properties map[string]*DataSchemaPropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// 2020-10-14T09:53:53Z
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListDocumentsResponseBodyDataSchema) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBodyDataSchema) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBodyDataSchema) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDocumentsResponseBodyDataSchema) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListDocumentsResponseBodyDataSchema) GetDescription() *string {
	return s.Description
}

func (s *ListDocumentsResponseBodyDataSchema) GetId() *string {
	return s.Id
}

func (s *ListDocumentsResponseBodyDataSchema) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDocumentsResponseBodyDataSchema) GetProperties() map[string]*DataSchemaPropertiesValue {
	return s.Properties
}

func (s *ListDocumentsResponseBodyDataSchema) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListDocumentsResponseBodyDataSchema) SetCreatedTime(v string) *ListDocumentsResponseBodyDataSchema {
	s.CreatedTime = &v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) SetDeleted(v bool) *ListDocumentsResponseBodyDataSchema {
	s.Deleted = &v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) SetDescription(v string) *ListDocumentsResponseBodyDataSchema {
	s.Description = &v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) SetId(v string) *ListDocumentsResponseBodyDataSchema {
	s.Id = &v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) SetInstanceId(v string) *ListDocumentsResponseBodyDataSchema {
	s.InstanceId = &v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) SetProperties(v map[string]*DataSchemaPropertiesValue) *ListDocumentsResponseBodyDataSchema {
	s.Properties = v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) SetUpdatedTime(v string) *ListDocumentsResponseBodyDataSchema {
	s.UpdatedTime = &v
	return s
}

func (s *ListDocumentsResponseBodyDataSchema) Validate() error {
	return dara.Validate(s)
}
