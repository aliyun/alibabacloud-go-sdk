// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RemoveImageWatermarkRequest
	GetImageURL() *string
}

type RemoveImageWatermarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RemoveImageWatermark/RemoveImageWatermark3.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageWatermarkRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RemoveImageWatermarkRequest) SetImageURL(v string) *RemoveImageWatermarkRequest {
	s.ImageURL = &v
	return s
}

func (s *RemoveImageWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
