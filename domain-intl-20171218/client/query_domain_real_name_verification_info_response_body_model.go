// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealNameVerificationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetDomainName() *string
	SetIdentityCredential(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetIdentityCredential() *string
	SetIdentityCredentialNo(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetIdentityCredentialNo() *string
	SetIdentityCredentialType(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetIdentityCredentialType() *string
	SetIdentityCredentialUrl(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetIdentityCredentialUrl() *string
	SetInstanceId(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetRequestId() *string
	SetSubmissionDate(v string) *QueryDomainRealNameVerificationInfoResponseBody
	GetSubmissionDate() *string
}

type QueryDomainRealNameVerificationInfoResponseBody struct {
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IdentityCredential     *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo   *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	IdentityCredentialUrl  *string `json:"IdentityCredentialUrl,omitempty" xml:"IdentityCredentialUrl,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubmissionDate         *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
}

func (s QueryDomainRealNameVerificationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealNameVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetIdentityCredential() *string {
	return s.IdentityCredential
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetIdentityCredentialNo() *string {
	return s.IdentityCredentialNo
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetIdentityCredentialType() *string {
	return s.IdentityCredentialType
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetIdentityCredentialUrl() *string {
	return s.IdentityCredentialUrl
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) GetSubmissionDate() *string {
	return s.SubmissionDate
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetDomainName(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredential(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredential = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredentialNo(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredentialNo = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredentialType(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredentialType = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredentialUrl(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredentialUrl = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetInstanceId(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetRequestId(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetSubmissionDate(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.SubmissionDate = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
