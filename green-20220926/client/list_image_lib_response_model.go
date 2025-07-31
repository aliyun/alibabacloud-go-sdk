// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageLibResponse
	GetStatusCode() *int32
	SetBody(v *ListImageLibResponseBody) *ListImageLibResponse
	GetBody() *ListImageLibResponseBody
}

type ListImageLibResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageLibResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageLibResponse) GoString() string {
	return s.String()
}

func (s *ListImageLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageLibResponse) GetBody() *ListImageLibResponseBody {
	return s.Body
}

func (s *ListImageLibResponse) SetHeaders(v map[string]*string) *ListImageLibResponse {
	s.Headers = v
	return s
}

func (s *ListImageLibResponse) SetStatusCode(v int32) *ListImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageLibResponse) SetBody(v *ListImageLibResponseBody) *ListImageLibResponse {
	s.Body = v
	return s
}

func (s *ListImageLibResponse) Validate() error {
	return dara.Validate(s)
}
