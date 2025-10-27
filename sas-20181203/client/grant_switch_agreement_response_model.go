// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSwitchAgreementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantSwitchAgreementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantSwitchAgreementResponse
	GetStatusCode() *int32
	SetBody(v *GrantSwitchAgreementResponseBody) *GrantSwitchAgreementResponse
	GetBody() *GrantSwitchAgreementResponseBody
}

type GrantSwitchAgreementResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantSwitchAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantSwitchAgreementResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantSwitchAgreementResponse) GoString() string {
	return s.String()
}

func (s *GrantSwitchAgreementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantSwitchAgreementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantSwitchAgreementResponse) GetBody() *GrantSwitchAgreementResponseBody {
	return s.Body
}

func (s *GrantSwitchAgreementResponse) SetHeaders(v map[string]*string) *GrantSwitchAgreementResponse {
	s.Headers = v
	return s
}

func (s *GrantSwitchAgreementResponse) SetStatusCode(v int32) *GrantSwitchAgreementResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantSwitchAgreementResponse) SetBody(v *GrantSwitchAgreementResponseBody) *GrantSwitchAgreementResponse {
	s.Body = v
	return s
}

func (s *GrantSwitchAgreementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
