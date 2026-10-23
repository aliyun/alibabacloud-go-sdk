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
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of historical artifacts.
	//
	// example:
	//
	// example
	Items []*ListWorkItemArtifactsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of returned artifacts.
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
	// The artifact ID.
	//
	// example:
	//
	// example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// The artifact key.
	//
	// example:
	//
	// example
	ArtifactKey *string `json:"artifactKey,omitempty" xml:"artifactKey,omitempty"`
	// The node attempt sequence number.
	//
	// example:
	//
	// 1
	AttemptSeq *int64 `json:"attemptSeq,omitempty" xml:"attemptSeq,omitempty"`
	// The content type.
	//
	// example:
	//
	// example
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// example
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The node key.
	//
	// example:
	//
	// example
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
	// The plan ID.
	//
	// example:
	//
	// example
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The reference information.
	//
	// example:
	//
	// example
	Ref *ListWorkItemArtifactsResponseBodyItemsRef `json:"ref,omitempty" xml:"ref,omitempty" type:"Struct"`
	// The run ID.
	//
	// example:
	//
	// example
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The SHA-256 hash of the content.
	//
	// example:
	//
	// example
	Sha256 *string `json:"sha256,omitempty" xml:"sha256,omitempty"`
	// The artifact size in bytes.
	//
	// example:
	//
	// 1
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// The artifact summary.
	//
	// example:
	//
	// example
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The artifact title.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The session turn sequence number.
	//
	// example:
	//
	// 1
	TurnSeq *int64 `json:"turnSeq,omitempty" xml:"turnSeq,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// example
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The work item ID.
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
	// The artifact ID.
	//
	// example:
	//
	// example
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The reference data type.
	//
	// example:
	//
	// WORK_ITEM_ARTIFACT
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The OpenAPI artifact access endpoint.
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
