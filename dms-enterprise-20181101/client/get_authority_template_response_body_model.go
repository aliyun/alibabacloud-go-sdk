// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityTemplateView(v *GetAuthorityTemplateResponseBodyAuthorityTemplateView) *GetAuthorityTemplateResponseBody
	GetAuthorityTemplateView() *GetAuthorityTemplateResponseBodyAuthorityTemplateView
	SetErrorCode(v string) *GetAuthorityTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAuthorityTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAuthorityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuthorityTemplateResponseBody
	GetSuccess() *bool
	SetTid(v int64) *GetAuthorityTemplateResponseBody
	GetTid() *int64
}

type GetAuthorityTemplateResponseBody struct {
	// The details of the permission template.
	AuthorityTemplateView *GetAuthorityTemplateResponseBodyAuthorityTemplateView `json:"AuthorityTemplateView,omitempty" xml:"AuthorityTemplateView,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetAuthorityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateResponseBody) GetAuthorityTemplateView() *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	return s.AuthorityTemplateView
}

func (s *GetAuthorityTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAuthorityTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAuthorityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuthorityTemplateResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *GetAuthorityTemplateResponseBody) SetAuthorityTemplateView(v *GetAuthorityTemplateResponseBodyAuthorityTemplateView) *GetAuthorityTemplateResponseBody {
	s.AuthorityTemplateView = v
	return s
}

func (s *GetAuthorityTemplateResponseBody) SetErrorCode(v string) *GetAuthorityTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAuthorityTemplateResponseBody) SetErrorMessage(v string) *GetAuthorityTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAuthorityTemplateResponseBody) SetRequestId(v string) *GetAuthorityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBody) SetSuccess(v bool) *GetAuthorityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuthorityTemplateResponseBody) SetTid(v int64) *GetAuthorityTemplateResponseBody {
	s.Tid = &v
	return s
}

func (s *GetAuthorityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuthorityTemplateResponseBodyAuthorityTemplateView struct {
	// The resource information in the permission template.
	AuthorityTemplateItemList *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList `json:"AuthorityTemplateItemList,omitempty" xml:"AuthorityTemplateItemList,omitempty" type:"Struct"`
	// The time when the permission template was created. The time is in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the permission template.
	//
	// example:
	//
	// 12***
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the permission template.
	//
	// example:
	//
	// This template is used for business testing.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission template.
	//
	// example:
	//
	// TestTemplate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the permission template.
	//
	// example:
	//
	// 1563
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAuthorityTemplateResponseBodyAuthorityTemplateView) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateResponseBodyAuthorityTemplateView) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) GetAuthorityTemplateItemList() *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList {
	return s.AuthorityTemplateItemList
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) GetDescription() *string {
	return s.Description
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) GetName() *string {
	return s.Name
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) SetAuthorityTemplateItemList(v *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList) *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.AuthorityTemplateItemList = v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) SetCreateTime(v string) *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.CreateTime = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) SetCreatorId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.CreatorId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) SetDescription(v string) *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.Description = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) SetName(v string) *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.Name = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) SetTemplateId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.TemplateId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateView) Validate() error {
	return dara.Validate(s)
}

type GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList struct {
	AuthorityTemplateItem []*GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem `json:"AuthorityTemplateItem,omitempty" xml:"AuthorityTemplateItem,omitempty" type:"Repeated"`
}

func (s GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList) GetAuthorityTemplateItem() []*GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	return s.AuthorityTemplateItem
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList) SetAuthorityTemplateItem(v []*GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList {
	s.AuthorityTemplateItem = v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemList) Validate() error {
	return dara.Validate(s)
}

type GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem struct {
	// Other information. For example, you can add the logon permission on an instance to the permission template.
	//
	// example:
	//
	// "permissionTypes": [
	//
	//             "LOGIN"
	//
	//           ]
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 43***
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 188****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// 12***
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The ID of the user who modified the resource.
	//
	// example:
	//
	// 51***
	ModifierId *int64 `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **INSTANCE**: instance
	//
	// 	- **LOGIC_DB**: logical database
	//
	// 	- **META_DB**: physical database
	//
	// 	- **LOGIC_TABLE**: logical table
	//
	// 	- **SINGLE_TABLE**: physical table
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// ExampleTable
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the permission template.
	//
	// example:
	//
	// 1563
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetAttribute() *string {
	return s.Attribute
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetDbId() *int64 {
	return s.DbId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetTableName() *string {
	return s.TableName
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetAttribute(v string) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.Attribute = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetDbId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.DbId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetInstanceId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetItemId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.ItemId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetModifierId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.ModifierId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetResourceType(v string) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.ResourceType = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetTableName(v string) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.TableName = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) SetTemplateId(v int64) *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem {
	s.TemplateId = &v
	return s
}

func (s *GetAuthorityTemplateResponseBodyAuthorityTemplateViewAuthorityTemplateItemListAuthorityTemplateItem) Validate() error {
	return dara.Validate(s)
}
