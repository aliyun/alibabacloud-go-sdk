// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTrainTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeTrainTicketRequest
	GetImageURL() *string
}

type RecognizeTrainTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTrainTicket/RecognizeTrainTicket3.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTrainTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTrainTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeTrainTicketRequest) SetImageURL(v string) *RecognizeTrainTicketRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeTrainTicketRequest) Validate() error {
	return dara.Validate(s)
}
