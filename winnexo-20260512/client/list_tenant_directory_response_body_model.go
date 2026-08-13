// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTenantDirectoryResponseBody
	GetCode() *string
	SetItems(v []*ListTenantDirectoryResponseBodyItems) *ListTenantDirectoryResponseBody
	GetItems() []*ListTenantDirectoryResponseBodyItems
	SetMessage(v string) *ListTenantDirectoryResponseBody
	GetMessage() *string
	SetPage(v int64) *ListTenantDirectoryResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListTenantDirectoryResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListTenantDirectoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListTenantDirectoryResponseBody
	GetTotalCount() *int64
}

type ListTenantDirectoryResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code  *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Items []*ListTenantDirectoryResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 内容总数
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTenantDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTenantDirectoryResponseBody) GetItems() []*ListTenantDirectoryResponseBodyItems {
	return s.Items
}

func (s *ListTenantDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTenantDirectoryResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListTenantDirectoryResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTenantDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantDirectoryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTenantDirectoryResponseBody) SetCode(v string) *ListTenantDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantDirectoryResponseBody) SetItems(v []*ListTenantDirectoryResponseBodyItems) *ListTenantDirectoryResponseBody {
	s.Items = v
	return s
}

func (s *ListTenantDirectoryResponseBody) SetMessage(v string) *ListTenantDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantDirectoryResponseBody) SetPage(v int64) *ListTenantDirectoryResponseBody {
	s.Page = &v
	return s
}

func (s *ListTenantDirectoryResponseBody) SetPageSize(v int64) *ListTenantDirectoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTenantDirectoryResponseBody) SetRequestId(v string) *ListTenantDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantDirectoryResponseBody) SetTotalCount(v int64) *ListTenantDirectoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTenantDirectoryResponseBody) Validate() error {
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

type ListTenantDirectoryResponseBodyItems struct {
	// 创建人名称
	//
	// example:
	//
	// string_value
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 目录描述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间戳
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间戳
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 目录 ID 或资源 ID
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 内容类型：directory 或 resource
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
	Name           *string                  `json:"name,omitempty" xml:"name,omitempty"`
	ObjectBindings []map[string]interface{} `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// 根知识库下失败资源数
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// 根知识库下成功资源数
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// 资源解析状态
	//
	// example:
	//
	// string_value
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// 根知识库下资源总数
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// 资源类型
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListTenantDirectoryResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDirectoryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTenantDirectoryResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListTenantDirectoryResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListTenantDirectoryResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListTenantDirectoryResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListTenantDirectoryResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListTenantDirectoryResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListTenantDirectoryResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListTenantDirectoryResponseBodyItems) GetObjectBindings() []map[string]interface{} {
	return s.ObjectBindings
}

func (s *ListTenantDirectoryResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListTenantDirectoryResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListTenantDirectoryResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListTenantDirectoryResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListTenantDirectoryResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTenantDirectoryResponseBodyItems) SetCreatorName(v string) *ListTenantDirectoryResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetDescription(v string) *ListTenantDirectoryResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetGmtCreate(v int64) *ListTenantDirectoryResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetGmtModified(v int64) *ListTenantDirectoryResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetItemId(v string) *ListTenantDirectoryResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetItemType(v string) *ListTenantDirectoryResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetName(v string) *ListTenantDirectoryResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetObjectBindings(v []map[string]interface{}) *ListTenantDirectoryResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetSourceFailedCount(v int64) *ListTenantDirectoryResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetSourceReadyCount(v int64) *ListTenantDirectoryResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetSourceStatus(v string) *ListTenantDirectoryResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetSourceTotalCount(v int64) *ListTenantDirectoryResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) SetSourceType(v string) *ListTenantDirectoryResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListTenantDirectoryResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
