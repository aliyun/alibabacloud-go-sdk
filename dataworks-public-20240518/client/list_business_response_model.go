// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBusinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBusinessResponse
	GetStatusCode() *int32
	SetBody(v *ListBusinessResponseBody) *ListBusinessResponse
	GetBody() *ListBusinessResponseBody
}

type ListBusinessResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBusinessResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessResponse) GoString() string {
	return s.String()
}

func (s *ListBusinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBusinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBusinessResponse) GetBody() *ListBusinessResponseBody {
	return s.Body
}

func (s *ListBusinessResponse) SetHeaders(v map[string]*string) *ListBusinessResponse {
	s.Headers = v
	return s
}

func (s *ListBusinessResponse) SetStatusCode(v int32) *ListBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusinessResponse) SetBody(v *ListBusinessResponseBody) *ListBusinessResponse {
	s.Body = v
	return s
}

func (s *ListBusinessResponse) Validate() error {
	return dara.Validate(s)
}
