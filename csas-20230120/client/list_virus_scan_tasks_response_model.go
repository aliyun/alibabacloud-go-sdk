// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanTasksResponseBody) *ListVirusScanTasksResponse
	GetBody() *ListVirusScanTasksResponseBody
}

type ListVirusScanTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTasksResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanTasksResponse) GetBody() *ListVirusScanTasksResponseBody {
	return s.Body
}

func (s *ListVirusScanTasksResponse) SetHeaders(v map[string]*string) *ListVirusScanTasksResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanTasksResponse) SetStatusCode(v int32) *ListVirusScanTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanTasksResponse) SetBody(v *ListVirusScanTasksResponseBody) *ListVirusScanTasksResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
