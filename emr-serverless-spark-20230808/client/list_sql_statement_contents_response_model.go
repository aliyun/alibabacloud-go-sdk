// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSqlStatementContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSqlStatementContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSqlStatementContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListSqlStatementContentsResponseBody) *ListSqlStatementContentsResponse
	GetBody() *ListSqlStatementContentsResponseBody
}

type ListSqlStatementContentsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSqlStatementContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSqlStatementContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSqlStatementContentsResponse) GoString() string {
	return s.String()
}

func (s *ListSqlStatementContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSqlStatementContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSqlStatementContentsResponse) GetBody() *ListSqlStatementContentsResponseBody {
	return s.Body
}

func (s *ListSqlStatementContentsResponse) SetHeaders(v map[string]*string) *ListSqlStatementContentsResponse {
	s.Headers = v
	return s
}

func (s *ListSqlStatementContentsResponse) SetStatusCode(v int32) *ListSqlStatementContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSqlStatementContentsResponse) SetBody(v *ListSqlStatementContentsResponseBody) *ListSqlStatementContentsResponse {
	s.Body = v
	return s
}

func (s *ListSqlStatementContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
