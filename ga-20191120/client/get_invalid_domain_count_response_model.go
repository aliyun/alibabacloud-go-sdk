// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInvalidDomainCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInvalidDomainCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInvalidDomainCountResponse
	GetStatusCode() *int32
	SetBody(v *GetInvalidDomainCountResponseBody) *GetInvalidDomainCountResponse
	GetBody() *GetInvalidDomainCountResponseBody
}

type GetInvalidDomainCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInvalidDomainCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInvalidDomainCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInvalidDomainCountResponse) GoString() string {
	return s.String()
}

func (s *GetInvalidDomainCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInvalidDomainCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInvalidDomainCountResponse) GetBody() *GetInvalidDomainCountResponseBody {
	return s.Body
}

func (s *GetInvalidDomainCountResponse) SetHeaders(v map[string]*string) *GetInvalidDomainCountResponse {
	s.Headers = v
	return s
}

func (s *GetInvalidDomainCountResponse) SetStatusCode(v int32) *GetInvalidDomainCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInvalidDomainCountResponse) SetBody(v *GetInvalidDomainCountResponseBody) *GetInvalidDomainCountResponse {
	s.Body = v
	return s
}

func (s *GetInvalidDomainCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
