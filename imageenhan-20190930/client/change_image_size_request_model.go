// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeImageSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *ChangeImageSizeRequest
	GetHeight() *int32
	SetUrl(v string) *ChangeImageSizeRequest
	GetUrl() *string
	SetWidth(v int32) *ChangeImageSizeRequest
	GetWidth() *int32
}

type ChangeImageSizeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 600
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ChangeImageSize/ChangeImageSize5.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeImageSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeImageSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeRequest) GetHeight() *int32 {
	return s.Height
}

func (s *ChangeImageSizeRequest) GetUrl() *string {
	return s.Url
}

func (s *ChangeImageSizeRequest) GetWidth() *int32 {
	return s.Width
}

func (s *ChangeImageSizeRequest) SetHeight(v int32) *ChangeImageSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeRequest) SetUrl(v string) *ChangeImageSizeRequest {
	s.Url = &v
	return s
}

func (s *ChangeImageSizeRequest) SetWidth(v int32) *ChangeImageSizeRequest {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeRequest) Validate() error {
	return dara.Validate(s)
}
