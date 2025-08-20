// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefineMaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RefineMaskRequest
	GetImageURL() *string
	SetMaskImageURL(v string) *RefineMaskRequest
	GetMaskImageURL() *string
}

type RefineMaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/RefineMask/RefineMask6.jpg
	MaskImageURL *string `json:"MaskImageURL,omitempty" xml:"MaskImageURL,omitempty"`
}

func (s RefineMaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RefineMaskRequest) GoString() string {
	return s.String()
}

func (s *RefineMaskRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RefineMaskRequest) GetMaskImageURL() *string {
	return s.MaskImageURL
}

func (s *RefineMaskRequest) SetImageURL(v string) *RefineMaskRequest {
	s.ImageURL = &v
	return s
}

func (s *RefineMaskRequest) SetMaskImageURL(v string) *RefineMaskRequest {
	s.MaskImageURL = &v
	return s
}

func (s *RefineMaskRequest) Validate() error {
	return dara.Validate(s)
}
