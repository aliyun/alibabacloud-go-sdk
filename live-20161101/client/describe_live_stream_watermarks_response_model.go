// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamWatermarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamWatermarksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamWatermarksResponseBody) *DescribeLiveStreamWatermarksResponse
	GetBody() *DescribeLiveStreamWatermarksResponseBody
}

type DescribeLiveStreamWatermarksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamWatermarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamWatermarksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarksResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamWatermarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamWatermarksResponse) GetBody() *DescribeLiveStreamWatermarksResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamWatermarksResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamWatermarksResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamWatermarksResponse) SetStatusCode(v int32) *DescribeLiveStreamWatermarksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponse) SetBody(v *DescribeLiveStreamWatermarksResponseBody) *DescribeLiveStreamWatermarksResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamWatermarksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
