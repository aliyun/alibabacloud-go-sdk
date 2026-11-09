// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomAuthIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomAuthIdentityResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomAuthIdentityResponseBody) *AddCustomAuthIdentityResponse
	GetBody() *AddCustomAuthIdentityResponseBody
}

type AddCustomAuthIdentityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomAuthIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomAuthIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthIdentityResponse) GoString() string {
	return s.String()
}

func (s *AddCustomAuthIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomAuthIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomAuthIdentityResponse) GetBody() *AddCustomAuthIdentityResponseBody {
	return s.Body
}

func (s *AddCustomAuthIdentityResponse) SetHeaders(v map[string]*string) *AddCustomAuthIdentityResponse {
	s.Headers = v
	return s
}

func (s *AddCustomAuthIdentityResponse) SetStatusCode(v int32) *AddCustomAuthIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomAuthIdentityResponse) SetBody(v *AddCustomAuthIdentityResponseBody) *AddCustomAuthIdentityResponse {
	s.Body = v
	return s
}

func (s *AddCustomAuthIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
