// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeyVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKeyVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKeyVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListKeyVersionsResponseBody) *ListKeyVersionsResponse
	GetBody() *ListKeyVersionsResponseBody
}

type ListKeyVersionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeyVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKeyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKeyVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKeyVersionsResponse) GetBody() *ListKeyVersionsResponseBody {
	return s.Body
}

func (s *ListKeyVersionsResponse) SetHeaders(v map[string]*string) *ListKeyVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListKeyVersionsResponse) SetStatusCode(v int32) *ListKeyVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeyVersionsResponse) SetBody(v *ListKeyVersionsResponseBody) *ListKeyVersionsResponse {
	s.Body = v
	return s
}

func (s *ListKeyVersionsResponse) Validate() error {
	return dara.Validate(s)
}
