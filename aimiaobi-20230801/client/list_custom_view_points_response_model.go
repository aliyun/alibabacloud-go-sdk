// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomViewPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomViewPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomViewPointsResponseBody) *ListCustomViewPointsResponse
	GetBody() *ListCustomViewPointsResponseBody
}

type ListCustomViewPointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomViewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomViewPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomViewPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomViewPointsResponse) GetBody() *ListCustomViewPointsResponseBody {
	return s.Body
}

func (s *ListCustomViewPointsResponse) SetHeaders(v map[string]*string) *ListCustomViewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomViewPointsResponse) SetStatusCode(v int32) *ListCustomViewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomViewPointsResponse) SetBody(v *ListCustomViewPointsResponseBody) *ListCustomViewPointsResponse {
	s.Body = v
	return s
}

func (s *ListCustomViewPointsResponse) Validate() error {
	return dara.Validate(s)
}
