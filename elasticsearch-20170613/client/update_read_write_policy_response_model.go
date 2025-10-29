// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReadWritePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateReadWritePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateReadWritePolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateReadWritePolicyResponseBody) *UpdateReadWritePolicyResponse
	GetBody() *UpdateReadWritePolicyResponseBody
}

type UpdateReadWritePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateReadWritePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateReadWritePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadWritePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateReadWritePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateReadWritePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateReadWritePolicyResponse) GetBody() *UpdateReadWritePolicyResponseBody {
	return s.Body
}

func (s *UpdateReadWritePolicyResponse) SetHeaders(v map[string]*string) *UpdateReadWritePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateReadWritePolicyResponse) SetStatusCode(v int32) *UpdateReadWritePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReadWritePolicyResponse) SetBody(v *UpdateReadWritePolicyResponseBody) *UpdateReadWritePolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateReadWritePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
