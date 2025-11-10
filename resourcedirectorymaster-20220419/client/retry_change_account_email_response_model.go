// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChangeAccountEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryChangeAccountEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryChangeAccountEmailResponse
	GetStatusCode() *int32
	SetBody(v *RetryChangeAccountEmailResponseBody) *RetryChangeAccountEmailResponse
	GetBody() *RetryChangeAccountEmailResponseBody
}

type RetryChangeAccountEmailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryChangeAccountEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryChangeAccountEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryChangeAccountEmailResponse) GetBody() *RetryChangeAccountEmailResponseBody {
	return s.Body
}

func (s *RetryChangeAccountEmailResponse) SetHeaders(v map[string]*string) *RetryChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *RetryChangeAccountEmailResponse) SetStatusCode(v int32) *RetryChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryChangeAccountEmailResponse) SetBody(v *RetryChangeAccountEmailResponseBody) *RetryChangeAccountEmailResponse {
	s.Body = v
	return s
}

func (s *RetryChangeAccountEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
