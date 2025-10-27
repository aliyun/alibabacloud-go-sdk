// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInclinedNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInclinedNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInclinedNodesResponseBody) *DescribeInclinedNodesResponse
	GetBody() *DescribeInclinedNodesResponseBody
}

type DescribeInclinedNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInclinedNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInclinedNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInclinedNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInclinedNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInclinedNodesResponse) GetBody() *DescribeInclinedNodesResponseBody {
	return s.Body
}

func (s *DescribeInclinedNodesResponse) SetHeaders(v map[string]*string) *DescribeInclinedNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInclinedNodesResponse) SetStatusCode(v int32) *DescribeInclinedNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInclinedNodesResponse) SetBody(v *DescribeInclinedNodesResponseBody) *DescribeInclinedNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeInclinedNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
