// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentHeadAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentHeadAdvanceRequest
	GetImageURLObject() io.Reader
	SetReturnForm(v string) *SegmentHeadAdvanceRequest
	GetReturnForm() *string
}

type SegmentHeadAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHead/SegmentHead1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentHeadAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHeadAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHeadAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentHeadAdvanceRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentHeadAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHeadAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentHeadAdvanceRequest) SetReturnForm(v string) *SegmentHeadAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentHeadAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
