// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntervenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntervenesResponse
	GetStatusCode() *int32
	SetBody(v *ListIntervenesResponseBody) *ListIntervenesResponse
	GetBody() *ListIntervenesResponseBody
}

type ListIntervenesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervenesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntervenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntervenesResponse) GetBody() *ListIntervenesResponseBody {
	return s.Body
}

func (s *ListIntervenesResponse) SetHeaders(v map[string]*string) *ListIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ListIntervenesResponse) SetStatusCode(v int32) *ListIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervenesResponse) SetBody(v *ListIntervenesResponseBody) *ListIntervenesResponse {
	s.Body = v
	return s
}

func (s *ListIntervenesResponse) Validate() error {
	return dara.Validate(s)
}
