// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsOnlineListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsOnlineListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsOnlineListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsOnlineListResponseBody) *DescribeLiveStreamsOnlineListResponse
	GetBody() *DescribeLiveStreamsOnlineListResponseBody
}

type DescribeLiveStreamsOnlineListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsOnlineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsOnlineListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsOnlineListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsOnlineListResponse) GetBody() *DescribeLiveStreamsOnlineListResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsOnlineListResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsOnlineListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetStatusCode(v int32) *DescribeLiveStreamsOnlineListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetBody(v *DescribeLiveStreamsOnlineListResponseBody) *DescribeLiveStreamsOnlineListResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
