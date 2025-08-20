// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentSkinAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetURLObject(v io.Reader) *SegmentSkinAdvanceRequest
	GetURLObject() io.Reader
}

type SegmentSkinAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSkin/SegmentSkin2.jpg
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentSkinAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkinAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkinAdvanceRequest) GetURLObject() io.Reader {
	return s.URLObject
}

func (s *SegmentSkinAdvanceRequest) SetURLObject(v io.Reader) *SegmentSkinAdvanceRequest {
	s.URLObject = v
	return s
}

func (s *SegmentSkinAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
