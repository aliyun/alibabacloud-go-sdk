// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody
	GetImages() []*ListImagesResponseBodyImages
	SetTotalCount(v int64) *ListImagesResponseBody
	GetTotalCount() *int64
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
}

type ListImagesResponseBody struct {
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetImages() []*ListImagesResponseBodyImages {
	return s.Images
}

func (s *ListImagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImagesResponseBodyImages struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// image-tzi7f9czc0cxs9s45t
	ImageId  *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUri *string                               `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels   []*ListImagesResponseBodyImagesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name     *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyImages) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetImageUri() *string {
	return s.ImageUri
}

func (s *ListImagesResponseBodyImages) GetLabels() []*ListImagesResponseBodyImagesLabels {
	return s.Labels
}

func (s *ListImagesResponseBodyImages) GetName() *string {
	return s.Name
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetGmtCreateTime(v string) *ListImagesResponseBodyImages {
	s.GmtCreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageUri(v string) *ListImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetLabels(v []*ListImagesResponseBodyImagesLabels) *ListImagesResponseBodyImages {
	s.Labels = v
	return s
}

func (s *ListImagesResponseBodyImages) SetName(v string) *ListImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImagesResponseBodyImagesLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImagesResponseBodyImagesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImagesLabels) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesLabels) GetKey() *string {
	return s.Key
}

func (s *ListImagesResponseBodyImagesLabels) GetValue() *string {
	return s.Value
}

func (s *ListImagesResponseBodyImagesLabels) SetKey(v string) *ListImagesResponseBodyImagesLabels {
	s.Key = &v
	return s
}

func (s *ListImagesResponseBodyImagesLabels) SetValue(v string) *ListImagesResponseBodyImagesLabels {
	s.Value = &v
	return s
}

func (s *ListImagesResponseBodyImagesLabels) Validate() error {
	return dara.Validate(s)
}
