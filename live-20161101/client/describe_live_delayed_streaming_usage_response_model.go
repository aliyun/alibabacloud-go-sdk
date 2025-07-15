// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDelayedStreamingUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDelayedStreamingUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDelayedStreamingUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDelayedStreamingUsageResponseBody) *DescribeLiveDelayedStreamingUsageResponse
	GetBody() *DescribeLiveDelayedStreamingUsageResponseBody
}

type DescribeLiveDelayedStreamingUsageResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDelayedStreamingUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDelayedStreamingUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayedStreamingUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayedStreamingUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDelayedStreamingUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDelayedStreamingUsageResponse) GetBody() *DescribeLiveDelayedStreamingUsageResponseBody {
	return s.Body
}

func (s *DescribeLiveDelayedStreamingUsageResponse) SetHeaders(v map[string]*string) *DescribeLiveDelayedStreamingUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponse) SetStatusCode(v int32) *DescribeLiveDelayedStreamingUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponse) SetBody(v *DescribeLiveDelayedStreamingUsageResponseBody) *DescribeLiveDelayedStreamingUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageResponse) Validate() error {
	return dara.Validate(s)
}
