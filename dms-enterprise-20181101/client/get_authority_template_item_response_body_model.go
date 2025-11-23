// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorityTemplateItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityTemplateItemList(v *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) *GetAuthorityTemplateItemResponseBody
	GetAuthorityTemplateItemList() *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList
	SetErrorCode(v string) *GetAuthorityTemplateItemResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAuthorityTemplateItemResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAuthorityTemplateItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuthorityTemplateItemResponseBody
	GetSuccess() *bool
	SetTid(v int64) *GetAuthorityTemplateItemResponseBody
	GetTid() *int64
}

type GetAuthorityTemplateItemResponseBody struct {
	// The permission templates.
	AuthorityTemplateItemList *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList `json:"AuthorityTemplateItemList,omitempty" xml:"AuthorityTemplateItemList,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5B96E35F-A58E-5399-9041-09CF9A1E46EA
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

func (s GetAuthorityTemplateItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateItemResponseBody) GetAuthorityTemplateItemList() *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList {
	return s.AuthorityTemplateItemList
}

func (s *GetAuthorityTemplateItemResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAuthorityTemplateItemResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAuthorityTemplateItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorityTemplateItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuthorityTemplateItemResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *GetAuthorityTemplateItemResponseBody) SetAuthorityTemplateItemList(v *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) *GetAuthorityTemplateItemResponseBody {
	s.AuthorityTemplateItemList = v
	return s
}

func (s *GetAuthorityTemplateItemResponseBody) SetErrorCode(v string) *GetAuthorityTemplateItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBody) SetErrorMessage(v string) *GetAuthorityTemplateItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBody) SetRequestId(v string) *GetAuthorityTemplateItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBody) SetSuccess(v bool) *GetAuthorityTemplateItemResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBody) SetTid(v int64) *GetAuthorityTemplateItemResponseBody {
	s.Tid = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBody) Validate() error {
	if s.AuthorityTemplateItemList != nil {
		if err := s.AuthorityTemplateItemList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList struct {
	AuthorityTemplateItem []*GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem `json:"AuthorityTemplateItem,omitempty" xml:"AuthorityTemplateItem,omitempty" type:"Repeated"`
}

func (s GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) GetAuthorityTemplateItem() []*GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	return s.AuthorityTemplateItem
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) SetAuthorityTemplateItem(v []*GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList {
	s.AuthorityTemplateItem = v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemList) Validate() error {
	if s.AuthorityTemplateItem != nil {
		for _, item := range s.AuthorityTemplateItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem struct {
	// The additional information. For example, permissions to log on to an instance are added to the permission template.
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
	// The ID of the user who modifies the resource.
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
	// 	- **LOGIC_TABLE**: physical table
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

func (s GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetAttribute() *string {
	return s.Attribute
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetDbId() *int64 {
	return s.DbId
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetTableName() *string {
	return s.TableName
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetAttribute(v string) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.Attribute = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetDbId(v int64) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.DbId = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetInstanceId(v int64) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetItemId(v int64) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.ItemId = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetModifierId(v int64) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.ModifierId = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetResourceType(v string) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.ResourceType = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetTableName(v string) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.TableName = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) SetTemplateId(v int64) *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem {
	s.TemplateId = &v
	return s
}

func (s *GetAuthorityTemplateItemResponseBodyAuthorityTemplateItemListAuthorityTemplateItem) Validate() error {
	return dara.Validate(s)
}
