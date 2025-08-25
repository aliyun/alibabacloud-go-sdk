// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDynamicImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v string) *GenerateDynamicImageRequest
	GetOperation() *string
	SetUrl(v string) *GenerateDynamicImageRequest
	GetUrl() *string
}

type GenerateDynamicImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://vigen-invi.oss-cn-shanghai.aliyuncs.com/temp/xl/dynamicPhoto/viapi_test_images/xxxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateDynamicImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageRequest) GetOperation() *string {
	return s.Operation
}

func (s *GenerateDynamicImageRequest) GetUrl() *string {
	return s.Url
}

func (s *GenerateDynamicImageRequest) SetOperation(v string) *GenerateDynamicImageRequest {
	s.Operation = &v
	return s
}

func (s *GenerateDynamicImageRequest) SetUrl(v string) *GenerateDynamicImageRequest {
	s.Url = &v
	return s
}

func (s *GenerateDynamicImageRequest) Validate() error {
	return dara.Validate(s)
}
