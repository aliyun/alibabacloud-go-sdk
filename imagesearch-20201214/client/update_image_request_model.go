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
	// The custom content. The content can be up to 4,096 characters in length.
	//
	// >This field is returned when you call the "<props="china">[SearchImageByPic](https://help.aliyun.com/document_detail/202282.html)<props="intl">[SearchImageByPic](https://www.alibabacloud.com/help/zh/image-search/latest/updateimage)" operation. For example, you can add text such as image descriptions.
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, go to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, see [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure to distinguish between them.
	//
	// This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The integer attribute. This attribute can be used to filter query results. This field is returned in query results.
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// The integer attribute. This attribute can be used to filter query results. This field is returned in query results.
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	// The integer attribute. This attribute can be used to filter query results. This field is returned in query results.
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	// The integer attribute. This attribute can be used to filter query results. This field is returned in query results.
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The image name. The name can be up to 256 characters in length.
	//
	// > - The combination of ProductId and PicName uniquely identifies an image.
	//
	// - If you add an image multiple times with the same ProductId and PicName, the most recently added image takes effect and the previously added images are replaced.
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The product ID. The ID can be up to 256 characters in length.
	//
	// >A product can have multiple images. You can customize the value of this parameter based on your business requirements. For example: top001, pants002.
	//
	// This parameter is required.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. It can be used to filter query results. This field is returned in query results.
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. It can be used to filter query results. This field is returned in query results.
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. It can be used to filter query results. This field is returned in query results.
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. It can be used to filter query results. This field is returned in query results.
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
