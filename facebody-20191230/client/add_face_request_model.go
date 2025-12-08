// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *AddFaceRequest
	GetDbName() *string
	SetEntityId(v string) *AddFaceRequest
	GetEntityId() *string
	SetExtraData(v string) *AddFaceRequest
	GetExtraData() *string
	SetImageUrl(v string) *AddFaceRequest
	GetImageUrl() *string
	SetQualityScoreThreshold(v float32) *AddFaceRequest
	GetQualityScoreThreshold() *float32
	SetSimilarityScoreThresholdBetweenEntity(v float32) *AddFaceRequest
	GetSimilarityScoreThresholdBetweenEntity() *float32
	SetSimilarityScoreThresholdInEntity(v float32) *AddFaceRequest
	GetSimilarityScoreThresholdInEntity() *float32
}

type AddFaceRequest struct {
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
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/imgsearch/demo/1.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s AddFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceRequest) GetDbName() *string {
	return s.DbName
}

func (s *AddFaceRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *AddFaceRequest) GetExtraData() *string {
	return s.ExtraData
}

func (s *AddFaceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *AddFaceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *AddFaceRequest) GetSimilarityScoreThresholdBetweenEntity() *float32 {
	return s.SimilarityScoreThresholdBetweenEntity
}

func (s *AddFaceRequest) GetSimilarityScoreThresholdInEntity() *float32 {
	return s.SimilarityScoreThresholdInEntity
}

func (s *AddFaceRequest) SetDbName(v string) *AddFaceRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceRequest) SetEntityId(v string) *AddFaceRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceRequest) SetExtraData(v string) *AddFaceRequest {
	s.ExtraData = &v
	return s
}

func (s *AddFaceRequest) SetImageUrl(v string) *AddFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *AddFaceRequest) SetQualityScoreThreshold(v float32) *AddFaceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *AddFaceRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *AddFaceRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *AddFaceRequest) SetSimilarityScoreThresholdInEntity(v float32) *AddFaceRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

func (s *AddFaceRequest) Validate() error {
	return dara.Validate(s)
}
