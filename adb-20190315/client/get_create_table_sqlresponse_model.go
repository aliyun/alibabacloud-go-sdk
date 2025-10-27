// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateTableSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCreateTableSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCreateTableSQLResponse
	GetStatusCode() *int32
	SetBody(v *GetCreateTableSQLResponseBody) *GetCreateTableSQLResponse
	GetBody() *GetCreateTableSQLResponseBody
}

type GetCreateTableSQLResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateTableSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateTableSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCreateTableSQLResponse) GoString() string {
	return s.String()
}

func (s *GetCreateTableSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCreateTableSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCreateTableSQLResponse) GetBody() *GetCreateTableSQLResponseBody {
	return s.Body
}

func (s *GetCreateTableSQLResponse) SetHeaders(v map[string]*string) *GetCreateTableSQLResponse {
	s.Headers = v
	return s
}

func (s *GetCreateTableSQLResponse) SetStatusCode(v int32) *GetCreateTableSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateTableSQLResponse) SetBody(v *GetCreateTableSQLResponseBody) *GetCreateTableSQLResponse {
	s.Body = v
	return s
}

func (s *GetCreateTableSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
