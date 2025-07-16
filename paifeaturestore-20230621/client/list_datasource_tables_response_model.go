// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourceTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasourceTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasourceTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasourceTablesResponseBody) *ListDatasourceTablesResponse
	GetBody() *ListDatasourceTablesResponseBody
}

type ListDatasourceTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasourceTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDatasourceTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasourceTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasourceTablesResponse) GetBody() *ListDatasourceTablesResponseBody {
	return s.Body
}

func (s *ListDatasourceTablesResponse) SetHeaders(v map[string]*string) *ListDatasourceTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDatasourceTablesResponse) SetStatusCode(v int32) *ListDatasourceTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasourceTablesResponse) SetBody(v *ListDatasourceTablesResponseBody) *ListDatasourceTablesResponse {
	s.Body = v
	return s
}

func (s *ListDatasourceTablesResponse) Validate() error {
	return dara.Validate(s)
}
