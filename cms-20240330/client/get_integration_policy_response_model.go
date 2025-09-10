// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntegrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntegrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetIntegrationPolicyResponseBody) *GetIntegrationPolicyResponse
	GetBody() *GetIntegrationPolicyResponseBody
}

type GetIntegrationPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntegrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntegrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntegrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntegrationPolicyResponse) GetBody() *GetIntegrationPolicyResponseBody {
	return s.Body
}

func (s *GetIntegrationPolicyResponse) SetHeaders(v map[string]*string) *GetIntegrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationPolicyResponse) SetStatusCode(v int32) *GetIntegrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationPolicyResponse) SetBody(v *GetIntegrationPolicyResponseBody) *GetIntegrationPolicyResponse {
	s.Body = v
	return s
}

func (s *GetIntegrationPolicyResponse) Validate() error {
	return dara.Validate(s)
}
