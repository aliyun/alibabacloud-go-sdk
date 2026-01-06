// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetUserTokenResponseBody) *GetUserTokenResponse
	GetBody() *GetUserTokenResponseBody
}

type GetUserTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserTokenResponse) GetBody() *GetUserTokenResponseBody {
	return s.Body
}

func (s *GetUserTokenResponse) SetHeaders(v map[string]*string) *GetUserTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUserTokenResponse) SetStatusCode(v int32) *GetUserTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserTokenResponse) SetBody(v *GetUserTokenResponseBody) *GetUserTokenResponse {
	s.Body = v
	return s
}

func (s *GetUserTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
