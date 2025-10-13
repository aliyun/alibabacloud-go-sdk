// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRootCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetRootCertificateResponseBody
	GetData() *string
	SetErrorCode(v string) *GetRootCertificateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetRootCertificateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetRootCertificateResponseBody
	GetHttpStatusCode() *string
	SetSuccess(v bool) *GetRootCertificateResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *GetRootCertificateResponseBody
	GetRequestId() *string
}

type GetRootCertificateResponseBody struct {
	// example:
	//
	// { "rootCertificate": "BEGIN xxxx"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s GetRootCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRootCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetRootCertificateResponseBody) GetData() *string {
	return s.Data
}

func (s *GetRootCertificateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRootCertificateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRootCertificateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetRootCertificateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRootCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRootCertificateResponseBody) SetData(v string) *GetRootCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *GetRootCertificateResponseBody) SetErrorCode(v string) *GetRootCertificateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRootCertificateResponseBody) SetErrorMessage(v string) *GetRootCertificateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRootCertificateResponseBody) SetHttpStatusCode(v string) *GetRootCertificateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRootCertificateResponseBody) SetSuccess(v bool) *GetRootCertificateResponseBody {
	s.Success = &v
	return s
}

func (s *GetRootCertificateResponseBody) SetRequestId(v string) *GetRootCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRootCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
