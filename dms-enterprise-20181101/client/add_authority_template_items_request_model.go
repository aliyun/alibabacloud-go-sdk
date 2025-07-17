// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthorityTemplateItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AddAuthorityTemplateItemsRequestItems) *AddAuthorityTemplateItemsRequest
	GetItems() []*AddAuthorityTemplateItemsRequestItems
	SetTemplateId(v int64) *AddAuthorityTemplateItemsRequest
	GetTemplateId() *int64
	SetTid(v int64) *AddAuthorityTemplateItemsRequest
	GetTid() *int64
}

type AddAuthorityTemplateItemsRequest struct {
	// This parameter is required.
	Items []*AddAuthorityTemplateItemsRequestItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 15***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddAuthorityTemplateItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuthorityTemplateItemsRequest) GoString() string {
	return s.String()
}

func (s *AddAuthorityTemplateItemsRequest) GetItems() []*AddAuthorityTemplateItemsRequestItems {
	return s.Items
}

func (s *AddAuthorityTemplateItemsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AddAuthorityTemplateItemsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddAuthorityTemplateItemsRequest) SetItems(v []*AddAuthorityTemplateItemsRequestItems) *AddAuthorityTemplateItemsRequest {
	s.Items = v
	return s
}

func (s *AddAuthorityTemplateItemsRequest) SetTemplateId(v int64) *AddAuthorityTemplateItemsRequest {
	s.TemplateId = &v
	return s
}

func (s *AddAuthorityTemplateItemsRequest) SetTid(v int64) *AddAuthorityTemplateItemsRequest {
	s.Tid = &v
	return s
}

func (s *AddAuthorityTemplateItemsRequest) Validate() error {
	return dara.Validate(s)
}

type AddAuthorityTemplateItemsRequestItems struct {
	// example:
	//
	// 2478****
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// 237****
	InstanceId      *int32    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PermissionTypes []*string `json:"PermissionTypes,omitempty" xml:"PermissionTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s AddAuthorityTemplateItemsRequestItems) String() string {
	return dara.Prettify(s)
}

func (s AddAuthorityTemplateItemsRequestItems) GoString() string {
	return s.String()
}

func (s *AddAuthorityTemplateItemsRequestItems) GetDbId() *int32 {
	return s.DbId
}

func (s *AddAuthorityTemplateItemsRequestItems) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *AddAuthorityTemplateItemsRequestItems) GetPermissionTypes() []*string {
	return s.PermissionTypes
}

func (s *AddAuthorityTemplateItemsRequestItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddAuthorityTemplateItemsRequestItems) GetTableName() *string {
	return s.TableName
}

func (s *AddAuthorityTemplateItemsRequestItems) SetDbId(v int32) *AddAuthorityTemplateItemsRequestItems {
	s.DbId = &v
	return s
}

func (s *AddAuthorityTemplateItemsRequestItems) SetInstanceId(v int32) *AddAuthorityTemplateItemsRequestItems {
	s.InstanceId = &v
	return s
}

func (s *AddAuthorityTemplateItemsRequestItems) SetPermissionTypes(v []*string) *AddAuthorityTemplateItemsRequestItems {
	s.PermissionTypes = v
	return s
}

func (s *AddAuthorityTemplateItemsRequestItems) SetResourceType(v string) *AddAuthorityTemplateItemsRequestItems {
	s.ResourceType = &v
	return s
}

func (s *AddAuthorityTemplateItemsRequestItems) SetTableName(v string) *AddAuthorityTemplateItemsRequestItems {
	s.TableName = &v
	return s
}

func (s *AddAuthorityTemplateItemsRequestItems) Validate() error {
	return dara.Validate(s)
}
