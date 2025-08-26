// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizePdfAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileURLObject(v io.Reader) *RecognizePdfAdvanceRequest
	GetFileURLObject() io.Reader
}

type RecognizePdfAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/ocr/xxxx.pdf
	FileURLObject io.Reader `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizePdfAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePdfAdvanceRequest) GetFileURLObject() io.Reader {
	return s.FileURLObject
}

func (s *RecognizePdfAdvanceRequest) SetFileURLObject(v io.Reader) *RecognizePdfAdvanceRequest {
	s.FileURLObject = v
	return s
}

func (s *RecognizePdfAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
