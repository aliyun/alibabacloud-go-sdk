// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGroupDirectoryResponseBody
	GetCode() *string
	SetItems(v []*ListGroupDirectoryResponseBodyItems) *ListGroupDirectoryResponseBody
	GetItems() []*ListGroupDirectoryResponseBodyItems
	SetMessage(v string) *ListGroupDirectoryResponseBody
	GetMessage() *string
	SetPage(v int64) *ListGroupDirectoryResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListGroupDirectoryResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListGroupDirectoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupDirectoryResponseBody
	GetTotalCount() *int64
}

type ListGroupDirectoryResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The immediate subdirectories and resources on the current page. The queried directory itself is not included, and results are not recursively expanded.
	//
	// example:
	//
	// []
	Items []*ListGroupDirectoryResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries after filtering and before pagination. This includes both physical content and referenced content that match the filter criteria.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGroupDirectoryResponseBody) GetItems() []*ListGroupDirectoryResponseBodyItems {
	return s.Items
}

func (s *ListGroupDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGroupDirectoryResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListGroupDirectoryResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupDirectoryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupDirectoryResponseBody) SetCode(v string) *ListGroupDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupDirectoryResponseBody) SetItems(v []*ListGroupDirectoryResponseBodyItems) *ListGroupDirectoryResponseBody {
	s.Items = v
	return s
}

func (s *ListGroupDirectoryResponseBody) SetMessage(v string) *ListGroupDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupDirectoryResponseBody) SetPage(v int64) *ListGroupDirectoryResponseBody {
	s.Page = &v
	return s
}

func (s *ListGroupDirectoryResponseBody) SetPageSize(v int64) *ListGroupDirectoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGroupDirectoryResponseBody) SetRequestId(v string) *ListGroupDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupDirectoryResponseBody) SetTotalCount(v int64) *ListGroupDirectoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupDirectoryResponseBody) Validate() error {
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

type ListGroupDirectoryResponseBodyItems struct {
	// The name of the directory creator or resource submitter.
	//
	// example:
	//
	// example
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The directory ownership category. This follows the service output, such as normal.
	//
	// example:
	//
	// example
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// The directory type. Physical directories within the space have a value of GROUP. Reference directories retain their original type.
	//
	// example:
	//
	// example
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
	// The creation timestamp, in seconds.
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification timestamp, in seconds.
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The directoryId of a directory or the sourceId of a resource.
	//
	// example:
	//
	// source_example
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The content type. Valid values: directory (subdirectory) and resource.
	//
	// example:
	//
	// resource
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The name of the last modifier.
	//
	// example:
	//
	// example
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// The content name.
	//
	// example:
	//
	// Project Resources
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of resource object bindings. This may be empty if metadata is missing or for referenced resources.
	//
	// example:
	//
	// []
	ObjectBindings []*ListGroupDirectoryResponseBodyItemsObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// Indicates whether the content is a read-only reference. A value of false does not indicate write permissions. Write operations still require creator or space administrator permissions.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The resource ownership category. This follows the service output.
	//
	// example:
	//
	// example
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The resource parsing status. This field has a value only for resource items.
	//
	// example:
	//
	// READY
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// The resource type. This field has a value only for resource items. The type display rules of the service are used.
	//
	// example:
	//
	// example
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListGroupDirectoryResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoryResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListGroupDirectoryResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListGroupDirectoryResponseBodyItems) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *ListGroupDirectoryResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListGroupDirectoryResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListGroupDirectoryResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListGroupDirectoryResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListGroupDirectoryResponseBodyItems) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListGroupDirectoryResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListGroupDirectoryResponseBodyItems) GetObjectBindings() []*ListGroupDirectoryResponseBodyItemsObjectBindings {
	return s.ObjectBindings
}

func (s *ListGroupDirectoryResponseBodyItems) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListGroupDirectoryResponseBodyItems) GetSourceKind() *string {
	return s.SourceKind
}

func (s *ListGroupDirectoryResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListGroupDirectoryResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGroupDirectoryResponseBodyItems) SetCreatorName(v string) *ListGroupDirectoryResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetDirectoryKind(v string) *ListGroupDirectoryResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetDirectoryType(v string) *ListGroupDirectoryResponseBodyItems {
	s.DirectoryType = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetGmtCreate(v int64) *ListGroupDirectoryResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetGmtModified(v int64) *ListGroupDirectoryResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetItemId(v string) *ListGroupDirectoryResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetItemType(v string) *ListGroupDirectoryResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetModifierName(v string) *ListGroupDirectoryResponseBodyItems {
	s.ModifierName = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetName(v string) *ListGroupDirectoryResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetObjectBindings(v []*ListGroupDirectoryResponseBodyItemsObjectBindings) *ListGroupDirectoryResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetReadOnly(v bool) *ListGroupDirectoryResponseBodyItems {
	s.ReadOnly = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetSourceKind(v string) *ListGroupDirectoryResponseBodyItems {
	s.SourceKind = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetSourceStatus(v string) *ListGroupDirectoryResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) SetSourceType(v string) *ListGroupDirectoryResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItems) Validate() error {
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

type ListGroupDirectoryResponseBodyItemsObjectBindings struct {
	// The name of the knowledge graph to which the binding belongs.
	//
	// example:
	//
	// example
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The business ID of the object.
	//
	// example:
	//
	// example
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The display name of the object.
	//
	// example:
	//
	// example
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The object type.
	//
	// example:
	//
	// example
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The display name of the object type.
	//
	// example:
	//
	// example
	ObjectTypeName *string `json:"objectTypeName,omitempty" xml:"objectTypeName,omitempty"`
}

func (s ListGroupDirectoryResponseBodyItemsObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoryResponseBodyItemsObjectBindings) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) GetObjectTypeName() *string {
	return s.ObjectTypeName
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) SetGraphName(v string) *ListGroupDirectoryResponseBodyItemsObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) SetObjectId(v string) *ListGroupDirectoryResponseBodyItemsObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) SetObjectName(v string) *ListGroupDirectoryResponseBodyItemsObjectBindings {
	s.ObjectName = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) SetObjectType(v string) *ListGroupDirectoryResponseBodyItemsObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) SetObjectTypeName(v string) *ListGroupDirectoryResponseBodyItemsObjectBindings {
	s.ObjectTypeName = &v
	return s
}

func (s *ListGroupDirectoryResponseBodyItemsObjectBindings) Validate() error {
	return dara.Validate(s)
}
