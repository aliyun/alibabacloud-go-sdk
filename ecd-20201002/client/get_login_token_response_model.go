// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoginTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoginTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse
	GetBody() *GetLoginTokenResponseBody
}

type GetLoginTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoginTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoginTokenResponse) GetBody() *GetLoginTokenResponseBody {
	return s.Body
}

func (s *GetLoginTokenResponse) SetHeaders(v map[string]*string) *GetLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLoginTokenResponse) SetStatusCode(v int32) *GetLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginTokenResponse) SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse {
	s.Body = v
	return s
}

func (s *GetLoginTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
