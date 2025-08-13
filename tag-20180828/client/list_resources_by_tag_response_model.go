// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesByTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcesByTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcesByTagResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcesByTagResponseBody) *ListResourcesByTagResponse
	GetBody() *ListResourcesByTagResponseBody
}

type ListResourcesByTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesByTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesByTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByTagResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcesByTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcesByTagResponse) GetBody() *ListResourcesByTagResponseBody {
	return s.Body
}

func (s *ListResourcesByTagResponse) SetHeaders(v map[string]*string) *ListResourcesByTagResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesByTagResponse) SetStatusCode(v int32) *ListResourcesByTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesByTagResponse) SetBody(v *ListResourcesByTagResponseBody) *ListResourcesByTagResponse {
	s.Body = v
	return s
}

func (s *ListResourcesByTagResponse) Validate() error {
	return dara.Validate(s)
}
