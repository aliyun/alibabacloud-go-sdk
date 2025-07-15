// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContactFlowsResponseBody
	GetCode() *string
	SetData(v *ListContactFlowsResponseBodyData) *ListContactFlowsResponseBody
	GetData() *ListContactFlowsResponseBodyData
	SetHttpStatusCode(v int32) *ListContactFlowsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListContactFlowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContactFlowsResponseBody
	GetRequestId() *string
}

type ListContactFlowsResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListContactFlowsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListContactFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContactFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContactFlowsResponseBody) GetData() *ListContactFlowsResponseBodyData {
	return s.Data
}

func (s *ListContactFlowsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListContactFlowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContactFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContactFlowsResponseBody) SetCode(v string) *ListContactFlowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetData(v *ListContactFlowsResponseBodyData) *ListContactFlowsResponseBody {
	s.Data = v
	return s
}

func (s *ListContactFlowsResponseBody) SetHttpStatusCode(v int32) *ListContactFlowsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetMessage(v string) *ListContactFlowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetRequestId(v string) *ListContactFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactFlowsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListContactFlowsResponseBodyData struct {
	List []*ListContactFlowsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContactFlowsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContactFlowsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponseBodyData) GetList() []*ListContactFlowsResponseBodyDataList {
	return s.List
}

func (s *ListContactFlowsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListContactFlowsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContactFlowsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListContactFlowsResponseBodyData) SetList(v []*ListContactFlowsResponseBodyDataList) *ListContactFlowsResponseBodyData {
	s.List = v
	return s
}

func (s *ListContactFlowsResponseBodyData) SetPageNumber(v int32) *ListContactFlowsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListContactFlowsResponseBodyData) SetPageSize(v int32) *ListContactFlowsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContactFlowsResponseBodyData) SetTotalCount(v int32) *ListContactFlowsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListContactFlowsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListContactFlowsResponseBodyDataList struct {
	// example:
	//
	// 78128960-bb00-4ddc-8e82-923a8c5bd22d
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// example:
	//
	// 2021-03-05 17:35:45.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// example:
	//
	// 1.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// db07c0bb-6b1f-47d2-b37e-2451c617562d
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// example:
	//
	// ccc-test
	Editor *string `json:"Editor,omitempty" xml:"Editor,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NumberList []*string `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// MAIN_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2021-03-08 15:34:49.0
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListContactFlowsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListContactFlowsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponseBodyDataList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListContactFlowsResponseBodyDataList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListContactFlowsResponseBodyDataList) GetDefinition() *string {
	return s.Definition
}

func (s *ListContactFlowsResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListContactFlowsResponseBodyDataList) GetDraftId() *string {
	return s.DraftId
}

func (s *ListContactFlowsResponseBodyDataList) GetEditor() *string {
	return s.Editor
}

func (s *ListContactFlowsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListContactFlowsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListContactFlowsResponseBodyDataList) GetNumberList() []*string {
	return s.NumberList
}

func (s *ListContactFlowsResponseBodyDataList) GetPublished() *bool {
	return s.Published
}

func (s *ListContactFlowsResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *ListContactFlowsResponseBodyDataList) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListContactFlowsResponseBodyDataList) SetContactFlowId(v string) *ListContactFlowsResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetCreatedTime(v string) *ListContactFlowsResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetDefinition(v string) *ListContactFlowsResponseBodyDataList {
	s.Definition = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetDescription(v string) *ListContactFlowsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetDraftId(v string) *ListContactFlowsResponseBodyDataList {
	s.DraftId = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetEditor(v string) *ListContactFlowsResponseBodyDataList {
	s.Editor = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetInstanceId(v string) *ListContactFlowsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetName(v string) *ListContactFlowsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetNumberList(v []*string) *ListContactFlowsResponseBodyDataList {
	s.NumberList = v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetPublished(v bool) *ListContactFlowsResponseBodyDataList {
	s.Published = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetType(v string) *ListContactFlowsResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetUpdatedTime(v string) *ListContactFlowsResponseBodyDataList {
	s.UpdatedTime = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
