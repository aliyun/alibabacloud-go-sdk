// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRefineMaskAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RefineMaskAdvanceRequest
	GetImageURLObject() io.Reader
	SetMaskImageURLObject(v io.Reader) *RefineMaskAdvanceRequest
	GetMaskImageURLObject() io.Reader
}

type RefineMaskAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask6.jpg
	MaskImageURLObject io.Reader `json:"MaskImageURL,omitempty" xml:"MaskImageURL,omitempty"`
}

func (s RefineMaskAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RefineMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RefineMaskAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RefineMaskAdvanceRequest) GetMaskImageURLObject() io.Reader {
	return s.MaskImageURLObject
}

func (s *RefineMaskAdvanceRequest) SetImageURLObject(v io.Reader) *RefineMaskAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RefineMaskAdvanceRequest) SetMaskImageURLObject(v io.Reader) *RefineMaskAdvanceRequest {
	s.MaskImageURLObject = v
	return s
}

func (s *RefineMaskAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
