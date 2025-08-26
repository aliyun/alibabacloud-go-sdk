// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTicketInvoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeTicketInvoiceRequest
	GetImageURL() *string
}

type RecognizeTicketInvoiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTicketInvoice/RecognizeTicketInvoice1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTicketInvoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeTicketInvoiceRequest) SetImageURL(v string) *RecognizeTicketInvoiceRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeTicketInvoiceRequest) Validate() error {
	return dara.Validate(s)
}
