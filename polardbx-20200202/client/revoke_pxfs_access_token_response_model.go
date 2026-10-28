// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePxfsAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokePxfsAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokePxfsAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *RevokePxfsAccessTokenResponseBody) *RevokePxfsAccessTokenResponse
	GetBody() *RevokePxfsAccessTokenResponseBody
}

type RevokePxfsAccessTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokePxfsAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokePxfsAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokePxfsAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *RevokePxfsAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokePxfsAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokePxfsAccessTokenResponse) GetBody() *RevokePxfsAccessTokenResponseBody {
	return s.Body
}

func (s *RevokePxfsAccessTokenResponse) SetHeaders(v map[string]*string) *RevokePxfsAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *RevokePxfsAccessTokenResponse) SetStatusCode(v int32) *RevokePxfsAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePxfsAccessTokenResponse) SetBody(v *RevokePxfsAccessTokenResponseBody) *RevokePxfsAccessTokenResponse {
	s.Body = v
	return s
}

func (s *RevokePxfsAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
