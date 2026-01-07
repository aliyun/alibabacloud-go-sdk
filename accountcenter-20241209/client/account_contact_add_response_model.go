// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountContactAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountContactAddResponse
	GetStatusCode() *int32
	SetBody(v *AccountContactAddResponseBody) *AccountContactAddResponse
	GetBody() *AccountContactAddResponseBody
}

type AccountContactAddResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountContactAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountContactAddResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountContactAddResponse) GoString() string {
	return s.String()
}

func (s *AccountContactAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountContactAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountContactAddResponse) GetBody() *AccountContactAddResponseBody {
	return s.Body
}

func (s *AccountContactAddResponse) SetHeaders(v map[string]*string) *AccountContactAddResponse {
	s.Headers = v
	return s
}

func (s *AccountContactAddResponse) SetStatusCode(v int32) *AccountContactAddResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountContactAddResponse) SetBody(v *AccountContactAddResponseBody) *AccountContactAddResponse {
	s.Body = v
	return s
}

func (s *AccountContactAddResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
