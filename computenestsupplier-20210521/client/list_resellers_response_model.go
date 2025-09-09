// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResellersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResellersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResellersResponse
	GetStatusCode() *int32
	SetBody(v *ListResellersResponseBody) *ListResellersResponse
	GetBody() *ListResellersResponseBody
}

type ListResellersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResellersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResellersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResellersResponse) GoString() string {
	return s.String()
}

func (s *ListResellersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResellersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResellersResponse) GetBody() *ListResellersResponseBody {
	return s.Body
}

func (s *ListResellersResponse) SetHeaders(v map[string]*string) *ListResellersResponse {
	s.Headers = v
	return s
}

func (s *ListResellersResponse) SetStatusCode(v int32) *ListResellersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResellersResponse) SetBody(v *ListResellersResponseBody) *ListResellersResponse {
	s.Body = v
	return s
}

func (s *ListResellersResponse) Validate() error {
	return dara.Validate(s)
}
