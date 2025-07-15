// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsBlockListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsBlockListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsBlockListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsBlockListResponseBody) *DescribeLiveStreamsBlockListResponse
	GetBody() *DescribeLiveStreamsBlockListResponseBody
}

type DescribeLiveStreamsBlockListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsBlockListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsBlockListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsBlockListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsBlockListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsBlockListResponse) GetBody() *DescribeLiveStreamsBlockListResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsBlockListResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsBlockListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetStatusCode(v int32) *DescribeLiveStreamsBlockListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetBody(v *DescribeLiveStreamsBlockListResponseBody) *DescribeLiveStreamsBlockListResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) Validate() error {
	return dara.Validate(s)
}
