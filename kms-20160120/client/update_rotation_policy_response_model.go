// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRotationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRotationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRotationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRotationPolicyResponseBody) *UpdateRotationPolicyResponse
	GetBody() *UpdateRotationPolicyResponseBody
}

type UpdateRotationPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRotationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRotationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRotationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRotationPolicyResponse) GetBody() *UpdateRotationPolicyResponseBody {
	return s.Body
}

func (s *UpdateRotationPolicyResponse) SetHeaders(v map[string]*string) *UpdateRotationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateRotationPolicyResponse) SetStatusCode(v int32) *UpdateRotationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRotationPolicyResponse) SetBody(v *UpdateRotationPolicyResponseBody) *UpdateRotationPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateRotationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
