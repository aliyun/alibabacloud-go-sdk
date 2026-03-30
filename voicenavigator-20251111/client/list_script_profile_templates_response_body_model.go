// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptProfileTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptProfileTemplatesResponseBody
	GetCode() *string
	SetData(v []*ListScriptProfileTemplatesResponseBodyData) *ListScriptProfileTemplatesResponseBody
	GetData() []*ListScriptProfileTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListScriptProfileTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptProfileTemplatesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListScriptProfileTemplatesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListScriptProfileTemplatesResponseBody
	GetRequestId() *string
}

type ListScriptProfileTemplatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListScriptProfileTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// E6E61E1A-D2DC-5ACF-AED4-A115B6691F98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListScriptProfileTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptProfileTemplatesResponseBody) GetData() []*ListScriptProfileTemplatesResponseBodyData {
	return s.Data
}

func (s *ListScriptProfileTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptProfileTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptProfileTemplatesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListScriptProfileTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptProfileTemplatesResponseBody) SetCode(v string) *ListScriptProfileTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetData(v []*ListScriptProfileTemplatesResponseBodyData) *ListScriptProfileTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetHttpStatusCode(v int32) *ListScriptProfileTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetMessage(v string) *ListScriptProfileTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetParams(v []*string) *ListScriptProfileTemplatesResponseBody {
	s.Params = v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetRequestId(v string) *ListScriptProfileTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptProfileTemplatesResponseBodyData struct {
	// example:
	//
	// 1752829090000
	CreatedTime *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Schema      *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// example:
	//
	// CCC_PROMPTS_DEFAULT
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 1752829090000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListScriptProfileTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetCreatedTime(v int64) *ListScriptProfileTemplatesResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetDescription(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetName(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetSchema(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Schema = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetTemplateId(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetUpdatedTime(v int64) *ListScriptProfileTemplatesResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
