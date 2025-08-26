// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeQuotaInvoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeQuotaInvoiceRequest
	GetImageURL() *string
}

type RecognizeQuotaInvoiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeQuotaInvoice/RecognizeQuotaInvoice1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQuotaInvoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeQuotaInvoiceRequest) SetImageURL(v string) *RecognizeQuotaInvoiceRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeQuotaInvoiceRequest) Validate() error {
	return dara.Validate(s)
}
