// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeVATInvoiceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *RecognizeVATInvoiceAdvanceRequest
	GetFileType() *string
	SetFileURLObject(v io.Reader) *RecognizeVATInvoiceAdvanceRequest
	GetFileURLObject() io.Reader
}

type RecognizeVATInvoiceAdvanceRequest struct {
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
	FileURLObject io.Reader `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizeVATInvoiceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceAdvanceRequest) GetFileType() *string {
	return s.FileType
}

func (s *RecognizeVATInvoiceAdvanceRequest) GetFileURLObject() io.Reader {
	return s.FileURLObject
}

func (s *RecognizeVATInvoiceAdvanceRequest) SetFileType(v string) *RecognizeVATInvoiceAdvanceRequest {
	s.FileType = &v
	return s
}

func (s *RecognizeVATInvoiceAdvanceRequest) SetFileURLObject(v io.Reader) *RecognizeVATInvoiceAdvanceRequest {
	s.FileURLObject = v
	return s
}

func (s *RecognizeVATInvoiceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
