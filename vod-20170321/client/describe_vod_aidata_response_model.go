// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodAIDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodAIDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodAIDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodAIDataResponseBody) *DescribeVodAIDataResponse
	GetBody() *DescribeVodAIDataResponseBody
}

type DescribeVodAIDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodAIDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodAIDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodAIDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodAIDataResponse) GetBody() *DescribeVodAIDataResponseBody {
	return s.Body
}

func (s *DescribeVodAIDataResponse) SetHeaders(v map[string]*string) *DescribeVodAIDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodAIDataResponse) SetStatusCode(v int32) *DescribeVodAIDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodAIDataResponse) SetBody(v *DescribeVodAIDataResponseBody) *DescribeVodAIDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodAIDataResponse) Validate() error {
	return dara.Validate(s)
}
