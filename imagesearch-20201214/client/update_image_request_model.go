// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomContent(v string) *UpdateImageRequest
	GetCustomContent() *string
	SetInstanceName(v string) *UpdateImageRequest
	GetInstanceName() *string
	SetIntAttr(v int32) *UpdateImageRequest
	GetIntAttr() *int32
	SetIntAttr2(v int32) *UpdateImageRequest
	GetIntAttr2() *int32
	SetIntAttr3(v int32) *UpdateImageRequest
	GetIntAttr3() *int32
	SetIntAttr4(v int32) *UpdateImageRequest
	GetIntAttr4() *int32
	SetPicName(v string) *UpdateImageRequest
	GetPicName() *string
	SetProductId(v string) *UpdateImageRequest
	GetProductId() *string
	SetStrAttr(v string) *UpdateImageRequest
	GetStrAttr() *string
	SetStrAttr2(v string) *UpdateImageRequest
	GetStrAttr2() *string
	SetStrAttr3(v string) *UpdateImageRequest
	GetStrAttr3() *string
	SetStrAttr4(v string) *UpdateImageRequest
	GetStrAttr4() *string
}

type UpdateImageRequest struct {
	// The user-defined content. The value can be up to 4,096 characters in length.
	//
	// >  If you set this parameter, the response includes this parameter and its value. You can add text such as an image description.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The attribute, which is an integer. The attribute can be used to filter images when you search for images. If you set this parameter, the response includes this parameter and its value.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 20
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The name of the image. The name can be up to 512 characters in length.
	//
	// > 	- An image is uniquely identified by the values of the ProductId and PicName parameters.
	//
	// >	- If you add an image whose product ID (ProductId) and image name (PicName) are the same as those of an existing image, the newly added image overwrites the existing image.
	//
	// This parameter is required.
	//
	// example:
	//
	// namexxx.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product. The ID can be up to 512 characters in length.
	//
	// >  A product may have multiple images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The attribute, which is a string. The value can be up to 128 characters in length. The attribute can be used to filter images. If you set this parameter, the response includes this parameter and its value.
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *UpdateImageRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateImageRequest) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *UpdateImageRequest) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *UpdateImageRequest) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *UpdateImageRequest) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *UpdateImageRequest) GetPicName() *string {
	return s.PicName
}

func (s *UpdateImageRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateImageRequest) GetStrAttr() *string {
	return s.StrAttr
}

func (s *UpdateImageRequest) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *UpdateImageRequest) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *UpdateImageRequest) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *UpdateImageRequest) SetCustomContent(v string) *UpdateImageRequest {
	s.CustomContent = &v
	return s
}

func (s *UpdateImageRequest) SetInstanceName(v string) *UpdateImageRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr(v int32) *UpdateImageRequest {
	s.IntAttr = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr2(v int32) *UpdateImageRequest {
	s.IntAttr2 = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr3(v int32) *UpdateImageRequest {
	s.IntAttr3 = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr4(v int32) *UpdateImageRequest {
	s.IntAttr4 = &v
	return s
}

func (s *UpdateImageRequest) SetPicName(v string) *UpdateImageRequest {
	s.PicName = &v
	return s
}

func (s *UpdateImageRequest) SetProductId(v string) *UpdateImageRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr(v string) *UpdateImageRequest {
	s.StrAttr = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr2(v string) *UpdateImageRequest {
	s.StrAttr2 = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr3(v string) *UpdateImageRequest {
	s.StrAttr3 = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr4(v string) *UpdateImageRequest {
	s.StrAttr4 = &v
	return s
}

func (s *UpdateImageRequest) Validate() error {
	return dara.Validate(s)
}
