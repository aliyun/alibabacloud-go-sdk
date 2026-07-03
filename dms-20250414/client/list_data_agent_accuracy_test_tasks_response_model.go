// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentAccuracyTestTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentAccuracyTestTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentAccuracyTestTasksResponseBody) *ListDataAgentAccuracyTestTasksResponse
	GetBody() *ListDataAgentAccuracyTestTasksResponseBody
}

type ListDataAgentAccuracyTestTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentAccuracyTestTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentAccuracyTestTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentAccuracyTestTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentAccuracyTestTasksResponse) GetBody() *ListDataAgentAccuracyTestTasksResponseBody {
	return s.Body
}

func (s *ListDataAgentAccuracyTestTasksResponse) SetHeaders(v map[string]*string) *ListDataAgentAccuracyTestTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponse) SetStatusCode(v int32) *ListDataAgentAccuracyTestTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponse) SetBody(v *ListDataAgentAccuracyTestTasksResponseBody) *ListDataAgentAccuracyTestTasksResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
