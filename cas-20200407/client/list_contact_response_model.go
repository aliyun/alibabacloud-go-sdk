// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContactResponse
	GetStatusCode() *int32
	SetBody(v *ListContactResponseBody) *ListContactResponse
	GetBody() *ListContactResponseBody
}

type ListContactResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContactResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContactResponse) GoString() string {
	return s.String()
}

func (s *ListContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContactResponse) GetBody() *ListContactResponseBody {
	return s.Body
}

func (s *ListContactResponse) SetHeaders(v map[string]*string) *ListContactResponse {
	s.Headers = v
	return s
}

func (s *ListContactResponse) SetStatusCode(v int32) *ListContactResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactResponse) SetBody(v *ListContactResponseBody) *ListContactResponse {
	s.Body = v
	return s
}

func (s *ListContactResponse) Validate() error {
	return dara.Validate(s)
}
