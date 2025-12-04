// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigInXMLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterConfigInXMLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterConfigInXMLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterConfigInXMLResponseBody) *DescribeDBClusterConfigInXMLResponse
	GetBody() *DescribeDBClusterConfigInXMLResponseBody
}

type DescribeDBClusterConfigInXMLResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigInXMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConfigInXMLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigInXMLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigInXMLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterConfigInXMLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterConfigInXMLResponse) GetBody() *DescribeDBClusterConfigInXMLResponseBody {
	return s.Body
}

func (s *DescribeDBClusterConfigInXMLResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigInXMLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponse) SetStatusCode(v int32) *DescribeDBClusterConfigInXMLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponse) SetBody(v *DescribeDBClusterConfigInXMLResponseBody) *DescribeDBClusterConfigInXMLResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
