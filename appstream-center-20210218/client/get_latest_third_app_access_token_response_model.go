// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestThirdAppAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLatestThirdAppAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLatestThirdAppAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetLatestThirdAppAccessTokenResponseBody) *GetLatestThirdAppAccessTokenResponse
	GetBody() *GetLatestThirdAppAccessTokenResponseBody
}

type GetLatestThirdAppAccessTokenResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLatestThirdAppAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLatestThirdAppAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLatestThirdAppAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLatestThirdAppAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLatestThirdAppAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLatestThirdAppAccessTokenResponse) GetBody() *GetLatestThirdAppAccessTokenResponseBody {
	return s.Body
}

func (s *GetLatestThirdAppAccessTokenResponse) SetHeaders(v map[string]*string) *GetLatestThirdAppAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLatestThirdAppAccessTokenResponse) SetStatusCode(v int32) *GetLatestThirdAppAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenResponse) SetBody(v *GetLatestThirdAppAccessTokenResponseBody) *GetLatestThirdAppAccessTokenResponse {
	s.Body = v
	return s
}

func (s *GetLatestThirdAppAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
