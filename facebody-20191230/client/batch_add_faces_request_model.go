// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *BatchAddFacesRequest
	GetDbName() *string
	SetEntityId(v string) *BatchAddFacesRequest
	GetEntityId() *string
	SetFaces(v []*BatchAddFacesRequestFaces) *BatchAddFacesRequest
	GetFaces() []*BatchAddFacesRequestFaces
	SetQualityScoreThreshold(v float32) *BatchAddFacesRequest
	GetQualityScoreThreshold() *float32
	SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesRequest
	GetSimilarityScoreThresholdBetweenEntity() *float32
	SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesRequest
	GetSimilarityScoreThresholdInEntity() *float32
}

type BatchAddFacesRequest struct {
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
	Faces []*BatchAddFacesRequestFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
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

func (s BatchAddFacesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFacesRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchAddFacesRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *BatchAddFacesRequest) GetFaces() []*BatchAddFacesRequestFaces {
	return s.Faces
}

func (s *BatchAddFacesRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *BatchAddFacesRequest) GetSimilarityScoreThresholdBetweenEntity() *float32 {
	return s.SimilarityScoreThresholdBetweenEntity
}

func (s *BatchAddFacesRequest) GetSimilarityScoreThresholdInEntity() *float32 {
	return s.SimilarityScoreThresholdInEntity
}

func (s *BatchAddFacesRequest) SetDbName(v string) *BatchAddFacesRequest {
	s.DbName = &v
	return s
}

func (s *BatchAddFacesRequest) SetEntityId(v string) *BatchAddFacesRequest {
	s.EntityId = &v
	return s
}

func (s *BatchAddFacesRequest) SetFaces(v []*BatchAddFacesRequestFaces) *BatchAddFacesRequest {
	s.Faces = v
	return s
}

func (s *BatchAddFacesRequest) SetQualityScoreThreshold(v float32) *BatchAddFacesRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *BatchAddFacesRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *BatchAddFacesRequest) SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

func (s *BatchAddFacesRequest) Validate() error {
	if s.Faces != nil {
		for _, item := range s.Faces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchAddFacesRequestFaces struct {
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/imgsearch/demo/1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BatchAddFacesRequestFaces) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesRequestFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesRequestFaces) GetExtraData() *string {
	return s.ExtraData
}

func (s *BatchAddFacesRequestFaces) GetImageURL() *string {
	return s.ImageURL
}

func (s *BatchAddFacesRequestFaces) SetExtraData(v string) *BatchAddFacesRequestFaces {
	s.ExtraData = &v
	return s
}

func (s *BatchAddFacesRequestFaces) SetImageURL(v string) *BatchAddFacesRequestFaces {
	s.ImageURL = &v
	return s
}

func (s *BatchAddFacesRequestFaces) Validate() error {
	return dara.Validate(s)
}
