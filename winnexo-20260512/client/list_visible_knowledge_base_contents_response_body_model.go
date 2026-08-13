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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code  *string                                              `json:"code,omitempty" xml:"code,omitempty"`
	Items []*ListVisibleKnowledgeBaseContentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 总数（不分页前的命中行数）
	//
	// example:
	//
	// 1
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
	// 目录创建者姓名（仅根目录列表时返回；下钻场景为 null）
	//
	// example:
	//
	// string_value
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 知识库描述（仅根目录列表时返回；下钻场景为 null）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目录 KB 归属类型（itemType=directory 时有值）：aliding_kb_root / aliding_kb_internal / normal
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
	// 唯一标识（目录为 directoryId，资源为 sourceId）
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 类型: directory 或 resource
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
	Name           *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	ObjectBindings []*ListVisibleKnowledgeBaseContentsResponseBodyItemsObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// 目录下失败资源数（仅根目录列表时返回；下钻场景为 null）
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// Source KB 归属类型（itemType=resource 时有值）：aliding_kb_doc / normal
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// 目录下成功资源数（仅根目录列表时返回；下钻场景为 null）
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// 资源状态（itemType=resource 时有值；本接口固定按 READY 过滤）
	//
	// example:
	//
	// string_value
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// 目录下资源总数（含子目录，仅根目录列表时返回；下钻场景为 null）
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
