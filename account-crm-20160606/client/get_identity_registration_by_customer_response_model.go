// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityRegistrationByCustomerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityRegistrationByCustomerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityRegistrationByCustomerResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityRegistrationByCustomerResponseBody) *GetIdentityRegistrationByCustomerResponse
	GetBody() *GetIdentityRegistrationByCustomerResponseBody
}

type GetIdentityRegistrationByCustomerResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityRegistrationByCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityRegistrationByCustomerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityRegistrationByCustomerResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityRegistrationByCustomerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityRegistrationByCustomerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityRegistrationByCustomerResponse) GetBody() *GetIdentityRegistrationByCustomerResponseBody {
	return s.Body
}

func (s *GetIdentityRegistrationByCustomerResponse) SetHeaders(v map[string]*string) *GetIdentityRegistrationByCustomerResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponse) SetStatusCode(v int32) *GetIdentityRegistrationByCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponse) SetBody(v *GetIdentityRegistrationByCustomerResponseBody) *GetIdentityRegistrationByCustomerResponse {
	s.Body = v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
