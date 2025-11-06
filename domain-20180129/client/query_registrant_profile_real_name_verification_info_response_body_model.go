// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfileRealNameVerificationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityCredential(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetIdentityCredential() *string
	SetIdentityCredentialNo(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetIdentityCredentialNo() *string
	SetIdentityCredentialType(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetIdentityCredentialType() *string
	SetIdentityCredentialUrl(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetIdentityCredentialUrl() *string
	SetModificationDate(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetModificationDate() *string
	SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetRegistrantProfileId() *int64
	SetRequestId(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetRequestId() *string
	SetSubmissionDate(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody
	GetSubmissionDate() *string
}

type QueryRegistrantProfileRealNameVerificationInfoResponseBody struct {
	// example:
	//
	// dGVzdA==
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// example:
	//
	// 4111111111111110**
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// example:
	//
	// SFZ
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	// example:
	//
	// http://test.oss-cn-hangzhou.aliyuncs.com/20170522/1219541161213057_070445190.jpg
	IdentityCredentialUrl *string `json:"IdentityCredentialUrl,omitempty" xml:"IdentityCredentialUrl,omitempty"`
	// example:
	//
	// 2017-05-22 19:04:49
	ModificationDate *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	// example:
	//
	// 1234567
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// 4D73432C-7600-4779-ACBB-C3B5CA145D32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2017-05-22 19:04:49
	SubmissionDate *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetIdentityCredential() *string {
	return s.IdentityCredential
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetIdentityCredentialNo() *string {
	return s.IdentityCredentialNo
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetIdentityCredentialType() *string {
	return s.IdentityCredentialType
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetIdentityCredentialUrl() *string {
	return s.IdentityCredentialUrl
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetModificationDate() *string {
	return s.ModificationDate
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) GetSubmissionDate() *string {
	return s.SubmissionDate
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredential(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredential = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredentialNo(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredentialNo = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredentialType(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredentialType = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredentialUrl(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredentialUrl = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetModificationDate(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.ModificationDate = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetRequestId(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetSubmissionDate(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.SubmissionDate = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
