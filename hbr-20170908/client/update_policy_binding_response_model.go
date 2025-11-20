// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolicyBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolicyBindingResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolicyBindingResponseBody) *UpdatePolicyBindingResponse
	GetBody() *UpdatePolicyBindingResponseBody
}

type UpdatePolicyBindingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolicyBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolicyBindingResponse) GetBody() *UpdatePolicyBindingResponseBody {
	return s.Body
}

func (s *UpdatePolicyBindingResponse) SetHeaders(v map[string]*string) *UpdatePolicyBindingResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicyBindingResponse) SetStatusCode(v int32) *UpdatePolicyBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyBindingResponse) SetBody(v *UpdatePolicyBindingResponseBody) *UpdatePolicyBindingResponse {
	s.Body = v
	return s
}

func (s *UpdatePolicyBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
