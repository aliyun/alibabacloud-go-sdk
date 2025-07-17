// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskOperationLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskOperationLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskOperationLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskOperationLogsResponseBody) *ListTaskOperationLogsResponse
	GetBody() *ListTaskOperationLogsResponseBody
}

type ListTaskOperationLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskOperationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskOperationLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskOperationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskOperationLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskOperationLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskOperationLogsResponse) GetBody() *ListTaskOperationLogsResponseBody {
	return s.Body
}

func (s *ListTaskOperationLogsResponse) SetHeaders(v map[string]*string) *ListTaskOperationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskOperationLogsResponse) SetStatusCode(v int32) *ListTaskOperationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskOperationLogsResponse) SetBody(v *ListTaskOperationLogsResponseBody) *ListTaskOperationLogsResponse {
	s.Body = v
	return s
}

func (s *ListTaskOperationLogsResponse) Validate() error {
	return dara.Validate(s)
}
