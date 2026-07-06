// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAgentStoragePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAgentStoragePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAgentStoragePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CheckAgentStoragePolicyResponseBody) *CheckAgentStoragePolicyResponse
	GetBody() *CheckAgentStoragePolicyResponseBody
}

type CheckAgentStoragePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAgentStoragePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAgentStoragePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentStoragePolicyResponse) GoString() string {
	return s.String()
}

func (s *CheckAgentStoragePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAgentStoragePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAgentStoragePolicyResponse) GetBody() *CheckAgentStoragePolicyResponseBody {
	return s.Body
}

func (s *CheckAgentStoragePolicyResponse) SetHeaders(v map[string]*string) *CheckAgentStoragePolicyResponse {
	s.Headers = v
	return s
}

func (s *CheckAgentStoragePolicyResponse) SetStatusCode(v int32) *CheckAgentStoragePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAgentStoragePolicyResponse) SetBody(v *CheckAgentStoragePolicyResponseBody) *CheckAgentStoragePolicyResponse {
	s.Body = v
	return s
}

func (s *CheckAgentStoragePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
