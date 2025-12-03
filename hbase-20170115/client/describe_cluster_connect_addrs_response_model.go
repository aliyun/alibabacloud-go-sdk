// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterConnectAddrsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterConnectAddrsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterConnectAddrsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterConnectAddrsResponseBody) *DescribeClusterConnectAddrsResponse
	GetBody() *DescribeClusterConnectAddrsResponseBody
}

type DescribeClusterConnectAddrsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterConnectAddrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterConnectAddrsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterConnectAddrsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterConnectAddrsResponse) GetBody() *DescribeClusterConnectAddrsResponseBody {
	return s.Body
}

func (s *DescribeClusterConnectAddrsResponse) SetHeaders(v map[string]*string) *DescribeClusterConnectAddrsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterConnectAddrsResponse) SetStatusCode(v int32) *DescribeClusterConnectAddrsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponse) SetBody(v *DescribeClusterConnectAddrsResponseBody) *DescribeClusterConnectAddrsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterConnectAddrsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
