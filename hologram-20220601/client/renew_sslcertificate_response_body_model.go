// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewSSLCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RenewSSLCertificateResponseBody
	GetData() *bool
	SetErrorCode(v string) *RenewSSLCertificateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RenewSSLCertificateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *RenewSSLCertificateResponseBody
	GetHttpStatusCode() *string
	SetSuccess(v bool) *RenewSSLCertificateResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *RenewSSLCertificateResponseBody
	GetRequestId() *string
}

type RenewSSLCertificateResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RenewSSLCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RenewSSLCertificateResponseBody) GetData() *bool {
	return s.Data
}

func (s *RenewSSLCertificateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RenewSSLCertificateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RenewSSLCertificateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RenewSSLCertificateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewSSLCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewSSLCertificateResponseBody) SetData(v bool) *RenewSSLCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *RenewSSLCertificateResponseBody) SetErrorCode(v string) *RenewSSLCertificateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RenewSSLCertificateResponseBody) SetErrorMessage(v string) *RenewSSLCertificateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RenewSSLCertificateResponseBody) SetHttpStatusCode(v string) *RenewSSLCertificateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewSSLCertificateResponseBody) SetSuccess(v bool) *RenewSSLCertificateResponseBody {
	s.Success = &v
	return s
}

func (s *RenewSSLCertificateResponseBody) SetRequestId(v string) *RenewSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewSSLCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
