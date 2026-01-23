// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityWatchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityWatchTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityWatchTasksResponseBody) *ListQualityWatchTasksResponse
	GetBody() *ListQualityWatchTasksResponseBody
}

type ListQualityWatchTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityWatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityWatchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponse) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityWatchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityWatchTasksResponse) GetBody() *ListQualityWatchTasksResponseBody {
	return s.Body
}

func (s *ListQualityWatchTasksResponse) SetHeaders(v map[string]*string) *ListQualityWatchTasksResponse {
	s.Headers = v
	return s
}

func (s *ListQualityWatchTasksResponse) SetStatusCode(v int32) *ListQualityWatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityWatchTasksResponse) SetBody(v *ListQualityWatchTasksResponseBody) *ListQualityWatchTasksResponse {
	s.Body = v
	return s
}

func (s *ListQualityWatchTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
