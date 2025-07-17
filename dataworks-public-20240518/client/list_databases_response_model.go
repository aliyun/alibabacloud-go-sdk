// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabasesResponseBody) *ListDatabasesResponse
	GetBody() *ListDatabasesResponseBody
}

type ListDatabasesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabasesResponse) GetBody() *ListDatabasesResponseBody {
	return s.Body
}

func (s *ListDatabasesResponse) SetHeaders(v map[string]*string) *ListDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListDatabasesResponse) SetStatusCode(v int32) *ListDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabasesResponse) SetBody(v *ListDatabasesResponseBody) *ListDatabasesResponse {
	s.Body = v
	return s
}

func (s *ListDatabasesResponse) Validate() error {
	return dara.Validate(s)
}
