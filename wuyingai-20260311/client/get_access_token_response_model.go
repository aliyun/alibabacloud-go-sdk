// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessTokenResponseBody) *GetAccessTokenResponse
	GetBody() *GetAccessTokenResponseBody
}

type GetAccessTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessTokenResponse) GetBody() *GetAccessTokenResponseBody {
	return s.Body
}

func (s *GetAccessTokenResponse) SetHeaders(v map[string]*string) *GetAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAccessTokenResponse) SetStatusCode(v int32) *GetAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessTokenResponse) SetBody(v *GetAccessTokenResponseBody) *GetAccessTokenResponse {
	s.Body = v
	return s
}

func (s *GetAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
