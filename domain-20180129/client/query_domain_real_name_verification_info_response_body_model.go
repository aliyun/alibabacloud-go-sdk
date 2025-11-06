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
	// example:
	//
	// aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// dGVzdA==
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// example:
	//
	// 5****************9
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// example:
	//
	// SFZ
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	// example:
	//
	// http://dbu-nap-p.oss-cn-hangzhou.aliyuncs.com/20190219/140692647406xxxx_5d6baea3e7314fd986afdd86e33exxxx.jpg
	IdentityCredentialUrl *string `json:"IdentityCredentialUrl,omitempty" xml:"IdentityCredentialUrl,omitempty"`
	// example:
	//
	// S2019270W570****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 4DF9D693-0D5B-4EB7-8922-7ECA6BD59314
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2018-03-28 00:41:42
	SubmissionDate *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
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
