// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iImitatePhotoStyleAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *ImitatePhotoStyleAdvanceRequest
	GetImageURLObject() io.Reader
	SetStyleUrlObject(v io.Reader) *ImitatePhotoStyleAdvanceRequest
	GetStyleUrlObject() io.Reader
}

type ImitatePhotoStyleAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ImitatePhotoStyle/ImitatePhotoStyle1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ImitatePhotoStyle/ImitatePhotoStyle7.jpg
	StyleUrlObject io.Reader `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ImitatePhotoStyleAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ImitatePhotoStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *ImitatePhotoStyleAdvanceRequest) GetStyleUrlObject() io.Reader {
	return s.StyleUrlObject
}

func (s *ImitatePhotoStyleAdvanceRequest) SetImageURLObject(v io.Reader) *ImitatePhotoStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ImitatePhotoStyleAdvanceRequest) SetStyleUrlObject(v io.Reader) *ImitatePhotoStyleAdvanceRequest {
	s.StyleUrlObject = v
	return s
}

func (s *ImitatePhotoStyleAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
