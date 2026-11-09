// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomAuthIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomAuthIdentityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomAuthIdentityResponseBody) *DeleteCustomAuthIdentityResponse
	GetBody() *DeleteCustomAuthIdentityResponseBody
}

type DeleteCustomAuthIdentityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomAuthIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomAuthIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthIdentityResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomAuthIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomAuthIdentityResponse) GetBody() *DeleteCustomAuthIdentityResponseBody {
	return s.Body
}

func (s *DeleteCustomAuthIdentityResponse) SetHeaders(v map[string]*string) *DeleteCustomAuthIdentityResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomAuthIdentityResponse) SetStatusCode(v int32) *DeleteCustomAuthIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomAuthIdentityResponse) SetBody(v *DeleteCustomAuthIdentityResponseBody) *DeleteCustomAuthIdentityResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomAuthIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
