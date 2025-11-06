// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupDomainAutoRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetupDomainAutoRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetupDomainAutoRenewResponse
	GetStatusCode() *int32
	SetBody(v *SetupDomainAutoRenewResponseBody) *SetupDomainAutoRenewResponse
	GetBody() *SetupDomainAutoRenewResponseBody
}

type SetupDomainAutoRenewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupDomainAutoRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupDomainAutoRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s SetupDomainAutoRenewResponse) GoString() string {
	return s.String()
}

func (s *SetupDomainAutoRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetupDomainAutoRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetupDomainAutoRenewResponse) GetBody() *SetupDomainAutoRenewResponseBody {
	return s.Body
}

func (s *SetupDomainAutoRenewResponse) SetHeaders(v map[string]*string) *SetupDomainAutoRenewResponse {
	s.Headers = v
	return s
}

func (s *SetupDomainAutoRenewResponse) SetStatusCode(v int32) *SetupDomainAutoRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupDomainAutoRenewResponse) SetBody(v *SetupDomainAutoRenewResponseBody) *SetupDomainAutoRenewResponse {
	s.Body = v
	return s
}

func (s *SetupDomainAutoRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
