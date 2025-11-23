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
	// The resources that you want to add to the permission template.
	//
	// This parameter is required.
	Items []*AddAuthorityTemplateItemsRequestItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the permission template. You can call the [CreateAuthorityTemplate](https://help.aliyun.com/document_detail/600705.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
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
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddAuthorityTemplateItemsRequestItems struct {
	// The database ID. Databases are divided into physical databases and logical databases.
	//
	// 	- To query the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To query the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// > This parameter is required if the ResourceType parameter is set to META_DB, LOGIC_DB, META_TABLE, or LOGIC_TABLE.
	//
	// example:
	//
	// 2478****
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The instance ID. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// > This parameter is required if the ResourceType parameter is set to INSTANCE.
	//
	// example:
	//
	// 237****
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The permission types.
	PermissionTypes []*string `json:"PermissionTypes,omitempty" xml:"PermissionTypes,omitempty" type:"Repeated"`
	// The type of the resource from which you want to remove tags. Valid values:
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
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The table name. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the name of the table.
	//
	// > This parameter is required if the ResourceType parameter is set to META_TABLE or LOGIC_TABLE.
	//
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
