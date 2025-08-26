// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeQrCodeAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*RecognizeQrCodeAdvanceRequestTasks) *RecognizeQrCodeAdvanceRequest
	GetTasks() []*RecognizeQrCodeAdvanceRequestTasks
}

type RecognizeQrCodeAdvanceRequest struct {
	// 1
	//
	// This parameter is required.
	Tasks []*RecognizeQrCodeAdvanceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s RecognizeQrCodeAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeAdvanceRequest) GetTasks() []*RecognizeQrCodeAdvanceRequestTasks {
	return s.Tasks
}

func (s *RecognizeQrCodeAdvanceRequest) SetTasks(v []*RecognizeQrCodeAdvanceRequestTasks) *RecognizeQrCodeAdvanceRequest {
	s.Tasks = v
	return s
}

func (s *RecognizeQrCodeAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type RecognizeQrCodeAdvanceRequestTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeQrCode/RecognizeQrCode6.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQrCodeAdvanceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeAdvanceRequestTasks) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeAdvanceRequestTasks) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeQrCodeAdvanceRequestTasks) SetImageURLObject(v io.Reader) *RecognizeQrCodeAdvanceRequestTasks {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeQrCodeAdvanceRequestTasks) Validate() error {
	return dara.Validate(s)
}
