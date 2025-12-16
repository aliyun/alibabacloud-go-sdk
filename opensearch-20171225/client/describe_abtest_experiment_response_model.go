// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeABTestExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeABTestExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeABTestExperimentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeABTestExperimentResponseBody) *DescribeABTestExperimentResponse
	GetBody() *DescribeABTestExperimentResponseBody
}

type DescribeABTestExperimentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeABTestExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeABTestExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeABTestExperimentResponse) GetBody() *DescribeABTestExperimentResponseBody {
	return s.Body
}

func (s *DescribeABTestExperimentResponse) SetHeaders(v map[string]*string) *DescribeABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestExperimentResponse) SetStatusCode(v int32) *DescribeABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeABTestExperimentResponse) SetBody(v *DescribeABTestExperimentResponseBody) *DescribeABTestExperimentResponse {
	s.Body = v
	return s
}

func (s *DescribeABTestExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
