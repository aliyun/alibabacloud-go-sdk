// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUnderstandVideoContentAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoURLObject(v io.Reader) *UnderstandVideoContentAdvanceRequest
	GetVideoURLObject() io.Reader
}

type UnderstandVideoContentAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/UnderstandVideoContent/UnderstandVideoContent1.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s UnderstandVideoContentAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnderstandVideoContentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *UnderstandVideoContentAdvanceRequest) SetVideoURLObject(v io.Reader) *UnderstandVideoContentAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *UnderstandVideoContentAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
