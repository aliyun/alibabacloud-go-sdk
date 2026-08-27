// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAdminKnowledgeBasesResponseBody
	GetCode() *string
	SetItems(v []*ListAdminKnowledgeBasesResponseBodyItems) *ListAdminKnowledgeBasesResponseBody
	GetItems() []*ListAdminKnowledgeBasesResponseBodyItems
	SetMessage(v string) *ListAdminKnowledgeBasesResponseBody
	GetMessage() *string
	SetPage(v int64) *ListAdminKnowledgeBasesResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListAdminKnowledgeBasesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAdminKnowledgeBasesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAdminKnowledgeBasesResponseBody
	GetTotal() *int64
}

type ListAdminKnowledgeBasesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of MCP cards.
	Items []*ListAdminKnowledgeBasesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAdminKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAdminKnowledgeBasesResponseBody) GetItems() []*ListAdminKnowledgeBasesResponseBodyItems {
	return s.Items
}

func (s *ListAdminKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAdminKnowledgeBasesResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListAdminKnowledgeBasesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdminKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAdminKnowledgeBasesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAdminKnowledgeBasesResponseBody) SetCode(v string) *ListAdminKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) SetItems(v []*ListAdminKnowledgeBasesResponseBodyItems) *ListAdminKnowledgeBasesResponseBody {
	s.Items = v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) SetMessage(v string) *ListAdminKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) SetPage(v int64) *ListAdminKnowledgeBasesResponseBody {
	s.Page = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) SetPageSize(v int64) *ListAdminKnowledgeBasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) SetRequestId(v string) *ListAdminKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) SetTotal(v int64) *ListAdminKnowledgeBasesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBody) Validate() error {
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

type ListAdminKnowledgeBasesResponseBodyItems struct {
	// The name of the creator.
	//
	// example:
	//
	// string_value
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
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
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ID of the data item. When tabId and orgId are the same, itemId uniquely identifies a data item. The maximum length is 128 characters.
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The data type (group, user, or role).
	//
	// example:
	//
	// string_value
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object bindings.
	ObjectBindings []*ListAdminKnowledgeBasesResponseBodyItemsObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The number of resources with FAILED status. This field is returned only for the top-level knowledge base directory list.
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// The knowledge base ownership type. Valid values: aliding_kb_doc (DingTalk knowledge base document) and normal (common knowledge).
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The number of resources with READY status. This field is returned only for the top-level knowledge base directory list.
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
	// The total number of resources in the directory and its subdirectories. This field is returned only for the top-level knowledge base directory list.
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// The source type.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListAdminKnowledgeBasesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetObjectBindings() []*ListAdminKnowledgeBasesResponseBodyItemsObjectBindings {
	return s.ObjectBindings
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetSourceKind() *string {
	return s.SourceKind
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetCreatorName(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetDescription(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetDirectoryKind(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetGmtCreate(v int64) *ListAdminKnowledgeBasesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetGmtModified(v int64) *ListAdminKnowledgeBasesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetItemId(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetItemType(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetName(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetObjectBindings(v []*ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) *ListAdminKnowledgeBasesResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetSourceFailedCount(v int64) *ListAdminKnowledgeBasesResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetSourceKind(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.SourceKind = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetSourceReadyCount(v int64) *ListAdminKnowledgeBasesResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetSourceStatus(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetSourceTotalCount(v int64) *ListAdminKnowledgeBasesResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) SetSourceType(v string) *ListAdminKnowledgeBasesResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItems) Validate() error {
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

type ListAdminKnowledgeBasesResponseBodyItemsObjectBindings struct {
	// The semantic graph name to which the object belongs. The object_id is unique within this graph.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the recommended item, which can be a feedId or a micro-application ID.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object name.
	//
	// example:
	//
	// string_value
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The object type, such as customer. This field has a value only when type is mention.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The display name of the object type (such as "Customer"), parsed from the graph schema. The value is null when the cache is missed.
	//
	// example:
	//
	// string_value
	ObjectTypeName *string `json:"objectTypeName,omitempty" xml:"objectTypeName,omitempty"`
}

func (s ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) GetObjectTypeName() *string {
	return s.ObjectTypeName
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) SetGraphName(v string) *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) SetObjectId(v string) *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) SetObjectName(v string) *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings {
	s.ObjectName = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) SetObjectType(v string) *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) SetObjectTypeName(v string) *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings {
	s.ObjectTypeName = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponseBodyItemsObjectBindings) Validate() error {
	return dara.Validate(s)
}
