// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRecommendationDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UploadRecommendationDataRequest
	GetRegionId() *string
	SetContent(v []*UploadRecommendationDataRequestContent) *UploadRecommendationDataRequest
	GetContent() []*UploadRecommendationDataRequestContent
	SetDataType(v string) *UploadRecommendationDataRequest
	GetDataType() *string
}

type UploadRecommendationDataRequest struct {
	RegionId *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Content  []*UploadRecommendationDataRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	DataType *string                                   `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s UploadRecommendationDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadRecommendationDataRequest) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadRecommendationDataRequest) GetContent() []*UploadRecommendationDataRequestContent {
	return s.Content
}

func (s *UploadRecommendationDataRequest) GetDataType() *string {
	return s.DataType
}

func (s *UploadRecommendationDataRequest) SetRegionId(v string) *UploadRecommendationDataRequest {
	s.RegionId = &v
	return s
}

func (s *UploadRecommendationDataRequest) SetContent(v []*UploadRecommendationDataRequestContent) *UploadRecommendationDataRequest {
	s.Content = v
	return s
}

func (s *UploadRecommendationDataRequest) SetDataType(v string) *UploadRecommendationDataRequest {
	s.DataType = &v
	return s
}

func (s *UploadRecommendationDataRequest) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadRecommendationDataRequestContent struct {
	Fields        *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s UploadRecommendationDataRequestContent) String() string {
	return dara.Prettify(s)
}

func (s UploadRecommendationDataRequestContent) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataRequestContent) GetFields() *string {
	return s.Fields
}

func (s *UploadRecommendationDataRequestContent) GetOperationType() *string {
	return s.OperationType
}

func (s *UploadRecommendationDataRequestContent) SetFields(v string) *UploadRecommendationDataRequestContent {
	s.Fields = &v
	return s
}

func (s *UploadRecommendationDataRequestContent) SetOperationType(v string) *UploadRecommendationDataRequestContent {
	s.OperationType = &v
	return s
}

func (s *UploadRecommendationDataRequestContent) Validate() error {
	return dara.Validate(s)
}
