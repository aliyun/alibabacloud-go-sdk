// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlsInspectCertificateDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertId(v string) *GetTlsInspectCertificateDownloadUrlResponseBody
	GetCaCertId() *string
	SetDownloadUrl(v string) *GetTlsInspectCertificateDownloadUrlResponseBody
	GetDownloadUrl() *string
	SetRequestId(v string) *GetTlsInspectCertificateDownloadUrlResponseBody
	GetRequestId() *string
}

type GetTlsInspectCertificateDownloadUrlResponseBody struct {
	CaCertId *string `json:"CaCertId,omitempty" xml:"CaCertId,omitempty"`
	// example:
	//
	// https://cfw-tls-inspect-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/tls_cert%2F2025-08-13%2F1850****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-******h4j6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTlsInspectCertificateDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTlsInspectCertificateDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) GetCaCertId() *string {
	return s.CaCertId
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) SetCaCertId(v string) *GetTlsInspectCertificateDownloadUrlResponseBody {
	s.CaCertId = &v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) SetDownloadUrl(v string) *GetTlsInspectCertificateDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) SetRequestId(v string) *GetTlsInspectCertificateDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
