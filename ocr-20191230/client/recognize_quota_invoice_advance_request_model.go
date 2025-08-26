// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeQuotaInvoiceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeQuotaInvoiceAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeQuotaInvoiceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeQuotaInvoice/RecognizeQuotaInvoice1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQuotaInvoiceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeQuotaInvoiceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeQuotaInvoiceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeQuotaInvoiceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
