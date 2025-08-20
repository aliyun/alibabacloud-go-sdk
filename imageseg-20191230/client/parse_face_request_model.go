// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *ParseFaceRequest
	GetImageURL() *string
}

type ParseFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ParseFace/ParseFace1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ParseFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ParseFaceRequest) GoString() string {
	return s.String()
}

func (s *ParseFaceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *ParseFaceRequest) SetImageURL(v string) *ParseFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *ParseFaceRequest) Validate() error {
	return dara.Validate(s)
}
