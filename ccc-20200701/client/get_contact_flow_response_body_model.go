// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContactFlowResponseBody
	GetCode() *string
	SetData(v *GetContactFlowResponseBodyData) *GetContactFlowResponseBody
	GetData() *GetContactFlowResponseBodyData
	SetHttpStatusCode(v int32) *GetContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContactFlowResponseBody
	GetRequestId() *string
}

type GetContactFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetContactFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2263B273-AC1B-44EB-BA98-87F2322C6780
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContactFlowResponseBody) GetData() *GetContactFlowResponseBodyData {
	return s.Data
}

func (s *GetContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContactFlowResponseBody) SetCode(v string) *GetContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetContactFlowResponseBody) SetData(v *GetContactFlowResponseBodyData) *GetContactFlowResponseBody {
	s.Data = v
	return s
}

func (s *GetContactFlowResponseBody) SetHttpStatusCode(v int32) *GetContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContactFlowResponseBody) SetMessage(v string) *GetContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetContactFlowResponseBody) SetRequestId(v string) *GetContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetContactFlowResponseBodyData struct {
	// example:
	//
	// 274601be-a6d5-4429-bcef-32b51d031c6e
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// example:
	//
	// 2021-07-14 10:48:43.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// example:
	//
	// 1.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 566399d7-5558-447c-a72f-9be2768b6a82
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// example:
	//
	// editor-xxx
	Editor *string `json:"Editor,omitempty" xml:"Editor,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// False
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// MAIN_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2021-07-14 10:48:43.0
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetContactFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContactFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContactFlowResponseBodyData) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *GetContactFlowResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetContactFlowResponseBodyData) GetDefinition() *string {
	return s.Definition
}

func (s *GetContactFlowResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetContactFlowResponseBodyData) GetDraftId() *string {
	return s.DraftId
}

func (s *GetContactFlowResponseBodyData) GetEditor() *string {
	return s.Editor
}

func (s *GetContactFlowResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetContactFlowResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetContactFlowResponseBodyData) GetPublished() *bool {
	return s.Published
}

func (s *GetContactFlowResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetContactFlowResponseBodyData) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *GetContactFlowResponseBodyData) SetContactFlowId(v string) *GetContactFlowResponseBodyData {
	s.ContactFlowId = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetCreatedTime(v string) *GetContactFlowResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetDefinition(v string) *GetContactFlowResponseBodyData {
	s.Definition = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetDescription(v string) *GetContactFlowResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetDraftId(v string) *GetContactFlowResponseBodyData {
	s.DraftId = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetEditor(v string) *GetContactFlowResponseBodyData {
	s.Editor = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetInstanceId(v string) *GetContactFlowResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetName(v string) *GetContactFlowResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetPublished(v bool) *GetContactFlowResponseBodyData {
	s.Published = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetType(v string) *GetContactFlowResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetContactFlowResponseBodyData) SetUpdatedTime(v string) *GetContactFlowResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetContactFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
