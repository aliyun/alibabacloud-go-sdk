// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentSkinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetURL(v string) *SegmentSkinRequest
	GetURL() *string
}

type SegmentSkinRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSkin/SegmentSkin2.jpg
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentSkinRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkinRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkinRequest) GetURL() *string {
	return s.URL
}

func (s *SegmentSkinRequest) SetURL(v string) *SegmentSkinRequest {
	s.URL = &v
	return s
}

func (s *SegmentSkinRequest) Validate() error {
	return dara.Validate(s)
}
