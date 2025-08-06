// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistDRMCertInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *RegistDRMCertInfoResponseBody
	GetCertId() *string
	SetRequestId(v string) *RegistDRMCertInfoResponseBody
	GetRequestId() *string
}

type RegistDRMCertInfoResponseBody struct {
	CertId    *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegistDRMCertInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegistDRMCertInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RegistDRMCertInfoResponseBody) GetCertId() *string {
	return s.CertId
}

func (s *RegistDRMCertInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegistDRMCertInfoResponseBody) SetCertId(v string) *RegistDRMCertInfoResponseBody {
	s.CertId = &v
	return s
}

func (s *RegistDRMCertInfoResponseBody) SetRequestId(v string) *RegistDRMCertInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegistDRMCertInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
