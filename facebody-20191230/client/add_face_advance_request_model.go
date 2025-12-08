// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAddFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *AddFaceAdvanceRequest
	GetDbName() *string
	SetEntityId(v string) *AddFaceAdvanceRequest
	GetEntityId() *string
	SetExtraData(v string) *AddFaceAdvanceRequest
	GetExtraData() *string
	SetImageUrlObject(v io.Reader) *AddFaceAdvanceRequest
	GetImageUrlObject() io.Reader
	SetQualityScoreThreshold(v float32) *AddFaceAdvanceRequest
	GetQualityScoreThreshold() *float32
	SetSimilarityScoreThresholdBetweenEntity(v float32) *AddFaceAdvanceRequest
	GetSimilarityScoreThresholdBetweenEntity() *float32
	SetSimilarityScoreThresholdInEntity(v float32) *AddFaceAdvanceRequest
	GetSimilarityScoreThresholdInEntity() *float32
}

type AddFaceAdvanceRequest struct {
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
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s AddFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceAdvanceRequest) GetDbName() *string {
	return s.DbName
}

func (s *AddFaceAdvanceRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *AddFaceAdvanceRequest) GetExtraData() *string {
	return s.ExtraData
}

func (s *AddFaceAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *AddFaceAdvanceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *AddFaceAdvanceRequest) GetSimilarityScoreThresholdBetweenEntity() *float32 {
	return s.SimilarityScoreThresholdBetweenEntity
}

func (s *AddFaceAdvanceRequest) GetSimilarityScoreThresholdInEntity() *float32 {
	return s.SimilarityScoreThresholdInEntity
}

func (s *AddFaceAdvanceRequest) SetDbName(v string) *AddFaceAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetEntityId(v string) *AddFaceAdvanceRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetExtraData(v string) *AddFaceAdvanceRequest {
	s.ExtraData = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetImageUrlObject(v io.Reader) *AddFaceAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *AddFaceAdvanceRequest) SetQualityScoreThreshold(v float32) *AddFaceAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *AddFaceAdvanceRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetSimilarityScoreThresholdInEntity(v float32) *AddFaceAdvanceRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

func (s *AddFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
