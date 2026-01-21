// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantBsnCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantBsnCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantBsnCodeResponse
	GetStatusCode() *int32
	SetBody(v *GrantBsnCodeResponseBody) *GrantBsnCodeResponse
	GetBody() *GrantBsnCodeResponseBody
}

type GrantBsnCodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantBsnCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantBsnCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantBsnCodeResponse) GoString() string {
	return s.String()
}

func (s *GrantBsnCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantBsnCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantBsnCodeResponse) GetBody() *GrantBsnCodeResponseBody {
	return s.Body
}

func (s *GrantBsnCodeResponse) SetHeaders(v map[string]*string) *GrantBsnCodeResponse {
	s.Headers = v
	return s
}

func (s *GrantBsnCodeResponse) SetStatusCode(v int32) *GrantBsnCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantBsnCodeResponse) SetBody(v *GrantBsnCodeResponseBody) *GrantBsnCodeResponse {
	s.Body = v
	return s
}

func (s *GrantBsnCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
