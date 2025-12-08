// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *BatchAddFacesShrinkRequest
	GetDbName() *string
	SetEntityId(v string) *BatchAddFacesShrinkRequest
	GetEntityId() *string
	SetFacesShrink(v string) *BatchAddFacesShrinkRequest
	GetFacesShrink() *string
	SetQualityScoreThreshold(v float32) *BatchAddFacesShrinkRequest
	GetQualityScoreThreshold() *float32
	SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesShrinkRequest
	GetSimilarityScoreThresholdBetweenEntity() *float32
	SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesShrinkRequest
	GetSimilarityScoreThresholdInEntity() *float32
}

type BatchAddFacesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// U1
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	FacesShrink *string `json:"Faces,omitempty" xml:"Faces,omitempty"`
	// example:
	//
	// 50.0
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
	// example:
	//
	// 50.0
	SimilarityScoreThresholdBetweenEntity *float32 `json:"SimilarityScoreThresholdBetweenEntity,omitempty" xml:"SimilarityScoreThresholdBetweenEntity,omitempty"`
	// example:
	//
	// 50.0
	SimilarityScoreThresholdInEntity *float32 `json:"SimilarityScoreThresholdInEntity,omitempty" xml:"SimilarityScoreThresholdInEntity,omitempty"`
}

func (s BatchAddFacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFacesShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchAddFacesShrinkRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *BatchAddFacesShrinkRequest) GetFacesShrink() *string {
	return s.FacesShrink
}

func (s *BatchAddFacesShrinkRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *BatchAddFacesShrinkRequest) GetSimilarityScoreThresholdBetweenEntity() *float32 {
	return s.SimilarityScoreThresholdBetweenEntity
}

func (s *BatchAddFacesShrinkRequest) GetSimilarityScoreThresholdInEntity() *float32 {
	return s.SimilarityScoreThresholdInEntity
}

func (s *BatchAddFacesShrinkRequest) SetDbName(v string) *BatchAddFacesShrinkRequest {
	s.DbName = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetEntityId(v string) *BatchAddFacesShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetFacesShrink(v string) *BatchAddFacesShrinkRequest {
	s.FacesShrink = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetQualityScoreThreshold(v float32) *BatchAddFacesShrinkRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesShrinkRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesShrinkRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
