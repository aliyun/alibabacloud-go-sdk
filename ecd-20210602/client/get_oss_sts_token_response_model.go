// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssStsTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssStsTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssStsTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetOssStsTokenResponseBody) *GetOssStsTokenResponse
	GetBody() *GetOssStsTokenResponseBody
}

type GetOssStsTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssStsTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssStsTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssStsTokenResponse) GoString() string {
	return s.String()
}

func (s *GetOssStsTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssStsTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssStsTokenResponse) GetBody() *GetOssStsTokenResponseBody {
	return s.Body
}

func (s *GetOssStsTokenResponse) SetHeaders(v map[string]*string) *GetOssStsTokenResponse {
	s.Headers = v
	return s
}

func (s *GetOssStsTokenResponse) SetStatusCode(v int32) *GetOssStsTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssStsTokenResponse) SetBody(v *GetOssStsTokenResponseBody) *GetOssStsTokenResponse {
	s.Body = v
	return s
}

func (s *GetOssStsTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
