// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQualificationDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int32) *QueryQualificationDetailResponseBody
	GetAuditStatus() *int32
	SetCredentials(v *QueryQualificationDetailResponseBodyCredentials) *QueryQualificationDetailResponseBody
	GetCredentials() *QueryQualificationDetailResponseBodyCredentials
	SetRequestId(v string) *QueryQualificationDetailResponseBody
	GetRequestId() *string
	SetTrackId(v string) *QueryQualificationDetailResponseBody
	GetTrackId() *string
}

type QueryQualificationDetailResponseBody struct {
	// example:
	//
	// 1
	AuditStatus *int32                                           `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	Credentials *QueryQualificationDetailResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 943a1662898a****0acbdbeca91
	TrackId *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s QueryQualificationDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryQualificationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponseBody) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *QueryQualificationDetailResponseBody) GetCredentials() *QueryQualificationDetailResponseBodyCredentials {
	return s.Credentials
}

func (s *QueryQualificationDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryQualificationDetailResponseBody) GetTrackId() *string {
	return s.TrackId
}

func (s *QueryQualificationDetailResponseBody) SetAuditStatus(v int32) *QueryQualificationDetailResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *QueryQualificationDetailResponseBody) SetCredentials(v *QueryQualificationDetailResponseBodyCredentials) *QueryQualificationDetailResponseBody {
	s.Credentials = v
	return s
}

func (s *QueryQualificationDetailResponseBody) SetRequestId(v string) *QueryQualificationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQualificationDetailResponseBody) SetTrackId(v string) *QueryQualificationDetailResponseBody {
	s.TrackId = &v
	return s
}

func (s *QueryQualificationDetailResponseBody) Validate() error {
	if s.Credentials != nil {
		if err := s.Credentials.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryQualificationDetailResponseBodyCredentials struct {
	QualificationCredential []*QueryQualificationDetailResponseBodyCredentialsQualificationCredential `json:"QualificationCredential,omitempty" xml:"QualificationCredential,omitempty" type:"Repeated"`
}

func (s QueryQualificationDetailResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s QueryQualificationDetailResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponseBodyCredentials) GetQualificationCredential() []*QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	return s.QualificationCredential
}

func (s *QueryQualificationDetailResponseBodyCredentials) SetQualificationCredential(v []*QueryQualificationDetailResponseBodyCredentialsQualificationCredential) *QueryQualificationDetailResponseBodyCredentials {
	s.QualificationCredential = v
	return s
}

func (s *QueryQualificationDetailResponseBodyCredentials) Validate() error {
	if s.QualificationCredential != nil {
		for _, item := range s.QualificationCredential {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryQualificationDetailResponseBodyCredentialsQualificationCredential struct {
	// example:
	//
	// 92610725MA7G2E****
	CredentialNo *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	// example:
	//
	// SHSQB
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	CredentialUrl  *string `json:"CredentialUrl,omitempty" xml:"CredentialUrl,omitempty"`
}

func (s QueryQualificationDetailResponseBodyCredentialsQualificationCredential) String() string {
	return dara.Prettify(s)
}

func (s QueryQualificationDetailResponseBodyCredentialsQualificationCredential) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) GetCredentialType() *string {
	return s.CredentialType
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) GetCredentialUrl() *string {
	return s.CredentialUrl
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) SetCredentialNo(v string) *QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	s.CredentialNo = &v
	return s
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) SetCredentialType(v string) *QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	s.CredentialType = &v
	return s
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) SetCredentialUrl(v string) *QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	s.CredentialUrl = &v
	return s
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) Validate() error {
	return dara.Validate(s)
}
