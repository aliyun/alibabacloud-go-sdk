// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSimulationTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSimulationTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSimulationTaskListResponseBody) *DescribeSimulationTaskListResponse
	GetBody() *DescribeSimulationTaskListResponseBody
}

type DescribeSimulationTaskListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimulationTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSimulationTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimulationTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSimulationTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSimulationTaskListResponse) GetBody() *DescribeSimulationTaskListResponseBody {
	return s.Body
}

func (s *DescribeSimulationTaskListResponse) SetHeaders(v map[string]*string) *DescribeSimulationTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimulationTaskListResponse) SetStatusCode(v int32) *DescribeSimulationTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimulationTaskListResponse) SetBody(v *DescribeSimulationTaskListResponseBody) *DescribeSimulationTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeSimulationTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
