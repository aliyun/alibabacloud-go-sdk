// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeQrCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*RecognizeQrCodeRequestTasks) *RecognizeQrCodeRequest
	GetTasks() []*RecognizeQrCodeRequestTasks
}

type RecognizeQrCodeRequest struct {
	// 1
	//
	// This parameter is required.
	Tasks []*RecognizeQrCodeRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s RecognizeQrCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeRequest) GetTasks() []*RecognizeQrCodeRequestTasks {
	return s.Tasks
}

func (s *RecognizeQrCodeRequest) SetTasks(v []*RecognizeQrCodeRequestTasks) *RecognizeQrCodeRequest {
	s.Tasks = v
	return s
}

func (s *RecognizeQrCodeRequest) Validate() error {
	return dara.Validate(s)
}

type RecognizeQrCodeRequestTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeQrCode/RecognizeQrCode6.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQrCodeRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeRequestTasks) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeRequestTasks) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeQrCodeRequestTasks) SetImageURL(v string) *RecognizeQrCodeRequestTasks {
	s.ImageURL = &v
	return s
}

func (s *RecognizeQrCodeRequestTasks) Validate() error {
	return dara.Validate(s)
}
