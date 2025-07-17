// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLineagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLineagesResponse
	GetStatusCode() *int32
	SetBody(v *ListLineagesResponseBody) *ListLineagesResponse
	GetBody() *ListLineagesResponseBody
}

type ListLineagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLineagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLineagesResponse) GoString() string {
	return s.String()
}

func (s *ListLineagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLineagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLineagesResponse) GetBody() *ListLineagesResponseBody {
	return s.Body
}

func (s *ListLineagesResponse) SetHeaders(v map[string]*string) *ListLineagesResponse {
	s.Headers = v
	return s
}

func (s *ListLineagesResponse) SetStatusCode(v int32) *ListLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLineagesResponse) SetBody(v *ListLineagesResponseBody) *ListLineagesResponse {
	s.Body = v
	return s
}

func (s *ListLineagesResponse) Validate() error {
	return dara.Validate(s)
}
