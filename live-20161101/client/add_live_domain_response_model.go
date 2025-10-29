// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveDomainResponseBody) *AddLiveDomainResponse
	GetBody() *AddLiveDomainResponseBody
}

type AddLiveDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveDomainResponse) GetBody() *AddLiveDomainResponseBody {
	return s.Body
}

func (s *AddLiveDomainResponse) SetHeaders(v map[string]*string) *AddLiveDomainResponse {
	s.Headers = v
	return s
}

func (s *AddLiveDomainResponse) SetStatusCode(v int32) *AddLiveDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveDomainResponse) SetBody(v *AddLiveDomainResponseBody) *AddLiveDomainResponse {
	s.Body = v
	return s
}

func (s *AddLiveDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
