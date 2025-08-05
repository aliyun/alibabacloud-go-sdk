// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsTablesResponseBody) *ListMmsTablesResponse
	GetBody() *ListMmsTablesResponseBody
}

type ListMmsTablesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsTablesResponse) GetBody() *ListMmsTablesResponseBody {
	return s.Body
}

func (s *ListMmsTablesResponse) SetHeaders(v map[string]*string) *ListMmsTablesResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTablesResponse) SetStatusCode(v int32) *ListMmsTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTablesResponse) SetBody(v *ListMmsTablesResponseBody) *ListMmsTablesResponse {
	s.Body = v
	return s
}

func (s *ListMmsTablesResponse) Validate() error {
	return dara.Validate(s)
}
