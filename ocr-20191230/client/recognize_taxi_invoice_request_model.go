// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTaxiInvoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeTaxiInvoiceRequest
	GetImageURL() *string
}

type RecognizeTaxiInvoiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTaxiInvoice/RecognizeTaxiInvoice2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTaxiInvoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeTaxiInvoiceRequest) SetImageURL(v string) *RecognizeTaxiInvoiceRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeTaxiInvoiceRequest) Validate() error {
	return dara.Validate(s)
}
