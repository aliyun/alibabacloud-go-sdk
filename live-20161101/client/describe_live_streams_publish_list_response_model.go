// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsPublishListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsPublishListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsPublishListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsPublishListResponseBody) *DescribeLiveStreamsPublishListResponse
	GetBody() *DescribeLiveStreamsPublishListResponseBody
}

type DescribeLiveStreamsPublishListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsPublishListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsPublishListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsPublishListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsPublishListResponse) GetBody() *DescribeLiveStreamsPublishListResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsPublishListResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsPublishListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetStatusCode(v int32) *DescribeLiveStreamsPublishListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetBody(v *DescribeLiveStreamsPublishListResponseBody) *DescribeLiveStreamsPublishListResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) Validate() error {
	return dara.Validate(s)
}
