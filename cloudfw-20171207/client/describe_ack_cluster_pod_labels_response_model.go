// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterPodLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAckClusterPodLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAckClusterPodLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAckClusterPodLabelsResponseBody) *DescribeAckClusterPodLabelsResponse
	GetBody() *DescribeAckClusterPodLabelsResponseBody
}

type DescribeAckClusterPodLabelsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAckClusterPodLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAckClusterPodLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterPodLabelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterPodLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAckClusterPodLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAckClusterPodLabelsResponse) GetBody() *DescribeAckClusterPodLabelsResponseBody {
	return s.Body
}

func (s *DescribeAckClusterPodLabelsResponse) SetHeaders(v map[string]*string) *DescribeAckClusterPodLabelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckClusterPodLabelsResponse) SetStatusCode(v int32) *DescribeAckClusterPodLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAckClusterPodLabelsResponse) SetBody(v *DescribeAckClusterPodLabelsResponseBody) *DescribeAckClusterPodLabelsResponse {
	s.Body = v
	return s
}

func (s *DescribeAckClusterPodLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
