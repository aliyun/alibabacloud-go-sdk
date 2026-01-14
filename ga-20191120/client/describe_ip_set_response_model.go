// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpSetResponseBody) *DescribeIpSetResponse
	GetBody() *DescribeIpSetResponseBody
}

type DescribeIpSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpSetResponse) GetBody() *DescribeIpSetResponseBody {
	return s.Body
}

func (s *DescribeIpSetResponse) SetHeaders(v map[string]*string) *DescribeIpSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpSetResponse) SetStatusCode(v int32) *DescribeIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpSetResponse) SetBody(v *DescribeIpSetResponseBody) *DescribeIpSetResponse {
	s.Body = v
	return s
}

func (s *DescribeIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
