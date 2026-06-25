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
	// The filter condition. The operators supported for int_attr include greater than (>), greater than or equal to (>=), less than (<), less than or equal to (<=), and equal to (=). The operators supported for str_attr include equal to (=) and not equal to (!=). Multiple conditions can be connected by using AND and OR.
	//
	// Examples:
	//
	// - int_attr >= 100.
	//
	// - str_attr != "value1".
	//
	// - int_attr = 1000 AND str_attr = "value1".
	//
	// >A maximum of 4096 characters are supported.
	//
	// example:
	//
	// int_attr=1000 AND str_attr="value1"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. The instance name must be unique within the same region. Make sure that you use the correct value.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Specifies whether to delete images based on the Filter parameter.
	//
	// > 1. If this parameter is set to true, images are deleted based on the Filter parameter, and Filter is required.
	//
	//  2. If this parameter is set to false, images are deleted based on ProductId or the combination of ProductId and PicName. ProductId is required.
	//
	// example:
	//
	// false
	IsDeleteByFilter *bool `json:"IsDeleteByFilter,omitempty" xml:"IsDeleteByFilter,omitempty"`
	// The image name.
	//
	//  - If you do not specify this parameter, all images under the specified ProductId are deleted.
	//
	//  - If you specify this parameter, the image specified by the combination of ProductId and PicName is deleted.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The product ID.
	//
	// > 1. A product can have multiple images. 2. If IsDeleteByFilter is set to false, this parameter is required.
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
