// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsOptimizationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApsOptimizationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApsOptimizationTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListApsOptimizationTasksResponseBody) *ListApsOptimizationTasksResponse
	GetBody() *ListApsOptimizationTasksResponseBody
}

type ListApsOptimizationTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApsOptimizationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApsOptimizationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationTasksResponse) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApsOptimizationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApsOptimizationTasksResponse) GetBody() *ListApsOptimizationTasksResponseBody {
	return s.Body
}

func (s *ListApsOptimizationTasksResponse) SetHeaders(v map[string]*string) *ListApsOptimizationTasksResponse {
	s.Headers = v
	return s
}

func (s *ListApsOptimizationTasksResponse) SetStatusCode(v int32) *ListApsOptimizationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApsOptimizationTasksResponse) SetBody(v *ListApsOptimizationTasksResponseBody) *ListApsOptimizationTasksResponse {
	s.Body = v
	return s
}

func (s *ListApsOptimizationTasksResponse) Validate() error {
	return dara.Validate(s)
}
