// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSSLDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSSLDetailsResponseBodyData) *ListSSLDetailsResponseBody
	GetData() *ListSSLDetailsResponseBodyData
	SetErrCode(v string) *ListSSLDetailsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListSSLDetailsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListSSLDetailsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListSSLDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSSLDetailsResponseBody
	GetSuccess() *bool
}

type ListSSLDetailsResponseBody struct {
	Data *ListSSLDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSSLDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSSLDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSSLDetailsResponseBody) GetData() *ListSSLDetailsResponseBodyData {
	return s.Data
}

func (s *ListSSLDetailsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListSSLDetailsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListSSLDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSSLDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSSLDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSSLDetailsResponseBody) SetData(v *ListSSLDetailsResponseBodyData) *ListSSLDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListSSLDetailsResponseBody) SetErrCode(v string) *ListSSLDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListSSLDetailsResponseBody) SetErrMessage(v string) *ListSSLDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListSSLDetailsResponseBody) SetHttpStatusCode(v int32) *ListSSLDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSSLDetailsResponseBody) SetRequestId(v string) *ListSSLDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSSLDetailsResponseBody) SetSuccess(v bool) *ListSSLDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSSLDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSSLDetailsResponseBodyData struct {
	// example:
	//
	// starrocks-ssl
	Aliases *string `json:"Aliases,omitempty" xml:"Aliases,omitempty"`
	// example:
	//
	// true
	CustomCert *bool `json:"CustomCert,omitempty" xml:"CustomCert,omitempty"`
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// example:
	//
	// true
	IsValid *bool `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	// example:
	//
	// CN=fe-sr-cn-xxx123-internal.starrocks.aliyuncs.com, OU=Aliyun Big Data Platform EMR StarRocks, O=Aliyun, L=BJ, ST=CN, C=CN
	IssuerDN *string `json:"IssuerDN,omitempty" xml:"IssuerDN,omitempty"`
	// example:
	//
	// 1751010702000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// 1751000702000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// example:
	//
	// MIIJkAIBAzCCCVYGCSqGSIb3DQEHAaCCCUcEgglDMIIJPzCCA+8GCSqGSIb3DQEHBqCCA+AwggPc
	//
	// AgEAMIID1QYJKoZIhvcNAQcBMBwGCiqGSIb3DQEMAQYwDgQI3e4V2eXgGFMCAggAgIIDqJ/L8sA
	//
	// ......
	//
	// AB3eKpMa7rQc==
	SslCertificateText *string `json:"SslCertificateText,omitempty" xml:"SslCertificateText,omitempty"`
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// CN=fe-{clusterId}-internal.starrocks.aliyuncs.com, OU=Aliyun Big Data Platform EMR StarRocks, O=Aliyun, L=BJ, ST=CN, C=CN
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
}

func (s ListSSLDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSSLDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSSLDetailsResponseBodyData) GetAliases() *string {
	return s.Aliases
}

func (s *ListSSLDetailsResponseBodyData) GetCustomCert() *bool {
	return s.CustomCert
}

func (s *ListSSLDetailsResponseBodyData) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *ListSSLDetailsResponseBodyData) GetIsValid() *bool {
	return s.IsValid
}

func (s *ListSSLDetailsResponseBodyData) GetIssuerDN() *string {
	return s.IssuerDN
}

func (s *ListSSLDetailsResponseBodyData) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *ListSSLDetailsResponseBodyData) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *ListSSLDetailsResponseBodyData) GetSslCertificateText() *string {
	return s.SslCertificateText
}

func (s *ListSSLDetailsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListSSLDetailsResponseBodyData) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *ListSSLDetailsResponseBodyData) SetAliases(v string) *ListSSLDetailsResponseBodyData {
	s.Aliases = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetCustomCert(v bool) *ListSSLDetailsResponseBodyData {
	s.CustomCert = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetEnableSSL(v bool) *ListSSLDetailsResponseBodyData {
	s.EnableSSL = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetIsValid(v bool) *ListSSLDetailsResponseBodyData {
	s.IsValid = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetIssuerDN(v string) *ListSSLDetailsResponseBodyData {
	s.IssuerDN = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetNotAfter(v int64) *ListSSLDetailsResponseBodyData {
	s.NotAfter = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetNotBefore(v int64) *ListSSLDetailsResponseBodyData {
	s.NotBefore = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetSslCertificateText(v string) *ListSSLDetailsResponseBodyData {
	s.SslCertificateText = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetStatus(v string) *ListSSLDetailsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) SetSubjectDN(v string) *ListSSLDetailsResponseBodyData {
	s.SubjectDN = &v
	return s
}

func (s *ListSSLDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
