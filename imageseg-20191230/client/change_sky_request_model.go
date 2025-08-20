// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSkyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *ChangeSkyRequest
	GetImageURL() *string
	SetReplaceImageURL(v string) *ChangeSkyRequest
	GetReplaceImageURL() *string
}

type ChangeSkyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky6.jpg
	ReplaceImageURL *string `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty"`
}

func (s ChangeSkyRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeSkyRequest) GoString() string {
	return s.String()
}

func (s *ChangeSkyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *ChangeSkyRequest) GetReplaceImageURL() *string {
	return s.ReplaceImageURL
}

func (s *ChangeSkyRequest) SetImageURL(v string) *ChangeSkyRequest {
	s.ImageURL = &v
	return s
}

func (s *ChangeSkyRequest) SetReplaceImageURL(v string) *ChangeSkyRequest {
	s.ReplaceImageURL = &v
	return s
}

func (s *ChangeSkyRequest) Validate() error {
	return dara.Validate(s)
}
