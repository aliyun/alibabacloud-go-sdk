// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIndexesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIndexesResponse
	GetStatusCode() *int32
	SetBody(v *ListIndexesResponseBody) *ListIndexesResponse
	GetBody() *ListIndexesResponseBody
}

type ListIndexesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponse) GoString() string {
	return s.String()
}

func (s *ListIndexesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIndexesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIndexesResponse) GetBody() *ListIndexesResponseBody {
	return s.Body
}

func (s *ListIndexesResponse) SetHeaders(v map[string]*string) *ListIndexesResponse {
	s.Headers = v
	return s
}

func (s *ListIndexesResponse) SetStatusCode(v int32) *ListIndexesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexesResponse) SetBody(v *ListIndexesResponseBody) *ListIndexesResponse {
	s.Body = v
	return s
}

func (s *ListIndexesResponse) Validate() error {
	return dara.Validate(s)
}
