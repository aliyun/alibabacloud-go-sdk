// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodStatisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodStatisResponseBody) *DescribeVodStatisResponse
	GetBody() *DescribeVodStatisResponseBody
}

type DescribeVodStatisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStatisResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodStatisResponse) GetBody() *DescribeVodStatisResponseBody {
	return s.Body
}

func (s *DescribeVodStatisResponse) SetHeaders(v map[string]*string) *DescribeVodStatisResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodStatisResponse) SetStatusCode(v int32) *DescribeVodStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodStatisResponse) SetBody(v *DescribeVodStatisResponseBody) *DescribeVodStatisResponse {
	s.Body = v
	return s
}

func (s *DescribeVodStatisResponse) Validate() error {
	return dara.Validate(s)
}
