// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventCountByThreatLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventCountByThreatLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventCountByThreatLevelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventCountByThreatLevelResponseBody) *DescribeEventCountByThreatLevelResponse
	GetBody() *DescribeEventCountByThreatLevelResponseBody
}

type DescribeEventCountByThreatLevelResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventCountByThreatLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventCountByThreatLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventCountByThreatLevelResponse) GetBody() *DescribeEventCountByThreatLevelResponseBody {
	return s.Body
}

func (s *DescribeEventCountByThreatLevelResponse) SetHeaders(v map[string]*string) *DescribeEventCountByThreatLevelResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) SetStatusCode(v int32) *DescribeEventCountByThreatLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) SetBody(v *DescribeEventCountByThreatLevelResponseBody) *DescribeEventCountByThreatLevelResponse {
	s.Body = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
