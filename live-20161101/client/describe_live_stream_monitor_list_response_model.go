// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMonitorListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamMonitorListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamMonitorListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamMonitorListResponseBody) *DescribeLiveStreamMonitorListResponse
	GetBody() *DescribeLiveStreamMonitorListResponseBody
}

type DescribeLiveStreamMonitorListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamMonitorListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamMonitorListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamMonitorListResponse) GetBody() *DescribeLiveStreamMonitorListResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamMonitorListResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamMonitorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponse) SetStatusCode(v int32) *DescribeLiveStreamMonitorListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponse) SetBody(v *DescribeLiveStreamMonitorListResponseBody) *DescribeLiveStreamMonitorListResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
