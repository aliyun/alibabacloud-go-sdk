// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelGalleryModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModels(v *ModelGalleryModel) *ListModelGalleryModelsResponseBody
	GetModels() *ModelGalleryModel
	SetRequestId(v string) *ListModelGalleryModelsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListModelGalleryModelsResponseBody
	GetTotalCount() *string
}

type ListModelGalleryModelsResponseBody struct {
	// example:
	//
	// []
	Models *ModelGalleryModel `json:"Models,omitempty" xml:"Models,omitempty"`
	// example:
	//
	// B6B54325-C98C-5937-87A3-2F96C07652EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelGalleryModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelGalleryModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelGalleryModelsResponseBody) GetModels() *ModelGalleryModel {
	return s.Models
}

func (s *ListModelGalleryModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelGalleryModelsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListModelGalleryModelsResponseBody) SetModels(v *ModelGalleryModel) *ListModelGalleryModelsResponseBody {
	s.Models = v
	return s
}

func (s *ListModelGalleryModelsResponseBody) SetRequestId(v string) *ListModelGalleryModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelGalleryModelsResponseBody) SetTotalCount(v string) *ListModelGalleryModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelGalleryModelsResponseBody) Validate() error {
	if s.Models != nil {
		if err := s.Models.Validate(); err != nil {
			return err
		}
	}
	return nil
}
