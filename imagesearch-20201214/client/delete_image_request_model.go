// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *DeleteImageRequest
	GetFilter() *string
	SetInstanceName(v string) *DeleteImageRequest
	GetInstanceName() *string
	SetIsDeleteByFilter(v bool) *DeleteImageRequest
	GetIsDeleteByFilter() *bool
	SetPicName(v string) *DeleteImageRequest
	GetPicName() *string
	SetProductId(v string) *DeleteImageRequest
	GetProductId() *string
}

type DeleteImageRequest struct {
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDeleteByFilter *bool   `json:"IsDeleteByFilter,omitempty" xml:"IsDeleteByFilter,omitempty"`
	// The name of the image.
	//
	// 	- If this parameter is not set, the system deletes all the images that correspond to the specified ProductId parameter.
	//
	// 	- If this parameter is set, the system deletes only the image that is specified by the ProductId and PicName parameters.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the commodity.
	//
	// >  A commodity may have multiple images.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) GetFilter() *string {
	return s.Filter
}

func (s *DeleteImageRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteImageRequest) GetIsDeleteByFilter() *bool {
	return s.IsDeleteByFilter
}

func (s *DeleteImageRequest) GetPicName() *string {
	return s.PicName
}

func (s *DeleteImageRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DeleteImageRequest) SetFilter(v string) *DeleteImageRequest {
	s.Filter = &v
	return s
}

func (s *DeleteImageRequest) SetInstanceName(v string) *DeleteImageRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteImageRequest) SetIsDeleteByFilter(v bool) *DeleteImageRequest {
	s.IsDeleteByFilter = &v
	return s
}

func (s *DeleteImageRequest) SetPicName(v string) *DeleteImageRequest {
	s.PicName = &v
	return s
}

func (s *DeleteImageRequest) SetProductId(v string) *DeleteImageRequest {
	s.ProductId = &v
	return s
}

func (s *DeleteImageRequest) Validate() error {
	return dara.Validate(s)
}
