// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckImageExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CheckImageExistsRequest
	GetInstanceName() *string
	SetPicName(v string) *CheckImageExistsRequest
	GetPicName() *string
	SetProductId(v string) *CheckImageExistsRequest
	GetProductId() *string
}

type CheckImageExistsRequest struct {
	// Image Search instance name. Supports up to 20 characters.
	//
	// If you have already purchased an Image Search instance, you can logon to the [Image Search console](https://imagesearch.console.aliyun.com/) to view it.
	//
	// If you have not purchased an Image Search instance, see [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// > The instance name here is not the instance ID. Please distinguish between them when using.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// Product ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s CheckImageExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckImageExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckImageExistsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CheckImageExistsRequest) GetPicName() *string {
	return s.PicName
}

func (s *CheckImageExistsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CheckImageExistsRequest) SetInstanceName(v string) *CheckImageExistsRequest {
	s.InstanceName = &v
	return s
}

func (s *CheckImageExistsRequest) SetPicName(v string) *CheckImageExistsRequest {
	s.PicName = &v
	return s
}

func (s *CheckImageExistsRequest) SetProductId(v string) *CheckImageExistsRequest {
	s.ProductId = &v
	return s
}

func (s *CheckImageExistsRequest) Validate() error {
	return dara.Validate(s)
}
