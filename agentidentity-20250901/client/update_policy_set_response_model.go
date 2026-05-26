// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicySetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolicySetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolicySetResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolicySetResponseBody) *UpdatePolicySetResponse
	GetBody() *UpdatePolicySetResponseBody
}

type UpdatePolicySetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicySetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicySetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicySetResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicySetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolicySetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolicySetResponse) GetBody() *UpdatePolicySetResponseBody {
	return s.Body
}

func (s *UpdatePolicySetResponse) SetHeaders(v map[string]*string) *UpdatePolicySetResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicySetResponse) SetStatusCode(v int32) *UpdatePolicySetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicySetResponse) SetBody(v *UpdatePolicySetResponseBody) *UpdatePolicySetResponse {
	s.Body = v
	return s
}

func (s *UpdatePolicySetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
