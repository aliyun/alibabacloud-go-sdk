// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeStatInfoResponseBody) *DescribeEdgeStatInfoResponse
	GetBody() *DescribeEdgeStatInfoResponseBody
}

type DescribeEdgeStatInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeStatInfoResponse) GetBody() *DescribeEdgeStatInfoResponseBody {
	return s.Body
}

func (s *DescribeEdgeStatInfoResponse) SetHeaders(v map[string]*string) *DescribeEdgeStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeStatInfoResponse) SetStatusCode(v int32) *DescribeEdgeStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeStatInfoResponse) SetBody(v *DescribeEdgeStatInfoResponseBody) *DescribeEdgeStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
