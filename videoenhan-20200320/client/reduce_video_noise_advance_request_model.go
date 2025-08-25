// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iReduceVideoNoiseAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrlObject(v io.Reader) *ReduceVideoNoiseAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type ReduceVideoNoiseAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ReduceVideoNoiseAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReduceVideoNoiseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *ReduceVideoNoiseAdvanceRequest) SetVideoUrlObject(v io.Reader) *ReduceVideoNoiseAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *ReduceVideoNoiseAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
