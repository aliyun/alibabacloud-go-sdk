// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthAndActiveWithHidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthAndActiveWithHidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthAndActiveWithHidResponse
	GetStatusCode() *int32
	SetBody(v *AuthAndActiveWithHidResponseBody) *AuthAndActiveWithHidResponse
	GetBody() *AuthAndActiveWithHidResponseBody
}

type AuthAndActiveWithHidResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthAndActiveWithHidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthAndActiveWithHidResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthAndActiveWithHidResponse) GoString() string {
	return s.String()
}

func (s *AuthAndActiveWithHidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthAndActiveWithHidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthAndActiveWithHidResponse) GetBody() *AuthAndActiveWithHidResponseBody {
	return s.Body
}

func (s *AuthAndActiveWithHidResponse) SetHeaders(v map[string]*string) *AuthAndActiveWithHidResponse {
	s.Headers = v
	return s
}

func (s *AuthAndActiveWithHidResponse) SetStatusCode(v int32) *AuthAndActiveWithHidResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthAndActiveWithHidResponse) SetBody(v *AuthAndActiveWithHidResponseBody) *AuthAndActiveWithHidResponse {
	s.Body = v
	return s
}

func (s *AuthAndActiveWithHidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
