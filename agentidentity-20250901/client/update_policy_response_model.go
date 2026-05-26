// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolicyResponseBody) *UpdatePolicyResponse
	GetBody() *UpdatePolicyResponseBody
}

type UpdatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolicyResponse) GetBody() *UpdatePolicyResponseBody {
	return s.Body
}

func (s *UpdatePolicyResponse) SetHeaders(v map[string]*string) *UpdatePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicyResponse) SetStatusCode(v int32) *UpdatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyResponse) SetBody(v *UpdatePolicyResponseBody) *UpdatePolicyResponse {
	s.Body = v
	return s
}

func (s *UpdatePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
