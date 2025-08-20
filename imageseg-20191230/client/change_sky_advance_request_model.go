// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iChangeSkyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest
	GetImageURLObject() io.Reader
	SetReplaceImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest
	GetReplaceImageURLObject() io.Reader
}

type ChangeSkyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky2.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ChangeSky/ChangeSky6.jpg
	ReplaceImageURLObject io.Reader `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty"`
}

func (s ChangeSkyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeSkyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *ChangeSkyAdvanceRequest) GetReplaceImageURLObject() io.Reader {
	return s.ReplaceImageURLObject
}

func (s *ChangeSkyAdvanceRequest) SetImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ChangeSkyAdvanceRequest) SetReplaceImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest {
	s.ReplaceImageURLObject = v
	return s
}

func (s *ChangeSkyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
