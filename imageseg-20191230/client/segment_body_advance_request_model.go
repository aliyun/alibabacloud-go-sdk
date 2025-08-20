// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentBodyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentBodyAdvanceRequest
	GetImageURLObject() io.Reader
	SetReturnForm(v string) *SegmentBodyAdvanceRequest
	GetReturnForm() *string
}

type SegmentBodyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentBody/SegmentBody1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentBodyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentBodyAdvanceRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentBodyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentBodyAdvanceRequest) SetReturnForm(v string) *SegmentBodyAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentBodyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
