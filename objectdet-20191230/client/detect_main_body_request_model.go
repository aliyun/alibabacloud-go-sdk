// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMainBodyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectMainBodyRequest
	GetImageURL() *string
}

type DetectMainBodyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectMainBody/DetectMainBody1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectMainBodyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectMainBodyRequest) GoString() string {
	return s.String()
}

func (s *DetectMainBodyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectMainBodyRequest) SetImageURL(v string) *DetectMainBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectMainBodyRequest) Validate() error {
	return dara.Validate(s)
}
