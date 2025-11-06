// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody
	GetTaskNo() *string
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
