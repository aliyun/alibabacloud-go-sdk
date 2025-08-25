// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRemoveImageWatermarkAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RemoveImageWatermarkAdvanceRequest
	GetImageURLObject() io.Reader
}

type RemoveImageWatermarkAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RemoveImageWatermark/RemoveImageWatermark3.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageWatermarkAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RemoveImageWatermarkAdvanceRequest) SetImageURLObject(v io.Reader) *RemoveImageWatermarkAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RemoveImageWatermarkAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
