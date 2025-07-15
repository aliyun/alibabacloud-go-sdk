// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionRiskyTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutionRiskyTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutionRiskyTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutionRiskyTasksResponseBody) *ListExecutionRiskyTasksResponse
	GetBody() *ListExecutionRiskyTasksResponseBody
}

type ListExecutionRiskyTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionRiskyTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutionRiskyTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionRiskyTasksResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutionRiskyTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutionRiskyTasksResponse) GetBody() *ListExecutionRiskyTasksResponseBody {
	return s.Body
}

func (s *ListExecutionRiskyTasksResponse) SetHeaders(v map[string]*string) *ListExecutionRiskyTasksResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionRiskyTasksResponse) SetStatusCode(v int32) *ListExecutionRiskyTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionRiskyTasksResponse) SetBody(v *ListExecutionRiskyTasksResponseBody) *ListExecutionRiskyTasksResponse {
	s.Body = v
	return s
}

func (s *ListExecutionRiskyTasksResponse) Validate() error {
	return dara.Validate(s)
}
