// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLiveDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLiveDomainResponse
	GetStatusCode() *int32
	SetBody(v *StartLiveDomainResponseBody) *StartLiveDomainResponse
	GetBody() *StartLiveDomainResponseBody
}

type StartLiveDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLiveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLiveDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *StartLiveDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLiveDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLiveDomainResponse) GetBody() *StartLiveDomainResponseBody {
	return s.Body
}

func (s *StartLiveDomainResponse) SetHeaders(v map[string]*string) *StartLiveDomainResponse {
	s.Headers = v
	return s
}

func (s *StartLiveDomainResponse) SetStatusCode(v int32) *StartLiveDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLiveDomainResponse) SetBody(v *StartLiveDomainResponseBody) *StartLiveDomainResponse {
	s.Body = v
	return s
}

func (s *StartLiveDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
