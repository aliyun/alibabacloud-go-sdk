// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWithSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWithSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWithSQLResponse
	GetStatusCode() *int32
	SetBody(v *QueryWithSQLResponseBody) *QueryWithSQLResponse
	GetBody() *QueryWithSQLResponseBody
}

type QueryWithSQLResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWithSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWithSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWithSQLResponse) GoString() string {
	return s.String()
}

func (s *QueryWithSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWithSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWithSQLResponse) GetBody() *QueryWithSQLResponseBody {
	return s.Body
}

func (s *QueryWithSQLResponse) SetHeaders(v map[string]*string) *QueryWithSQLResponse {
	s.Headers = v
	return s
}

func (s *QueryWithSQLResponse) SetStatusCode(v int32) *QueryWithSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWithSQLResponse) SetBody(v *QueryWithSQLResponseBody) *QueryWithSQLResponse {
	s.Body = v
	return s
}

func (s *QueryWithSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
