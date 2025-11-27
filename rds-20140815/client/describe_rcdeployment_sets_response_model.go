// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCDeploymentSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCDeploymentSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCDeploymentSetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCDeploymentSetsResponseBody) *DescribeRCDeploymentSetsResponse
	GetBody() *DescribeRCDeploymentSetsResponseBody
}

type DescribeRCDeploymentSetsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCDeploymentSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCDeploymentSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCDeploymentSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCDeploymentSetsResponse) GetBody() *DescribeRCDeploymentSetsResponseBody {
	return s.Body
}

func (s *DescribeRCDeploymentSetsResponse) SetHeaders(v map[string]*string) *DescribeRCDeploymentSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCDeploymentSetsResponse) SetStatusCode(v int32) *DescribeRCDeploymentSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponse) SetBody(v *DescribeRCDeploymentSetsResponseBody) *DescribeRCDeploymentSetsResponse {
	s.Body = v
	return s
}

func (s *DescribeRCDeploymentSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
