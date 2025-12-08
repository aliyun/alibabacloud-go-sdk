// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectBodyCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectBodyCountRequest
	GetImageURL() *string
}

type DetectBodyCountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectBodyCount/DetectBodyCount3.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectBodyCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectBodyCountRequest) GoString() string {
	return s.String()
}

func (s *DetectBodyCountRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectBodyCountRequest) SetImageURL(v string) *DetectBodyCountRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectBodyCountRequest) Validate() error {
	return dara.Validate(s)
}
