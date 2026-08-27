// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalDirectoryContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPersonalDirectoryContentsResponseBody
	GetCode() *string
	SetItems(v []*ListPersonalDirectoryContentsResponseBodyItems) *ListPersonalDirectoryContentsResponseBody
	GetItems() []*ListPersonalDirectoryContentsResponseBodyItems
	SetMessage(v string) *ListPersonalDirectoryContentsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListPersonalDirectoryContentsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalDirectoryContentsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListPersonalDirectoryContentsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListPersonalDirectoryContentsResponseBody
	GetTotal() *int64
}

type ListPersonalDirectoryContentsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The file information.
	Items []*ListPersonalDirectoryContentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries returned per page. Default value: 10.
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
	// The total number of records.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPersonalDirectoryContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPersonalDirectoryContentsResponseBody) GetItems() []*ListPersonalDirectoryContentsResponseBodyItems {
	return s.Items
}

func (s *ListPersonalDirectoryContentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPersonalDirectoryContentsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalDirectoryContentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalDirectoryContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPersonalDirectoryContentsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListPersonalDirectoryContentsResponseBody) SetCode(v string) *ListPersonalDirectoryContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) SetItems(v []*ListPersonalDirectoryContentsResponseBodyItems) *ListPersonalDirectoryContentsResponseBody {
	s.Items = v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) SetMessage(v string) *ListPersonalDirectoryContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) SetPage(v int64) *ListPersonalDirectoryContentsResponseBody {
	s.Page = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) SetPageSize(v int64) *ListPersonalDirectoryContentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) SetRequestId(v string) *ListPersonalDirectoryContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) SetTotal(v int64) *ListPersonalDirectoryContentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBody) Validate() error {
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

type ListPersonalDirectoryContentsResponseBodyItems struct {
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
	// The signing record ID.
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The item type.
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
	ObjectBindings []*ListPersonalDirectoryContentsResponseBodyItemsObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The number of resources in the FAILED state. This field is returned only when the top-level KB directory list is queried.
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// The knowledge base affiliation type. Valid values: aliding_kb_doc (DingTalk knowledge base document), normal (common knowledge).
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The number of resources in the READY state. This field is returned only when the top-level KB directory list is queried.
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
	// The total number of resources under the directory and its subdirectories. This field is returned only when the top-level KB directory list is queried.
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// The data source type.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListPersonalDirectoryContentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetObjectBindings() []*ListPersonalDirectoryContentsResponseBodyItemsObjectBindings {
	return s.ObjectBindings
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetSourceKind() *string {
	return s.SourceKind
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetCreatorName(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetDescription(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetDirectoryKind(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetGmtCreate(v int64) *ListPersonalDirectoryContentsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetGmtModified(v int64) *ListPersonalDirectoryContentsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetItemId(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetItemType(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetName(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetObjectBindings(v []*ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) *ListPersonalDirectoryContentsResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetSourceFailedCount(v int64) *ListPersonalDirectoryContentsResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetSourceKind(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.SourceKind = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetSourceReadyCount(v int64) *ListPersonalDirectoryContentsResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetSourceStatus(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetSourceTotalCount(v int64) *ListPersonalDirectoryContentsResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) SetSourceType(v string) *ListPersonalDirectoryContentsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItems) Validate() error {
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

type ListPersonalDirectoryContentsResponseBodyItemsObjectBindings struct {
	// The bound object ID.
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
	// The bound object type, such as customer or project.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The display name of the object type, such as "Customer". This value is resolved from the graph schema. The value is null when the cache is missed.
	//
	// example:
	//
	// string_value
	ObjectTypeName *string `json:"objectTypeName,omitempty" xml:"objectTypeName,omitempty"`
}

func (s ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) GetObjectTypeName() *string {
	return s.ObjectTypeName
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) SetObjectId(v string) *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) SetObjectName(v string) *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings {
	s.ObjectName = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) SetObjectType(v string) *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) SetObjectTypeName(v string) *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings {
	s.ObjectTypeName = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponseBodyItemsObjectBindings) Validate() error {
	return dara.Validate(s)
}
