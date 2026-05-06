// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnlineUserCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOnlineUserCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOnlineUserCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOnlineUserCountResponseBody) *DescribeOnlineUserCountResponse
	GetBody() *DescribeOnlineUserCountResponseBody
}

type DescribeOnlineUserCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOnlineUserCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOnlineUserCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineUserCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnlineUserCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOnlineUserCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOnlineUserCountResponse) GetBody() *DescribeOnlineUserCountResponseBody {
	return s.Body
}

func (s *DescribeOnlineUserCountResponse) SetHeaders(v map[string]*string) *DescribeOnlineUserCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnlineUserCountResponse) SetStatusCode(v int32) *DescribeOnlineUserCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnlineUserCountResponse) SetBody(v *DescribeOnlineUserCountResponseBody) *DescribeOnlineUserCountResponse {
	s.Body = v
	return s
}

func (s *DescribeOnlineUserCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
