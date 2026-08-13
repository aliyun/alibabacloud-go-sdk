// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVisibleKnowledgeBasesResponseBody
	GetCode() *string
	SetItems(v []*ListVisibleKnowledgeBasesResponseBodyItems) *ListVisibleKnowledgeBasesResponseBody
	GetItems() []*ListVisibleKnowledgeBasesResponseBodyItems
	SetMessage(v string) *ListVisibleKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVisibleKnowledgeBasesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListVisibleKnowledgeBasesResponseBody
	GetTotal() *int64
}

type ListVisibleKnowledgeBasesResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code  *string                                       `json:"code,omitempty" xml:"code,omitempty"`
	Items []*ListVisibleKnowledgeBasesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 返回条数（不分页，等于 len(items)）
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListVisibleKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetItems() []*ListVisibleKnowledgeBasesResponseBodyItems {
	return s.Items
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetCode(v string) *ListVisibleKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetItems(v []*ListVisibleKnowledgeBasesResponseBodyItems) *ListVisibleKnowledgeBasesResponseBody {
	s.Items = v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetMessage(v string) *ListVisibleKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetRequestId(v string) *ListVisibleKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetTotal(v int64) *ListVisibleKnowledgeBasesResponseBody {
	s.Total = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) Validate() error {
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

type ListVisibleKnowledgeBasesResponseBodyItems struct {
	// 目录创建者姓名（来自 rbj_user_tenant_mapping.user_display_name）
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
	// 目录唯一标识（租户内唯一）
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件 OSS URL
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 目录及子目录下状态为 FAILED 的资源数
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// 目录及子目录下状态为 READY 的资源数
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// 目录及子目录下的资源总数
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
}

func (s ListVisibleKnowledgeBasesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetPath() *string {
	return s.Path
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetCreatorName(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetDescription(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetDirectoryId(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.DirectoryId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetGmtCreate(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetGmtModified(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetName(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetPath(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.Path = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetSourceFailedCount(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetSourceReadyCount(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetSourceTotalCount(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
