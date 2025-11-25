// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogEmptyCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebAccessLogEmptyCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebAccessLogEmptyCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebAccessLogEmptyCountResponseBody) *DescribeWebAccessLogEmptyCountResponse
	GetBody() *DescribeWebAccessLogEmptyCountResponseBody
}

type DescribeWebAccessLogEmptyCountResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebAccessLogEmptyCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebAccessLogEmptyCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogEmptyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogEmptyCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebAccessLogEmptyCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebAccessLogEmptyCountResponse) GetBody() *DescribeWebAccessLogEmptyCountResponseBody {
	return s.Body
}

func (s *DescribeWebAccessLogEmptyCountResponse) SetHeaders(v map[string]*string) *DescribeWebAccessLogEmptyCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponse) SetStatusCode(v int32) *DescribeWebAccessLogEmptyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponse) SetBody(v *DescribeWebAccessLogEmptyCountResponseBody) *DescribeWebAccessLogEmptyCountResponse {
	s.Body = v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
