// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesFromLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageIds(v string) *DeleteImagesFromLibRequest
	GetImageIds() *string
	SetLibId(v string) *DeleteImagesFromLibRequest
	GetLibId() *string
	SetRegionId(v string) *DeleteImagesFromLibRequest
	GetRegionId() *string
}

type DeleteImagesFromLibRequest struct {
	// example:
	//
	// [158794]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteImagesFromLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesFromLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesFromLibRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *DeleteImagesFromLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteImagesFromLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImagesFromLibRequest) SetImageIds(v string) *DeleteImagesFromLibRequest {
	s.ImageIds = &v
	return s
}

func (s *DeleteImagesFromLibRequest) SetLibId(v string) *DeleteImagesFromLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteImagesFromLibRequest) SetRegionId(v string) *DeleteImagesFromLibRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagesFromLibRequest) Validate() error {
	return dara.Validate(s)
}
