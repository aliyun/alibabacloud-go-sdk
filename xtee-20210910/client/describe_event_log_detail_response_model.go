// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventLogDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventLogDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventLogDetailResponseBody) *DescribeEventLogDetailResponse
	GetBody() *DescribeEventLogDetailResponseBody
}

type DescribeEventLogDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventLogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventLogDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventLogDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventLogDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventLogDetailResponse) GetBody() *DescribeEventLogDetailResponseBody {
	return s.Body
}

func (s *DescribeEventLogDetailResponse) SetHeaders(v map[string]*string) *DescribeEventLogDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventLogDetailResponse) SetStatusCode(v int32) *DescribeEventLogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventLogDetailResponse) SetBody(v *DescribeEventLogDetailResponseBody) *DescribeEventLogDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeEventLogDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
