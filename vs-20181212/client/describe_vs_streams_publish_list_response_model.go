// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsPublishListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsStreamsPublishListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsStreamsPublishListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsStreamsPublishListResponseBody) *DescribeVsStreamsPublishListResponse
	GetBody() *DescribeVsStreamsPublishListResponseBody
}

type DescribeVsStreamsPublishListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsStreamsPublishListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsStreamsPublishListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsStreamsPublishListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsStreamsPublishListResponse) GetBody() *DescribeVsStreamsPublishListResponseBody {
	return s.Body
}

func (s *DescribeVsStreamsPublishListResponse) SetHeaders(v map[string]*string) *DescribeVsStreamsPublishListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStreamsPublishListResponse) SetStatusCode(v int32) *DescribeVsStreamsPublishListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponse) SetBody(v *DescribeVsStreamsPublishListResponseBody) *DescribeVsStreamsPublishListResponse {
	s.Body = v
	return s
}

func (s *DescribeVsStreamsPublishListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
