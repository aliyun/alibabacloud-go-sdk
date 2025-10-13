// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIntegrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIntegrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIntegrationPolicyResponseBody) *DeleteIntegrationPolicyResponse
	GetBody() *DeleteIntegrationPolicyResponseBody
}

type DeleteIntegrationPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntegrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIntegrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIntegrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIntegrationPolicyResponse) GetBody() *DeleteIntegrationPolicyResponseBody {
	return s.Body
}

func (s *DeleteIntegrationPolicyResponse) SetHeaders(v map[string]*string) *DeleteIntegrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntegrationPolicyResponse) SetStatusCode(v int32) *DeleteIntegrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntegrationPolicyResponse) SetBody(v *DeleteIntegrationPolicyResponseBody) *DeleteIntegrationPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteIntegrationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
