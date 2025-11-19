// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodEditingUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodEditingUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodEditingUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodEditingUsageDataResponseBody) *DescribeVodEditingUsageDataResponse
	GetBody() *DescribeVodEditingUsageDataResponseBody
}

type DescribeVodEditingUsageDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodEditingUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodEditingUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodEditingUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodEditingUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodEditingUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodEditingUsageDataResponse) GetBody() *DescribeVodEditingUsageDataResponseBody {
	return s.Body
}

func (s *DescribeVodEditingUsageDataResponse) SetHeaders(v map[string]*string) *DescribeVodEditingUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodEditingUsageDataResponse) SetStatusCode(v int32) *DescribeVodEditingUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponse) SetBody(v *DescribeVodEditingUsageDataResponseBody) *DescribeVodEditingUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodEditingUsageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
