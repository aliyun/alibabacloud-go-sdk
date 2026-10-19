// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemArtifactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkItemArtifactsResponseBody
	GetCode() *string
	SetItems(v []*ListWorkItemArtifactsResponseBodyItems) *ListWorkItemArtifactsResponseBody
	GetItems() []*ListWorkItemArtifactsResponseBodyItems
	SetMessage(v string) *ListWorkItemArtifactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkItemArtifactsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListWorkItemArtifactsResponseBody
	GetTotalCount() *int64
}

type ListWorkItemArtifactsResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 历史产物列表
	//
	// example:
	//
	// example
	Items []*ListWorkItemArtifactsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 返回产物数量
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListWorkItemArtifactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkItemArtifactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkItemArtifactsResponseBody) GetItems() []*ListWorkItemArtifactsResponseBodyItems {
	return s.Items
}

func (s *ListWorkItemArtifactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkItemArtifactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkItemArtifactsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkItemArtifactsResponseBody) SetCode(v string) *ListWorkItemArtifactsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBody) SetItems(v []*ListWorkItemArtifactsResponseBodyItems) *ListWorkItemArtifactsResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkItemArtifactsResponseBody) SetMessage(v string) *ListWorkItemArtifactsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBody) SetRequestId(v string) *ListWorkItemArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBody) SetTotalCount(v int64) *ListWorkItemArtifactsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBody) Validate() error {
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

type ListWorkItemArtifactsResponseBodyItems struct {
	// 产物ID
	//
	// example:
	//
	// example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// 产物键
	//
	// example:
	//
	// example
	ArtifactKey *string `json:"artifactKey,omitempty" xml:"artifactKey,omitempty"`
	// 节点尝试序号
	//
	// example:
	//
	// 1
	AttemptSeq *int64 `json:"attemptSeq,omitempty" xml:"attemptSeq,omitempty"`
	// 内容类型
	//
	// example:
	//
	// example
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 创建时间(ISO8601)
	//
	// example:
	//
	// example
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 节点键
	//
	// example:
	//
	// example
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
	// 计划ID
	//
	// example:
	//
	// example
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// 引用信息
	//
	// example:
	//
	// example
	Ref *ListWorkItemArtifactsResponseBodyItemsRef `json:"ref,omitempty" xml:"ref,omitempty" type:"Struct"`
	// 运行ID
	//
	// example:
	//
	// example
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// 内容SHA-256
	//
	// example:
	//
	// example
	Sha256 *string `json:"sha256,omitempty" xml:"sha256,omitempty"`
	// 产物大小(字节)
	//
	// example:
	//
	// 1
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// 产物摘要
	//
	// example:
	//
	// example
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 产物标题
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 会话轮次序号
	//
	// example:
	//
	// 1
	TurnSeq *int64 `json:"turnSeq,omitempty" xml:"turnSeq,omitempty"`
	// 产物类型
	//
	// example:
	//
	// example
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 工作项ID
	//
	// example:
	//
	// example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ListWorkItemArtifactsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemArtifactsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetArtifactKey() *string {
	return s.ArtifactKey
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetAttemptSeq() *int64 {
	return s.AttemptSeq
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetContentType() *string {
	return s.ContentType
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetNodeKey() *string {
	return s.NodeKey
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetPlanId() *string {
	return s.PlanId
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetRef() *ListWorkItemArtifactsResponseBodyItemsRef {
	return s.Ref
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetRunId() *string {
	return s.RunId
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetSha256() *string {
	return s.Sha256
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetSummary() *string {
	return s.Summary
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetTurnSeq() *int64 {
	return s.TurnSeq
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *ListWorkItemArtifactsResponseBodyItems) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetArtifactId(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.ArtifactId = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetArtifactKey(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.ArtifactKey = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetAttemptSeq(v int64) *ListWorkItemArtifactsResponseBodyItems {
	s.AttemptSeq = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetContentType(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.ContentType = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetCreatedAt(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetNodeKey(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.NodeKey = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetPlanId(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.PlanId = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetRef(v *ListWorkItemArtifactsResponseBodyItemsRef) *ListWorkItemArtifactsResponseBodyItems {
	s.Ref = v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetRunId(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.RunId = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetSha256(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.Sha256 = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetSizeBytes(v int64) *ListWorkItemArtifactsResponseBodyItems {
	s.SizeBytes = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetSummary(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.Summary = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetTitle(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetTurnSeq(v int64) *ListWorkItemArtifactsResponseBodyItems {
	s.TurnSeq = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetType(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.Type = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) SetWorkItemId(v string) *ListWorkItemArtifactsResponseBodyItems {
	s.WorkItemId = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItems) Validate() error {
	if s.Ref != nil {
		if err := s.Ref.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkItemArtifactsResponseBodyItemsRef struct {
	// 产物ID
	//
	// example:
	//
	// example
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 引用类型
	//
	// example:
	//
	// WORK_ITEM_ARTIFACT
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// OpenAPI 产物访问入口
	//
	// example:
	//
	// example
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListWorkItemArtifactsResponseBodyItemsRef) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemArtifactsResponseBodyItemsRef) GoString() string {
	return s.String()
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) GetId() *string {
	return s.Id
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) GetKind() *string {
	return s.Kind
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) GetUrl() *string {
	return s.Url
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) SetId(v string) *ListWorkItemArtifactsResponseBodyItemsRef {
	s.Id = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) SetKind(v string) *ListWorkItemArtifactsResponseBodyItemsRef {
	s.Kind = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) SetUrl(v string) *ListWorkItemArtifactsResponseBodyItemsRef {
	s.Url = &v
	return s
}

func (s *ListWorkItemArtifactsResponseBodyItemsRef) Validate() error {
	return dara.Validate(s)
}
