// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeBankCardAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeBankCardAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeBankCardAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeBankCard/yhk3.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBankCardAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBankCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeBankCardAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeBankCardAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeBankCardAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
