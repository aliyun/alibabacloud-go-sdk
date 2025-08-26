// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeTicketInvoiceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeTicketInvoiceAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeTicketInvoiceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTicketInvoice/RecognizeTicketInvoice1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTicketInvoiceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeTicketInvoiceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTicketInvoiceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTicketInvoiceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
