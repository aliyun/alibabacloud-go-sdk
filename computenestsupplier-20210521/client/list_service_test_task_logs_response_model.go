// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestTaskLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceTestTaskLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceTestTaskLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceTestTaskLogsResponseBody) *ListServiceTestTaskLogsResponse
	GetBody() *ListServiceTestTaskLogsResponseBody
}

type ListServiceTestTaskLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTestTaskLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTestTaskLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTaskLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceTestTaskLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceTestTaskLogsResponse) GetBody() *ListServiceTestTaskLogsResponseBody {
	return s.Body
}

func (s *ListServiceTestTaskLogsResponse) SetHeaders(v map[string]*string) *ListServiceTestTaskLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTestTaskLogsResponse) SetStatusCode(v int32) *ListServiceTestTaskLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTestTaskLogsResponse) SetBody(v *ListServiceTestTaskLogsResponseBody) *ListServiceTestTaskLogsResponse {
	s.Body = v
	return s
}

func (s *ListServiceTestTaskLogsResponse) Validate() error {
	return dara.Validate(s)
}
