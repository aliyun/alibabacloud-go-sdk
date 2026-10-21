// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticGlobalSqlExecutionHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticGlobalSqlExecutionHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticGlobalSqlExecutionHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticGlobalSqlExecutionHistoriesResponseBody) *ListAgenticGlobalSqlExecutionHistoriesResponse
	GetBody() *ListAgenticGlobalSqlExecutionHistoriesResponseBody
}

type ListAgenticGlobalSqlExecutionHistoriesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticGlobalSqlExecutionHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticGlobalSqlExecutionHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticGlobalSqlExecutionHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) GetBody() *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	return s.Body
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) SetHeaders(v map[string]*string) *ListAgenticGlobalSqlExecutionHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) SetStatusCode(v int32) *ListAgenticGlobalSqlExecutionHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) SetBody(v *ListAgenticGlobalSqlExecutionHistoriesResponseBody) *ListAgenticGlobalSqlExecutionHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
