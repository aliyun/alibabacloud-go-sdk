// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticSqlExecutionHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticSqlExecutionHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticSqlExecutionHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticSqlExecutionHistoriesResponseBody) *ListAgenticSqlExecutionHistoriesResponse
	GetBody() *ListAgenticSqlExecutionHistoriesResponseBody
}

type ListAgenticSqlExecutionHistoriesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticSqlExecutionHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticSqlExecutionHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSqlExecutionHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticSqlExecutionHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticSqlExecutionHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticSqlExecutionHistoriesResponse) GetBody() *ListAgenticSqlExecutionHistoriesResponseBody {
	return s.Body
}

func (s *ListAgenticSqlExecutionHistoriesResponse) SetHeaders(v map[string]*string) *ListAgenticSqlExecutionHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponse) SetStatusCode(v int32) *ListAgenticSqlExecutionHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponse) SetBody(v *ListAgenticSqlExecutionHistoriesResponseBody) *ListAgenticSqlExecutionHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
