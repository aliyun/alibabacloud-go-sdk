// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePxfsAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePxfsAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreatePxfsAccessTokenResponseBody) *CreatePxfsAccessTokenResponse
	GetBody() *CreatePxfsAccessTokenResponseBody
}

type CreatePxfsAccessTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePxfsAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePxfsAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *CreatePxfsAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePxfsAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePxfsAccessTokenResponse) GetBody() *CreatePxfsAccessTokenResponseBody {
	return s.Body
}

func (s *CreatePxfsAccessTokenResponse) SetHeaders(v map[string]*string) *CreatePxfsAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *CreatePxfsAccessTokenResponse) SetStatusCode(v int32) *CreatePxfsAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePxfsAccessTokenResponse) SetBody(v *CreatePxfsAccessTokenResponseBody) *CreatePxfsAccessTokenResponse {
	s.Body = v
	return s
}

func (s *CreatePxfsAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
