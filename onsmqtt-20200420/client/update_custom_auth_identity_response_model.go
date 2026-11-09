// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAuthIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomAuthIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomAuthIdentityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomAuthIdentityResponseBody) *UpdateCustomAuthIdentityResponse
	GetBody() *UpdateCustomAuthIdentityResponseBody
}

type UpdateCustomAuthIdentityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomAuthIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomAuthIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAuthIdentityResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomAuthIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomAuthIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomAuthIdentityResponse) GetBody() *UpdateCustomAuthIdentityResponseBody {
	return s.Body
}

func (s *UpdateCustomAuthIdentityResponse) SetHeaders(v map[string]*string) *UpdateCustomAuthIdentityResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomAuthIdentityResponse) SetStatusCode(v int32) *UpdateCustomAuthIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomAuthIdentityResponse) SetBody(v *UpdateCustomAuthIdentityResponseBody) *UpdateCustomAuthIdentityResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomAuthIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
