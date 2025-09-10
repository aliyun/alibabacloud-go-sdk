// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iClassifyCommodityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *ClassifyCommodityAdvanceRequest
	GetImageURLObject() io.Reader
}

type ClassifyCommodityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/goodstech/ClassifyCommodity/ClassifyCommodity1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyCommodityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ClassifyCommodityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *ClassifyCommodityAdvanceRequest) SetImageURLObject(v io.Reader) *ClassifyCommodityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ClassifyCommodityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
