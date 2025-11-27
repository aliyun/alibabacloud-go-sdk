// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsOnlineListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsStreamsOnlineListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsStreamsOnlineListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsStreamsOnlineListResponseBody) *DescribeVsStreamsOnlineListResponse
	GetBody() *DescribeVsStreamsOnlineListResponseBody
}

type DescribeVsStreamsOnlineListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsStreamsOnlineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsStreamsOnlineListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsStreamsOnlineListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsStreamsOnlineListResponse) GetBody() *DescribeVsStreamsOnlineListResponseBody {
	return s.Body
}

func (s *DescribeVsStreamsOnlineListResponse) SetHeaders(v map[string]*string) *DescribeVsStreamsOnlineListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStreamsOnlineListResponse) SetStatusCode(v int32) *DescribeVsStreamsOnlineListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponse) SetBody(v *DescribeVsStreamsOnlineListResponseBody) *DescribeVsStreamsOnlineListResponse {
	s.Body = v
	return s
}

func (s *DescribeVsStreamsOnlineListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
