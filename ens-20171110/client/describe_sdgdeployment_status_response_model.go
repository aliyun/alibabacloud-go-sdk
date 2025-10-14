// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGDeploymentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSDGDeploymentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSDGDeploymentStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSDGDeploymentStatusResponseBody) *DescribeSDGDeploymentStatusResponse
	GetBody() *DescribeSDGDeploymentStatusResponseBody
}

type DescribeSDGDeploymentStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSDGDeploymentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSDGDeploymentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGDeploymentStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDGDeploymentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSDGDeploymentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSDGDeploymentStatusResponse) GetBody() *DescribeSDGDeploymentStatusResponseBody {
	return s.Body
}

func (s *DescribeSDGDeploymentStatusResponse) SetHeaders(v map[string]*string) *DescribeSDGDeploymentStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDGDeploymentStatusResponse) SetStatusCode(v int32) *DescribeSDGDeploymentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponse) SetBody(v *DescribeSDGDeploymentStatusResponseBody) *DescribeSDGDeploymentStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSDGDeploymentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
