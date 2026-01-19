// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentResponseBody
	GetCode() *string
	SetData(v *GetDocumentResponseBodyData) *GetDocumentResponseBody
	GetData() *GetDocumentResponseBodyData
	SetHttpStatusCode(v int32) *GetDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocumentResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetDocumentResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetDocumentResponseBody
	GetRequestId() *string
}

type GetDocumentResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentResponseBody) GetData() *GetDocumentResponseBodyData {
	return s.Data
}

func (s *GetDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentResponseBody) SetCode(v string) *GetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentResponseBody) SetData(v *GetDocumentResponseBodyData) *GetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentResponseBody) SetHttpStatusCode(v int32) *GetDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentResponseBody) SetMessage(v string) *GetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentResponseBody) SetParams(v []*string) *GetDocumentResponseBody {
	s.Params = v
	return s
}

func (s *GetDocumentResponseBody) SetRequestId(v string) *GetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentResponseBodyData struct {
	// example:
	//
	// {"name":"tom"}
	Document map[string]interface{} `json:"Document,omitempty" xml:"Document,omitempty"`
	// scheme
	Schema *GetDocumentResponseBodyDataSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Struct"`
}

func (s GetDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentResponseBodyData) GetDocument() map[string]interface{} {
	return s.Document
}

func (s *GetDocumentResponseBodyData) GetSchema() *GetDocumentResponseBodyDataSchema {
	return s.Schema
}

func (s *GetDocumentResponseBodyData) SetDocument(v map[string]interface{}) *GetDocumentResponseBodyData {
	s.Document = v
	return s
}

func (s *GetDocumentResponseBodyData) SetSchema(v *GetDocumentResponseBodyDataSchema) *GetDocumentResponseBodyData {
	s.Schema = v
	return s
}

func (s *GetDocumentResponseBodyData) Validate() error {
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentResponseBodyDataSchema struct {
	// example:
	//
	// 2024-10-16T02:12:12Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// schame id
	//
	// example:
	//
	// 400135
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 57236c4f-48e9-49ca-a560-8697ec6c185e
	InstanceId *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Properties map[string]*DataSchemaPropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// 2024-10-16T02:12:12Z
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetDocumentResponseBodyDataSchema) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentResponseBodyDataSchema) GoString() string {
	return s.String()
}

func (s *GetDocumentResponseBodyDataSchema) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetDocumentResponseBodyDataSchema) GetDeleted() *bool {
	return s.Deleted
}

func (s *GetDocumentResponseBodyDataSchema) GetDescription() *string {
	return s.Description
}

func (s *GetDocumentResponseBodyDataSchema) GetId() *string {
	return s.Id
}

func (s *GetDocumentResponseBodyDataSchema) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDocumentResponseBodyDataSchema) GetProperties() map[string]*DataSchemaPropertiesValue {
	return s.Properties
}

func (s *GetDocumentResponseBodyDataSchema) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *GetDocumentResponseBodyDataSchema) SetCreatedTime(v string) *GetDocumentResponseBodyDataSchema {
	s.CreatedTime = &v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) SetDeleted(v bool) *GetDocumentResponseBodyDataSchema {
	s.Deleted = &v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) SetDescription(v string) *GetDocumentResponseBodyDataSchema {
	s.Description = &v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) SetId(v string) *GetDocumentResponseBodyDataSchema {
	s.Id = &v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) SetInstanceId(v string) *GetDocumentResponseBodyDataSchema {
	s.InstanceId = &v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) SetProperties(v map[string]*DataSchemaPropertiesValue) *GetDocumentResponseBodyDataSchema {
	s.Properties = v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) SetUpdatedTime(v string) *GetDocumentResponseBodyDataSchema {
	s.UpdatedTime = &v
	return s
}

func (s *GetDocumentResponseBodyDataSchema) Validate() error {
	return dara.Validate(s)
}
