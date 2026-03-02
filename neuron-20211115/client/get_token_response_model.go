// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateTokenResult) *GetTokenResponse
	GetBody() *CreateTokenResult
}

type GetTokenResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenResponse) GetBody() *CreateTokenResult {
	return s.Body
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *CreateTokenResult) *GetTokenResponse {
	s.Body = v
	return s
}

func (s *GetTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
