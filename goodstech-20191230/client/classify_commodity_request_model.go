// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClassifyCommodityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *ClassifyCommodityRequest
	GetImageURL() *string
}

type ClassifyCommodityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/goodstech/ClassifyCommodity/ClassifyCommodity1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyCommodityRequest) String() string {
	return dara.Prettify(s)
}

func (s ClassifyCommodityRequest) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *ClassifyCommodityRequest) SetImageURL(v string) *ClassifyCommodityRequest {
	s.ImageURL = &v
	return s
}

func (s *ClassifyCommodityRequest) Validate() error {
	return dara.Validate(s)
}
