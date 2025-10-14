// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateAttributeDto(v *GetCertificateAttributeResponseBodyCertificateAttributeDto) *GetCertificateAttributeResponseBody
	GetCertificateAttributeDto() *GetCertificateAttributeResponseBodyCertificateAttributeDto
	SetErrorCode(v string) *GetCertificateAttributeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetCertificateAttributeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetCertificateAttributeResponseBody
	GetHttpStatusCode() *string
	SetSuccess(v bool) *GetCertificateAttributeResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *GetCertificateAttributeResponseBody
	GetRequestId() *string
}

type GetCertificateAttributeResponseBody struct {
	CertificateAttributeDto *GetCertificateAttributeResponseBodyCertificateAttributeDto `json:"CertificateAttributeDto,omitempty" xml:"CertificateAttributeDto,omitempty" type:"Struct"`
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

func (s GetCertificateAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateAttributeResponseBody) GetCertificateAttributeDto() *GetCertificateAttributeResponseBodyCertificateAttributeDto {
	return s.CertificateAttributeDto
}

func (s *GetCertificateAttributeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCertificateAttributeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCertificateAttributeResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetCertificateAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCertificateAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateAttributeResponseBody) SetCertificateAttributeDto(v *GetCertificateAttributeResponseBodyCertificateAttributeDto) *GetCertificateAttributeResponseBody {
	s.CertificateAttributeDto = v
	return s
}

func (s *GetCertificateAttributeResponseBody) SetErrorCode(v string) *GetCertificateAttributeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCertificateAttributeResponseBody) SetErrorMessage(v string) *GetCertificateAttributeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCertificateAttributeResponseBody) SetHttpStatusCode(v string) *GetCertificateAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCertificateAttributeResponseBody) SetSuccess(v bool) *GetCertificateAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *GetCertificateAttributeResponseBody) SetRequestId(v string) *GetCertificateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateAttributeResponseBody) Validate() error {
	if s.CertificateAttributeDto != nil {
		if err := s.CertificateAttributeDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCertificateAttributeResponseBodyCertificateAttributeDto struct {
	// example:
	//
	// true
	EnableSSL *bool `json:"enableSSL,omitempty" xml:"enableSSL,omitempty"`
	// example:
	//
	// 1790583135000
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// example:
	//
	// effective
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetCertificateAttributeResponseBodyCertificateAttributeDto) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateAttributeResponseBodyCertificateAttributeDto) GoString() string {
	return s.String()
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) GetStatus() *string {
	return s.Status
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) SetEnableSSL(v bool) *GetCertificateAttributeResponseBodyCertificateAttributeDto {
	s.EnableSSL = &v
	return s
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) SetExpirationTime(v int64) *GetCertificateAttributeResponseBodyCertificateAttributeDto {
	s.ExpirationTime = &v
	return s
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) SetStatus(v string) *GetCertificateAttributeResponseBodyCertificateAttributeDto {
	s.Status = &v
	return s
}

func (s *GetCertificateAttributeResponseBodyCertificateAttributeDto) Validate() error {
	return dara.Validate(s)
}
