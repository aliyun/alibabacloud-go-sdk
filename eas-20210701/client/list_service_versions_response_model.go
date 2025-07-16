// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceVersionsResponseBody) *ListServiceVersionsResponse
	GetBody() *ListServiceVersionsResponseBody
}

type ListServiceVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceVersionsResponse) GetBody() *ListServiceVersionsResponseBody {
	return s.Body
}

func (s *ListServiceVersionsResponse) SetHeaders(v map[string]*string) *ListServiceVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceVersionsResponse) SetStatusCode(v int32) *ListServiceVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceVersionsResponse) SetBody(v *ListServiceVersionsResponseBody) *ListServiceVersionsResponse {
	s.Body = v
	return s
}

func (s *ListServiceVersionsResponse) Validate() error {
	return dara.Validate(s)
}
