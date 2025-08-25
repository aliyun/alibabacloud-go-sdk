// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iChangeImageSizeAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *ChangeImageSizeAdvanceRequest
	GetHeight() *int32
	SetUrlObject(v io.Reader) *ChangeImageSizeAdvanceRequest
	GetUrlObject() io.Reader
	SetWidth(v int32) *ChangeImageSizeAdvanceRequest
	GetWidth() *int32
}

type ChangeImageSizeAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeImageSizeAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeImageSizeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeAdvanceRequest) GetHeight() *int32 {
	return s.Height
}

func (s *ChangeImageSizeAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *ChangeImageSizeAdvanceRequest) GetWidth() *int32 {
	return s.Width
}

func (s *ChangeImageSizeAdvanceRequest) SetHeight(v int32) *ChangeImageSizeAdvanceRequest {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) SetUrlObject(v io.Reader) *ChangeImageSizeAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) SetWidth(v int32) *ChangeImageSizeAdvanceRequest {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
