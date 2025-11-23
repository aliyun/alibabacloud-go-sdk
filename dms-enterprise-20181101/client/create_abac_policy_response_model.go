// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAbacPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAbacPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAbacPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAbacPolicyResponseBody) *CreateAbacPolicyResponse
	GetBody() *CreateAbacPolicyResponseBody
}

type CreateAbacPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAbacPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAbacPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAbacPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAbacPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAbacPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAbacPolicyResponse) GetBody() *CreateAbacPolicyResponseBody {
	return s.Body
}

func (s *CreateAbacPolicyResponse) SetHeaders(v map[string]*string) *CreateAbacPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAbacPolicyResponse) SetStatusCode(v int32) *CreateAbacPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAbacPolicyResponse) SetBody(v *CreateAbacPolicyResponseBody) *CreateAbacPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAbacPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
