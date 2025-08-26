// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizePdfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *RecognizePdfRequest
	GetFileURL() *string
}

type RecognizePdfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/ocr/xxxx.pdf
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizePdfRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfRequest) GoString() string {
	return s.String()
}

func (s *RecognizePdfRequest) GetFileURL() *string {
	return s.FileURL
}

func (s *RecognizePdfRequest) SetFileURL(v string) *RecognizePdfRequest {
	s.FileURL = &v
	return s
}

func (s *RecognizePdfRequest) Validate() error {
	return dara.Validate(s)
}
