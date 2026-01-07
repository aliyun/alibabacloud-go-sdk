// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactEditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountContactEditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountContactEditResponse
	GetStatusCode() *int32
	SetBody(v *AccountContactEditResponseBody) *AccountContactEditResponse
	GetBody() *AccountContactEditResponseBody
}

type AccountContactEditResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountContactEditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountContactEditResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountContactEditResponse) GoString() string {
	return s.String()
}

func (s *AccountContactEditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountContactEditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountContactEditResponse) GetBody() *AccountContactEditResponseBody {
	return s.Body
}

func (s *AccountContactEditResponse) SetHeaders(v map[string]*string) *AccountContactEditResponse {
	s.Headers = v
	return s
}

func (s *AccountContactEditResponse) SetStatusCode(v int32) *AccountContactEditResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountContactEditResponse) SetBody(v *AccountContactEditResponseBody) *AccountContactEditResponse {
	s.Body = v
	return s
}

func (s *AccountContactEditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
