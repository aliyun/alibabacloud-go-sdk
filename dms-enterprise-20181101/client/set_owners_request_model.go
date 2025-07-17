// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOwnersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerIds(v string) *SetOwnersRequest
	GetOwnerIds() *string
	SetOwnerType(v string) *SetOwnersRequest
	GetOwnerType() *string
	SetResourceId(v string) *SetOwnersRequest
	GetResourceId() *string
	SetTid(v int64) *SetOwnersRequest
	GetTid() *int64
}

type SetOwnersRequest struct {
	// The ID of the user whom you want to specify as an owner. Separate multiple IDs with commas (,). You can call the [GetUser](https://help.aliyun.com/document_detail/147098.html) or [ListUsers](https://help.aliyun.com/document_detail/141938.html) operation to query the ID of the user.
	//
	// >  The value of the OwnerIds parameter is that of the UserId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51****
	OwnerIds *string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty"`
	// The type of the owner. Valid values:
	//
	// 	- INSTANCE: an owner of an instance.
	//
	// 	- DATABASE: an owner of a physical database.
	//
	// 	- LOGIC_DATABASE: an owner of a logical database.
	//
	// 	- TABLE: an owner of a physical table.
	//
	// 	- LOGIC_TABLE: an owner of a logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The ID of the resource. The ID of the resource varies with the owner type. The owner types and resource IDs have the following mappings:
	//
	// 	- INSTANCE: the ID of an instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) operation to query the ID of the instance.
	//
	// 	- DATABASE: the ID of a physical database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of the physical database.
	//
	// 	- LOGIC_DATABASE: the ID of a logical database. You can call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of the logical database.
	//
	// 	- TABLE: the ID of a physical table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the ID of the physical table.
	//
	// 	- LOGIC_DATABASE: the ID of a logical table. You can call the [ListLogicTables](https://help.aliyun.com/document_detail/141875.html) operation to query the ID of the logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SetOwnersRequest) String() string {
	return dara.Prettify(s)
}

func (s SetOwnersRequest) GoString() string {
	return s.String()
}

func (s *SetOwnersRequest) GetOwnerIds() *string {
	return s.OwnerIds
}

func (s *SetOwnersRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *SetOwnersRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *SetOwnersRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SetOwnersRequest) SetOwnerIds(v string) *SetOwnersRequest {
	s.OwnerIds = &v
	return s
}

func (s *SetOwnersRequest) SetOwnerType(v string) *SetOwnersRequest {
	s.OwnerType = &v
	return s
}

func (s *SetOwnersRequest) SetResourceId(v string) *SetOwnersRequest {
	s.ResourceId = &v
	return s
}

func (s *SetOwnersRequest) SetTid(v int64) *SetOwnersRequest {
	s.Tid = &v
	return s
}

func (s *SetOwnersRequest) Validate() error {
	return dara.Validate(s)
}
