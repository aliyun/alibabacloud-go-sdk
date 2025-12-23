// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeploymentJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeploymentJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeploymentJobStatusResponseBody) *DescribeDeploymentJobStatusResponse
	GetBody() *DescribeDeploymentJobStatusResponseBody
}

type DescribeDeploymentJobStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeploymentJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeploymentJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeploymentJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeploymentJobStatusResponse) GetBody() *DescribeDeploymentJobStatusResponseBody {
	return s.Body
}

func (s *DescribeDeploymentJobStatusResponse) SetHeaders(v map[string]*string) *DescribeDeploymentJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeploymentJobStatusResponse) SetStatusCode(v int32) *DescribeDeploymentJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponse) SetBody(v *DescribeDeploymentJobStatusResponseBody) *DescribeDeploymentJobStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDeploymentJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
