// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse
	GetBody() *CreatePolicyResponseBody
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicyResponse) GetBody() *CreatePolicyResponseBody {
	return s.Body
}

func (s *CreatePolicyResponse) SetHeaders(v map[string]*string) *CreatePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyResponse) SetStatusCode(v int32) *CreatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

func (s *CreatePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
