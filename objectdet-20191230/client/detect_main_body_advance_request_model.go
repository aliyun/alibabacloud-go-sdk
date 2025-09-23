// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectMainBodyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectMainBodyAdvanceRequest
	GetImageURLObject() io.Reader
}

type DetectMainBodyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectMainBody/DetectMainBody1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectMainBodyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectMainBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectMainBodyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectMainBodyAdvanceRequest) SetImageURLObject(v io.Reader) *DetectMainBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectMainBodyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
