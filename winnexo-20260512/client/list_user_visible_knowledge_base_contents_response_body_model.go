// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBaseContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetCode() *string
	SetItems(v []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetItems() []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems
	SetMessage(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetTotalCount() *int64
}

type ListUserVisibleKnowledgeBaseContentsResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code  *string                                                  `json:"code,omitempty" xml:"code,omitempty"`
	Items []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 命中总数
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetItems() []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	return s.Items
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetCode(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetItems(v []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Items = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetMessage(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Page = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetRequestId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetTotalCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) Validate() error {
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

type ListUserVisibleKnowledgeBaseContentsResponseBodyItems struct {
	// 创建人名称
	//
	// example:
	//
	// 张三
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 知识库描述
	//
	// example:
	//
	// 产品资料与使用说明
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目录归属类型
	//
	// example:
	//
	// normal
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// 目录类型
	//
	// example:
	//
	// TENANT
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
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
	// 目录 ID 或资源 ID
	//
	// example:
	//
	// source_example
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 内容类型：directory 或 resource
	//
	// example:
	//
	// resource
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// 资源是否存在待审批的知识库提交记录
	//
	// example:
	//
	// false
	KbSubmissionPending *bool `json:"kbSubmissionPending,omitempty" xml:"kbSubmissionPending,omitempty"`
	// 更新人名称
	//
	// example:
	//
	// 李四
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// 目录或资源名称
	//
	// example:
	//
	// 产品说明.pdf
	Name           *string                  `json:"name,omitempty" xml:"name,omitempty"`
	ObjectBindings []map[string]interface{} `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// 知识库对数字员工的可见模式
	//
	// example:
	//
	// PUBLIC
	OoVisibilityMode *string `json:"ooVisibilityMode,omitempty" xml:"ooVisibilityMode,omitempty"`
	// 是否为只读关联内容
	//
	// example:
	//
	// false
	ReadOnly   *bool                                                              `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	ShareInfos []*ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos `json:"shareInfos,omitempty" xml:"shareInfos,omitempty" type:"Repeated"`
	// 是否已直接共享到企业知识库
	//
	// example:
	//
	// false
	Shared *bool `json:"shared,omitempty" xml:"shared,omitempty"`
	// 目录 FAILED 资源数
	//
	// example:
	//
	// 0
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// 资源归属类型
	//
	// example:
	//
	// normal
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// 目录 READY 资源数
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// 资源状态；本接口只返回 READY 资源
	//
	// example:
	//
	// READY
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// 目录资源总数
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// 资源类型，目录项为空
	//
	// example:
	//
	// FILE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetKbSubmissionPending() *bool {
	return s.KbSubmissionPending
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetObjectBindings() []map[string]interface{} {
	return s.ObjectBindings
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetShareInfos() []*ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	return s.ShareInfos
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetShared() *bool {
	return s.Shared
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceKind() *string {
	return s.SourceKind
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetCreatorName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetDescription(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetDirectoryKind(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetDirectoryType(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.DirectoryType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetGmtCreate(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetGmtModified(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetItemId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetItemType(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetKbSubmissionPending(v bool) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.KbSubmissionPending = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetModifierName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ModifierName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetObjectBindings(v []map[string]interface{}) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetOoVisibilityMode(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.OoVisibilityMode = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetReadOnly(v bool) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ReadOnly = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetShareInfos(v []*ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ShareInfos = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetShared(v bool) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Shared = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceFailedCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceKind(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceKind = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceReadyCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceStatus(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceTotalCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceType(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) Validate() error {
	if s.ShareInfos != nil {
		for _, item := range s.ShareInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos struct {
	// 知识库提交审批单 ID
	//
	// example:
	//
	// submission_example
	SubmissionId *string `json:"submissionId,omitempty" xml:"submissionId,omitempty"`
	// 提交人用户 ID
	//
	// example:
	//
	// 1
	SubmitterId *int64 `json:"submitterId,omitempty" xml:"submitterId,omitempty"`
	// 提交人名称
	//
	// example:
	//
	// 张三
	SubmitterName *string `json:"submitterName,omitempty" xml:"submitterName,omitempty"`
	// 目标目录 ID
	//
	// example:
	//
	// dir_target
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// 目标目录名称
	//
	// example:
	//
	// 共享目录
	TargetDirectoryName *string `json:"targetDirectoryName,omitempty" xml:"targetDirectoryName,omitempty"`
	// 目标企业知识库根目录 ID
	//
	// example:
	//
	// dir_kb_root
	TargetKbRootDirectoryId *string `json:"targetKbRootDirectoryId,omitempty" xml:"targetKbRootDirectoryId,omitempty"`
	// 目标企业知识库名称
	//
	// example:
	//
	// 产品知识库
	TargetKbRootDirectoryName *string `json:"targetKbRootDirectoryName,omitempty" xml:"targetKbRootDirectoryName,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetSubmitterId() *int64 {
	return s.SubmitterId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetSubmitterName() *string {
	return s.SubmitterName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetDirectoryName() *string {
	return s.TargetDirectoryName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetKbRootDirectoryId() *string {
	return s.TargetKbRootDirectoryId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetKbRootDirectoryName() *string {
	return s.TargetKbRootDirectoryName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetSubmissionId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.SubmissionId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetSubmitterId(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.SubmitterId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetSubmitterName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.SubmitterName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetDirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetDirectoryName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetDirectoryName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetKbRootDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetKbRootDirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetKbRootDirectoryName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetKbRootDirectoryName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) Validate() error {
	return dara.Validate(s)
}
