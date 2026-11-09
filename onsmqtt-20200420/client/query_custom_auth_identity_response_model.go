// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomAuthIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomAuthIdentityResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomAuthIdentityResponseBody) *QueryCustomAuthIdentityResponse
	GetBody() *QueryCustomAuthIdentityResponseBody
}

type QueryCustomAuthIdentityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomAuthIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomAuthIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthIdentityResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomAuthIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomAuthIdentityResponse) GetBody() *QueryCustomAuthIdentityResponseBody {
	return s.Body
}

func (s *QueryCustomAuthIdentityResponse) SetHeaders(v map[string]*string) *QueryCustomAuthIdentityResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomAuthIdentityResponse) SetStatusCode(v int32) *QueryCustomAuthIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomAuthIdentityResponse) SetBody(v *QueryCustomAuthIdentityResponseBody) *QueryCustomAuthIdentityResponse {
	s.Body = v
	return s
}

func (s *QueryCustomAuthIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
