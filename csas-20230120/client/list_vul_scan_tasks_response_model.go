// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulScanTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVulScanTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVulScanTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListVulScanTasksResponseBody) *ListVulScanTasksResponse
	GetBody() *ListVulScanTasksResponseBody
}

type ListVulScanTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVulScanTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVulScanTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanTasksResponse) GoString() string {
	return s.String()
}

func (s *ListVulScanTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVulScanTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVulScanTasksResponse) GetBody() *ListVulScanTasksResponseBody {
	return s.Body
}

func (s *ListVulScanTasksResponse) SetHeaders(v map[string]*string) *ListVulScanTasksResponse {
	s.Headers = v
	return s
}

func (s *ListVulScanTasksResponse) SetStatusCode(v int32) *ListVulScanTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVulScanTasksResponse) SetBody(v *ListVulScanTasksResponseBody) *ListVulScanTasksResponse {
	s.Body = v
	return s
}

func (s *ListVulScanTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
