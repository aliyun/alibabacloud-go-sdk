// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationTaskCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSimulationTaskCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSimulationTaskCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSimulationTaskCountResponseBody) *DescribeSimulationTaskCountResponse
	GetBody() *DescribeSimulationTaskCountResponseBody
}

type DescribeSimulationTaskCountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimulationTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSimulationTaskCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationTaskCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimulationTaskCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSimulationTaskCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSimulationTaskCountResponse) GetBody() *DescribeSimulationTaskCountResponseBody {
	return s.Body
}

func (s *DescribeSimulationTaskCountResponse) SetHeaders(v map[string]*string) *DescribeSimulationTaskCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimulationTaskCountResponse) SetStatusCode(v int32) *DescribeSimulationTaskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimulationTaskCountResponse) SetBody(v *DescribeSimulationTaskCountResponseBody) *DescribeSimulationTaskCountResponse {
	s.Body = v
	return s
}

func (s *DescribeSimulationTaskCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
