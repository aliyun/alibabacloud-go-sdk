// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAppDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAppDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAppDomainResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAppDomainResponseBody) *UnbindAppDomainResponse
	GetBody() *UnbindAppDomainResponseBody
}

type UnbindAppDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAppDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAppDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppDomainResponse) GoString() string {
	return s.String()
}

func (s *UnbindAppDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAppDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAppDomainResponse) GetBody() *UnbindAppDomainResponseBody {
	return s.Body
}

func (s *UnbindAppDomainResponse) SetHeaders(v map[string]*string) *UnbindAppDomainResponse {
	s.Headers = v
	return s
}

func (s *UnbindAppDomainResponse) SetStatusCode(v int32) *UnbindAppDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAppDomainResponse) SetBody(v *UnbindAppDomainResponseBody) *UnbindAppDomainResponse {
	s.Body = v
	return s
}

func (s *UnbindAppDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
