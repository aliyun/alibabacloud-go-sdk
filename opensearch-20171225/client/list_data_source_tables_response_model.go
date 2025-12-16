// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceTablesResponseBody) *ListDataSourceTablesResponse
	GetBody() *ListDataSourceTablesResponseBody
}

type ListDataSourceTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceTablesResponse) GetBody() *ListDataSourceTablesResponseBody {
	return s.Body
}

func (s *ListDataSourceTablesResponse) SetHeaders(v map[string]*string) *ListDataSourceTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTablesResponse) SetStatusCode(v int32) *ListDataSourceTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTablesResponse) SetBody(v *ListDataSourceTablesResponseBody) *ListDataSourceTablesResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
