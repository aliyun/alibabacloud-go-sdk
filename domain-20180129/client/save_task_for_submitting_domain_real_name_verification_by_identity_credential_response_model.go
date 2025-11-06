// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse
	GetBody() *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse struct {
	Headers    map[string]*string                                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) GetBody() *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody {
	return s.Body
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) SetBody(v *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
