// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelOverallDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelOverallDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelOverallDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelOverallDataResponseBody) *DescribeChannelOverallDataResponse
	GetBody() *DescribeChannelOverallDataResponseBody
}

type DescribeChannelOverallDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelOverallDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelOverallDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelOverallDataResponse) GetBody() *DescribeChannelOverallDataResponseBody {
	return s.Body
}

func (s *DescribeChannelOverallDataResponse) SetHeaders(v map[string]*string) *DescribeChannelOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelOverallDataResponse) SetStatusCode(v int32) *DescribeChannelOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelOverallDataResponse) SetBody(v *DescribeChannelOverallDataResponseBody) *DescribeChannelOverallDataResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelOverallDataResponse) Validate() error {
	return dara.Validate(s)
}
