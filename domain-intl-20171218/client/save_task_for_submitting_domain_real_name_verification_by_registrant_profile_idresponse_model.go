// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse
	GetBody() *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse struct {
	Headers    map[string]*string                                                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) GetBody() *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody {
	return s.Body
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) SetBody(v *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
