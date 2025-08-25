// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImitatePhotoStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *ImitatePhotoStyleRequest
	GetImageURL() *string
	SetStyleUrl(v string) *ImitatePhotoStyleRequest
	GetStyleUrl() *string
}

type ImitatePhotoStyleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ImitatePhotoStyle/ImitatePhotoStyle1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ImitatePhotoStyle/ImitatePhotoStyle7.jpg
	StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ImitatePhotoStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s ImitatePhotoStyleRequest) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *ImitatePhotoStyleRequest) GetStyleUrl() *string {
	return s.StyleUrl
}

func (s *ImitatePhotoStyleRequest) SetImageURL(v string) *ImitatePhotoStyleRequest {
	s.ImageURL = &v
	return s
}

func (s *ImitatePhotoStyleRequest) SetStyleUrl(v string) *ImitatePhotoStyleRequest {
	s.StyleUrl = &v
	return s
}

func (s *ImitatePhotoStyleRequest) Validate() error {
	return dara.Validate(s)
}
