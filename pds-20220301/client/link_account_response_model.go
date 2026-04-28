// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LinkAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LinkAccountResponse
	GetStatusCode() *int32
	SetBody(v *Token) *LinkAccountResponse
	GetBody() *Token
}

type LinkAccountResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Token             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s LinkAccountResponse) GoString() string {
	return s.String()
}

func (s *LinkAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LinkAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LinkAccountResponse) GetBody() *Token {
	return s.Body
}

func (s *LinkAccountResponse) SetHeaders(v map[string]*string) *LinkAccountResponse {
	s.Headers = v
	return s
}

func (s *LinkAccountResponse) SetStatusCode(v int32) *LinkAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkAccountResponse) SetBody(v *Token) *LinkAccountResponse {
	s.Body = v
	return s
}

func (s *LinkAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
