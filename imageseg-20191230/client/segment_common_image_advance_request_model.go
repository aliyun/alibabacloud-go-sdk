// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentCommonImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentCommonImageAdvanceRequest
	GetImageURLObject() io.Reader
	SetReturnForm(v string) *SegmentCommonImageAdvanceRequest
	GetReturnForm() *string
}

type SegmentCommonImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommonImage/SegmentCommonImage1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommonImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommonImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentCommonImageAdvanceRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentCommonImageAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentCommonImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentCommonImageAdvanceRequest) SetReturnForm(v string) *SegmentCommonImageAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentCommonImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
