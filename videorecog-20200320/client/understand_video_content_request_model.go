// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnderstandVideoContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoURL(v string) *UnderstandVideoContentRequest
	GetVideoURL() *string
}

type UnderstandVideoContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/UnderstandVideoContent/UnderstandVideoContent1.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s UnderstandVideoContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UnderstandVideoContentRequest) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *UnderstandVideoContentRequest) SetVideoURL(v string) *UnderstandVideoContentRequest {
	s.VideoURL = &v
	return s
}

func (s *UnderstandVideoContentRequest) Validate() error {
	return dara.Validate(s)
}
