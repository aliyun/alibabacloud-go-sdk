// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstanceOperationLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskInstanceOperationLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskInstanceOperationLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskInstanceOperationLogsResponseBody) *ListTaskInstanceOperationLogsResponse
	GetBody() *ListTaskInstanceOperationLogsResponseBody
}

type ListTaskInstanceOperationLogsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskInstanceOperationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskInstanceOperationLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstanceOperationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskInstanceOperationLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskInstanceOperationLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskInstanceOperationLogsResponse) GetBody() *ListTaskInstanceOperationLogsResponseBody {
	return s.Body
}

func (s *ListTaskInstanceOperationLogsResponse) SetHeaders(v map[string]*string) *ListTaskInstanceOperationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskInstanceOperationLogsResponse) SetStatusCode(v int32) *ListTaskInstanceOperationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskInstanceOperationLogsResponse) SetBody(v *ListTaskInstanceOperationLogsResponseBody) *ListTaskInstanceOperationLogsResponse {
	s.Body = v
	return s
}

func (s *ListTaskInstanceOperationLogsResponse) Validate() error {
	return dara.Validate(s)
}
