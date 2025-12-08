// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchAddFacesAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *BatchAddFacesAdvanceRequest
	GetDbName() *string
	SetEntityId(v string) *BatchAddFacesAdvanceRequest
	GetEntityId() *string
	SetFaces(v []*BatchAddFacesAdvanceRequestFaces) *BatchAddFacesAdvanceRequest
	GetFaces() []*BatchAddFacesAdvanceRequestFaces
	SetQualityScoreThreshold(v float32) *BatchAddFacesAdvanceRequest
	GetQualityScoreThreshold() *float32
	SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesAdvanceRequest
	GetSimilarityScoreThresholdBetweenEntity() *float32
	SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesAdvanceRequest
	GetSimilarityScoreThresholdInEntity() *float32
}

type BatchAddFacesAdvanceRequest struct {
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
	Faces []*BatchAddFacesAdvanceRequestFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
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

func (s BatchAddFacesAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFacesAdvanceRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchAddFacesAdvanceRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *BatchAddFacesAdvanceRequest) GetFaces() []*BatchAddFacesAdvanceRequestFaces {
	return s.Faces
}

func (s *BatchAddFacesAdvanceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *BatchAddFacesAdvanceRequest) GetSimilarityScoreThresholdBetweenEntity() *float32 {
	return s.SimilarityScoreThresholdBetweenEntity
}

func (s *BatchAddFacesAdvanceRequest) GetSimilarityScoreThresholdInEntity() *float32 {
	return s.SimilarityScoreThresholdInEntity
}

func (s *BatchAddFacesAdvanceRequest) SetDbName(v string) *BatchAddFacesAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *BatchAddFacesAdvanceRequest) SetEntityId(v string) *BatchAddFacesAdvanceRequest {
	s.EntityId = &v
	return s
}

func (s *BatchAddFacesAdvanceRequest) SetFaces(v []*BatchAddFacesAdvanceRequestFaces) *BatchAddFacesAdvanceRequest {
	s.Faces = v
	return s
}

func (s *BatchAddFacesAdvanceRequest) SetQualityScoreThreshold(v float32) *BatchAddFacesAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *BatchAddFacesAdvanceRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesAdvanceRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *BatchAddFacesAdvanceRequest) SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesAdvanceRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

func (s *BatchAddFacesAdvanceRequest) Validate() error {
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

type BatchAddFacesAdvanceRequestFaces struct {
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/imgsearch/demo/1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BatchAddFacesAdvanceRequestFaces) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesAdvanceRequestFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesAdvanceRequestFaces) GetExtraData() *string {
	return s.ExtraData
}

func (s *BatchAddFacesAdvanceRequestFaces) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *BatchAddFacesAdvanceRequestFaces) SetExtraData(v string) *BatchAddFacesAdvanceRequestFaces {
	s.ExtraData = &v
	return s
}

func (s *BatchAddFacesAdvanceRequestFaces) SetImageURLObject(v io.Reader) *BatchAddFacesAdvanceRequestFaces {
	s.ImageURLObject = v
	return s
}

func (s *BatchAddFacesAdvanceRequestFaces) Validate() error {
	return dara.Validate(s)
}
