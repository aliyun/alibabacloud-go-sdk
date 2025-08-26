// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeTaxiInvoiceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeTaxiInvoiceAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeTaxiInvoiceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTaxiInvoice/RecognizeTaxiInvoice2.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTaxiInvoiceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeTaxiInvoiceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTaxiInvoiceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTaxiInvoiceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
