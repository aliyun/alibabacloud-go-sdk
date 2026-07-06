// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStoragePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentStoragePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentStoragePolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentStoragePolicyResponseBody) *UpdateAgentStoragePolicyResponse
	GetBody() *UpdateAgentStoragePolicyResponseBody
}

type UpdateAgentStoragePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentStoragePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentStoragePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStoragePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentStoragePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentStoragePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentStoragePolicyResponse) GetBody() *UpdateAgentStoragePolicyResponseBody {
	return s.Body
}

func (s *UpdateAgentStoragePolicyResponse) SetHeaders(v map[string]*string) *UpdateAgentStoragePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentStoragePolicyResponse) SetStatusCode(v int32) *UpdateAgentStoragePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentStoragePolicyResponse) SetBody(v *UpdateAgentStoragePolicyResponseBody) *UpdateAgentStoragePolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentStoragePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
