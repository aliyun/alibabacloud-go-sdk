// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBaseContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVisibleKnowledgeBaseContentsResponseBody
	GetCode() *string
	SetItems(v []*ListVisibleKnowledgeBaseContentsResponseBodyItems) *ListVisibleKnowledgeBaseContentsResponseBody
	GetItems() []*ListVisibleKnowledgeBaseContentsResponseBodyItems
	SetMessage(v string) *ListVisibleKnowledgeBaseContentsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListVisibleKnowledgeBaseContentsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListVisibleKnowledgeBaseContentsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListVisibleKnowledgeBaseContentsResponseBody
	GetTotal() *int64
}

type ListVisibleKnowledgeBaseContentsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of MCP cards.
	Items []*ListVisibleKnowledgeBaseContentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C474BFC7-7B11-5D92-971E-74AA82EC495B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of context libraries that match the query conditions.
	//
	// example:
	//
	// 3
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetItems() []*ListVisibleKnowledgeBaseContentsResponseBodyItems {
	return s.Items
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetCode(v string) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetItems(v []*ListVisibleKnowledgeBaseContentsResponseBodyItems) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.Items = v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetMessage(v string) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetPage(v int64) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.Page = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetRequestId(v string) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) SetTotal(v int64) *ListVisibleKnowledgeBaseContentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBody) Validate() error {
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

type ListVisibleKnowledgeBaseContentsResponseBodyItems struct {
	// The name of the creator.
	//
	// example:
	//
	// admin
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The description.
	//
	// example:
	//
	// Created by taishan-module-recovery
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory type.
	//
	// example:
	//
	// string_value
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-11-14T02:18:27Z
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-11-26T08:46:25Z
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ID of the data item. When tabId and orgId are the same, itemId uniquely identifies a data item. The maximum length is 128 characters.
	//
	// example:
	//
	// 8525934734583554048_prod
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The item type.
	//
	// example:
	//
	// item
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The skill name.
	//
	// example:
	//
	// cs-default-umodel-1782181212383_k8s.metric.k8s_csi_node_pv_node_cn-heyuan-acdr-1/c80cf3a4f9d6c496781591bd17d006c6f
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object bindings.
	ObjectBindings []*ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The number of resources in the FAILED state. This field is returned only when listing top-level knowledge base directories.
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// The knowledge base affiliation type. Valid values: aliding_kb_doc (DingTalk knowledge base document) and normal (common knowledge).
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The number of resources in the READY state. This field is returned only when listing top-level knowledge base directories.
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// The resource status. This field has a value only when itemType is resource.
	//
	// example:
	//
	// string_value
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// The total number of resources under the directory and its subdirectories. This field is returned only when listing top-level knowledge base directories.
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// The source type.
	//
	// example:
	//
	// AGENT
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetObjectBindings() []*ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings {
	return s.ObjectBindings
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceKind() *string {
	return s.SourceKind
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetCreatorName(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetDescription(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetDirectoryKind(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetGmtCreate(v int64) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetGmtModified(v int64) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetItemId(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetItemType(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetName(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetObjectBindings(v []*ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceFailedCount(v int64) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceKind(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceKind = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceReadyCount(v int64) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceStatus(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceTotalCount(v int64) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceType(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItems) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings struct {
	// The semantic graph name to which the object belongs. The object_id is unique within this graph.
	//
	// example:
	//
	// product
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the recommended item, which can be a **feedId*	- or a micro-application ID.
	//
	// example:
	//
	// 2676
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object name.
	//
	// example:
	//
	// 0bf4cf71-a55d-43f7-9d1e-3f9a6110ae6b
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The data type.
	//
	// example:
	//
	// table
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The display name of the object type (such as "Customer"), parsed from the graph schema. The value is null when the cache is missed.
	//
	// example:
	//
	// string_value
	ObjectTypeName *string `json:"objectTypeName,omitempty" xml:"objectTypeName,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) GetObjectTypeName() *string {
	return s.ObjectTypeName
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) SetGraphName(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) SetObjectId(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) SetObjectName(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings {
	s.ObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) SetObjectType(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) SetObjectTypeName(v string) *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings {
	s.ObjectTypeName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings) Validate() error {
	return dara.Validate(s)
}
