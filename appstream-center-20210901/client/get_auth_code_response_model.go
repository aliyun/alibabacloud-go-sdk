// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthCodeResponseBody) *GetAuthCodeResponse
	GetBody() *GetAuthCodeResponseBody
}

type GetAuthCodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthCodeResponse) GetBody() *GetAuthCodeResponseBody {
	return s.Body
}

func (s *GetAuthCodeResponse) SetHeaders(v map[string]*string) *GetAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAuthCodeResponse) SetStatusCode(v int32) *GetAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthCodeResponse) SetBody(v *GetAuthCodeResponseBody) *GetAuthCodeResponse {
	s.Body = v
	return s
}

func (s *GetAuthCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
