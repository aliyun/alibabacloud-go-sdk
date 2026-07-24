// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlExecutionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchSqlExecutionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchSqlExecutionResultResponse
	GetStatusCode() *int32
	SetBody(v *FetchSqlExecutionResultResponseBody) *FetchSqlExecutionResultResponse
	GetBody() *FetchSqlExecutionResultResponseBody
}

type FetchSqlExecutionResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchSqlExecutionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchSqlExecutionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlExecutionResultResponse) GoString() string {
	return s.String()
}

func (s *FetchSqlExecutionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchSqlExecutionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchSqlExecutionResultResponse) GetBody() *FetchSqlExecutionResultResponseBody {
	return s.Body
}

func (s *FetchSqlExecutionResultResponse) SetHeaders(v map[string]*string) *FetchSqlExecutionResultResponse {
	s.Headers = v
	return s
}

func (s *FetchSqlExecutionResultResponse) SetStatusCode(v int32) *FetchSqlExecutionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchSqlExecutionResultResponse) SetBody(v *FetchSqlExecutionResultResponseBody) *FetchSqlExecutionResultResponse {
	s.Body = v
	return s
}

func (s *FetchSqlExecutionResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
