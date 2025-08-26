// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeTrainTicketAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeTrainTicketAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeTrainTicketAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTrainTicket/RecognizeTrainTicket3.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTrainTicketAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTrainTicketAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeTrainTicketAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTrainTicketAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTrainTicketAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
