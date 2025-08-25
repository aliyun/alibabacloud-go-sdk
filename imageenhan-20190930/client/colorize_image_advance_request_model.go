// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iColorizeImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *ColorizeImageAdvanceRequest
	GetImageURLObject() io.Reader
}

type ColorizeImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ColorizeImage/ColorizeImage1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ColorizeImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ColorizeImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ColorizeImageAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *ColorizeImageAdvanceRequest) SetImageURLObject(v io.Reader) *ColorizeImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ColorizeImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
