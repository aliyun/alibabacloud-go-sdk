// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlowSQLRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSlowSQLRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSlowSQLRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListSlowSQLRecordsResponseBody) *ListSlowSQLRecordsResponse
	GetBody() *ListSlowSQLRecordsResponseBody
}

type ListSlowSQLRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSlowSQLRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSlowSQLRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSlowSQLRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListSlowSQLRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSlowSQLRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSlowSQLRecordsResponse) GetBody() *ListSlowSQLRecordsResponseBody {
	return s.Body
}

func (s *ListSlowSQLRecordsResponse) SetHeaders(v map[string]*string) *ListSlowSQLRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListSlowSQLRecordsResponse) SetStatusCode(v int32) *ListSlowSQLRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlowSQLRecordsResponse) SetBody(v *ListSlowSQLRecordsResponseBody) *ListSlowSQLRecordsResponse {
	s.Body = v
	return s
}

func (s *ListSlowSQLRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
