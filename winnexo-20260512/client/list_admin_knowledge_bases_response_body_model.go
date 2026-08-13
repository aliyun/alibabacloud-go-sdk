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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code  *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Items []*ListAdminKnowledgeBasesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 符合条件的总数（应用 keyword/sourceTypes 后，分页前）
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
	// 目录创建者姓名（仅 KB 顶层目录列表时返回）
	//
	// example:
	//
	// string_value
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 描述（仅 KB 顶层目录列表时返回）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目录归属类型（itemType=directory 时有值）：normal / aliding_kb_root / aliding_kb_internal
	//
	// example:
	//
	// string_value
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// 创建时间戳（毫秒）
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间戳（毫秒）
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 唯一标识（itemType=directory 时为 directory_id；itemType=resource 时为 source_id）
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 类型：directory / resource
	//
	// example:
	//
	// string_value
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name           *string                                                   `json:"name,omitempty" xml:"name,omitempty"`
	ObjectBindings []*ListAdminKnowledgeBasesResponseBodyItemsObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// 状态为 FAILED 的资源数（仅 KB 顶层目录列表时返回）
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// 资源归属类型（itemType=resource 时有值）：aliding_kb_doc / normal
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// 状态为 READY 的资源数（仅 KB 顶层目录列表时返回）
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// 资源状态（itemType=resource 时有值）
	//
	// example:
	//
	// string_value
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// 目录及子目录下资源总数（仅 KB 顶层目录列表时返回）
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// 资源类型（itemType=resource 时有值）
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
	// 对象归属的语义图谱名（object_id 在该 graph 下唯一）
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 对象唯一 ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象显示名（如客户名称），由图谱 schema 解析；缓存缺失时为 null
	//
	// example:
	//
	// string_value
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// 对象类型（如 customer / opportunity），对应图谱 schema 中的 object_type
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 对象类型显示名（如"客户"），由图谱 schema 解析；缓存缺失时为 null
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
