// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainDnsChallengeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDomainDnsChallengeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDomainDnsChallengeResponse
	GetStatusCode() *int32
	SetBody(v *GetDomainDnsChallengeResponseBody) *GetDomainDnsChallengeResponse
	GetBody() *GetDomainDnsChallengeResponseBody
}

type GetDomainDnsChallengeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainDnsChallengeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainDnsChallengeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDomainDnsChallengeResponse) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDomainDnsChallengeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDomainDnsChallengeResponse) GetBody() *GetDomainDnsChallengeResponseBody {
	return s.Body
}

func (s *GetDomainDnsChallengeResponse) SetHeaders(v map[string]*string) *GetDomainDnsChallengeResponse {
	s.Headers = v
	return s
}

func (s *GetDomainDnsChallengeResponse) SetStatusCode(v int32) *GetDomainDnsChallengeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainDnsChallengeResponse) SetBody(v *GetDomainDnsChallengeResponseBody) *GetDomainDnsChallengeResponse {
	s.Body = v
	return s
}

func (s *GetDomainDnsChallengeResponse) Validate() error {
	return dara.Validate(s)
}
