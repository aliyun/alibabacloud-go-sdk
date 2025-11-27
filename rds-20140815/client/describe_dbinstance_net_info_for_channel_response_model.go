// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoForChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoForChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceNetInfoForChannelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceNetInfoForChannelResponseBody) *DescribeDBInstanceNetInfoForChannelResponse
	GetBody() *DescribeDBInstanceNetInfoForChannelResponseBody
}

type DescribeDBInstanceNetInfoForChannelResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceNetInfoForChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) GetBody() *DescribeDBInstanceNetInfoForChannelResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoForChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) SetStatusCode(v int32) *DescribeDBInstanceNetInfoForChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) SetBody(v *DescribeDBInstanceNetInfoForChannelResponseBody) *DescribeDBInstanceNetInfoForChannelResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
