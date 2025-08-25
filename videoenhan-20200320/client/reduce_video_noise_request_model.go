// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceVideoNoiseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrl(v string) *ReduceVideoNoiseRequest
	GetVideoUrl() *string
}

type ReduceVideoNoiseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ReduceVideoNoiseRequest) String() string {
	return dara.Prettify(s)
}

func (s ReduceVideoNoiseRequest) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *ReduceVideoNoiseRequest) SetVideoUrl(v string) *ReduceVideoNoiseRequest {
	s.VideoUrl = &v
	return s
}

func (s *ReduceVideoNoiseRequest) Validate() error {
	return dara.Validate(s)
}
