// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListTablesResponseBody) *ListTablesResponse
	GetBody() *ListTablesResponseBody
}

type ListTablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTablesResponse) GetBody() *ListTablesResponseBody {
	return s.Body
}

func (s *ListTablesResponse) SetHeaders(v map[string]*string) *ListTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTablesResponse) SetStatusCode(v int32) *ListTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablesResponse) SetBody(v *ListTablesResponseBody) *ListTablesResponse {
	s.Body = v
	return s
}

func (s *ListTablesResponse) Validate() error {
	return dara.Validate(s)
}
