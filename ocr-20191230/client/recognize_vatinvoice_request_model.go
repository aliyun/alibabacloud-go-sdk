// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVATInvoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *RecognizeVATInvoiceRequest
	GetFileType() *string
	SetFileURL(v string) *RecognizeVATInvoiceRequest
	GetFileURL() *string
}

type RecognizeVATInvoiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// jpg
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeVATInvoice/RecognizeVATInvoice3.jpg
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizeVATInvoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceRequest) GetFileType() *string {
	return s.FileType
}

func (s *RecognizeVATInvoiceRequest) GetFileURL() *string {
	return s.FileURL
}

func (s *RecognizeVATInvoiceRequest) SetFileType(v string) *RecognizeVATInvoiceRequest {
	s.FileType = &v
	return s
}

func (s *RecognizeVATInvoiceRequest) SetFileURL(v string) *RecognizeVATInvoiceRequest {
	s.FileURL = &v
	return s
}

func (s *RecognizeVATInvoiceRequest) Validate() error {
	return dara.Validate(s)
}
