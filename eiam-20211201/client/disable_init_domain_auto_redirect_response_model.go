// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInitDomainAutoRedirectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInitDomainAutoRedirectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInitDomainAutoRedirectResponse
	GetStatusCode() *int32
	SetBody(v *DisableInitDomainAutoRedirectResponseBody) *DisableInitDomainAutoRedirectResponse
	GetBody() *DisableInitDomainAutoRedirectResponseBody
}

type DisableInitDomainAutoRedirectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInitDomainAutoRedirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInitDomainAutoRedirectResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInitDomainAutoRedirectResponse) GoString() string {
	return s.String()
}

func (s *DisableInitDomainAutoRedirectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInitDomainAutoRedirectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInitDomainAutoRedirectResponse) GetBody() *DisableInitDomainAutoRedirectResponseBody {
	return s.Body
}

func (s *DisableInitDomainAutoRedirectResponse) SetHeaders(v map[string]*string) *DisableInitDomainAutoRedirectResponse {
	s.Headers = v
	return s
}

func (s *DisableInitDomainAutoRedirectResponse) SetStatusCode(v int32) *DisableInitDomainAutoRedirectResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInitDomainAutoRedirectResponse) SetBody(v *DisableInitDomainAutoRedirectResponseBody) *DisableInitDomainAutoRedirectResponse {
	s.Body = v
	return s
}

func (s *DisableInitDomainAutoRedirectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
