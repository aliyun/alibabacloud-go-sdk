// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamVodListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamVodListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamVodListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamVodListResponseBody) *DescribeStreamVodListResponse
	GetBody() *DescribeStreamVodListResponseBody
}

type DescribeStreamVodListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamVodListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamVodListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamVodListResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamVodListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamVodListResponse) GetBody() *DescribeStreamVodListResponseBody {
	return s.Body
}

func (s *DescribeStreamVodListResponse) SetHeaders(v map[string]*string) *DescribeStreamVodListResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamVodListResponse) SetStatusCode(v int32) *DescribeStreamVodListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamVodListResponse) SetBody(v *DescribeStreamVodListResponseBody) *DescribeStreamVodListResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamVodListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
