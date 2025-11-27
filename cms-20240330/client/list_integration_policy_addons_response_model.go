// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyAddonsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyAddonsResponseBody) *ListIntegrationPolicyAddonsResponse
	GetBody() *ListIntegrationPolicyAddonsResponseBody
}

type ListIntegrationPolicyAddonsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyAddonsResponse) GetBody() *ListIntegrationPolicyAddonsResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyAddonsResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyAddonsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponse) SetStatusCode(v int32) *ListIntegrationPolicyAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponse) SetBody(v *ListIntegrationPolicyAddonsResponseBody) *ListIntegrationPolicyAddonsResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
