// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterTDEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterTDEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterTDEResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterTDEResponseBody) *DescribeDBClusterTDEResponse
	GetBody() *DescribeDBClusterTDEResponseBody
}

type DescribeDBClusterTDEResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterTDEResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterTDEResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterTDEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterTDEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterTDEResponse) GetBody() *DescribeDBClusterTDEResponseBody {
	return s.Body
}

func (s *DescribeDBClusterTDEResponse) SetHeaders(v map[string]*string) *DescribeDBClusterTDEResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterTDEResponse) SetStatusCode(v int32) *DescribeDBClusterTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterTDEResponse) SetBody(v *DescribeDBClusterTDEResponseBody) *DescribeDBClusterTDEResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterTDEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
