// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceDdosCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceDdosCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceDdosCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceDdosCountResponseBody) *DescribeRCInstanceDdosCountResponse
	GetBody() *DescribeRCInstanceDdosCountResponseBody
}

type DescribeRCInstanceDdosCountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceDdosCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceDdosCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceDdosCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceDdosCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceDdosCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceDdosCountResponse) GetBody() *DescribeRCInstanceDdosCountResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceDdosCountResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceDdosCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceDdosCountResponse) SetStatusCode(v int32) *DescribeRCInstanceDdosCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceDdosCountResponse) SetBody(v *DescribeRCInstanceDdosCountResponseBody) *DescribeRCInstanceDdosCountResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceDdosCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
