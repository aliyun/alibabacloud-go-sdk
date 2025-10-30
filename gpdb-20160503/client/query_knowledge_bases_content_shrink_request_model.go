// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryKnowledgeBasesContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *QueryKnowledgeBasesContentShrinkRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryKnowledgeBasesContentShrinkRequest
	GetDBInstanceId() *string
	SetMergeMethod(v string) *QueryKnowledgeBasesContentShrinkRequest
	GetMergeMethod() *string
	SetMergeMethodArgsShrink(v string) *QueryKnowledgeBasesContentShrinkRequest
	GetMergeMethodArgsShrink() *string
	SetOwnerId(v int64) *QueryKnowledgeBasesContentShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryKnowledgeBasesContentShrinkRequest
	GetRegionId() *string
	SetRerankFactor(v float64) *QueryKnowledgeBasesContentShrinkRequest
	GetRerankFactor() *float64
	SetSourceCollectionShrink(v string) *QueryKnowledgeBasesContentShrinkRequest
	GetSourceCollectionShrink() *string
	SetTopK(v int64) *QueryKnowledgeBasesContentShrinkRequest
	GetTopK() *int64
}

type QueryKnowledgeBasesContentShrinkRequest struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// RRF
	MergeMethod           *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	MergeMethodArgsShrink *string `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// This parameter is required.
	SourceCollectionShrink *string `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty"`
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s QueryKnowledgeBasesContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetMergeMethod() *string {
	return s.MergeMethod
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetMergeMethodArgsShrink() *string {
	return s.MergeMethodArgsShrink
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetSourceCollectionShrink() *string {
	return s.SourceCollectionShrink
}

func (s *QueryKnowledgeBasesContentShrinkRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetContent(v string) *QueryKnowledgeBasesContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetDBInstanceId(v string) *QueryKnowledgeBasesContentShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetMergeMethod(v string) *QueryKnowledgeBasesContentShrinkRequest {
	s.MergeMethod = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetMergeMethodArgsShrink(v string) *QueryKnowledgeBasesContentShrinkRequest {
	s.MergeMethodArgsShrink = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetOwnerId(v int64) *QueryKnowledgeBasesContentShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetRegionId(v string) *QueryKnowledgeBasesContentShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetRerankFactor(v float64) *QueryKnowledgeBasesContentShrinkRequest {
	s.RerankFactor = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetSourceCollectionShrink(v string) *QueryKnowledgeBasesContentShrinkRequest {
	s.SourceCollectionShrink = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) SetTopK(v int64) *QueryKnowledgeBasesContentShrinkRequest {
	s.TopK = &v
	return s
}

func (s *QueryKnowledgeBasesContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
