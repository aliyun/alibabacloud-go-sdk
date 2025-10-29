// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordVodConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRecordVodConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRecordVodConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRecordVodConfigsResponseBody) *DescribeLiveRecordVodConfigsResponse
	GetBody() *DescribeLiveRecordVodConfigsResponseBody
}

type DescribeLiveRecordVodConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRecordVodConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRecordVodConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRecordVodConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRecordVodConfigsResponse) GetBody() *DescribeLiveRecordVodConfigsResponseBody {
	return s.Body
}

func (s *DescribeLiveRecordVodConfigsResponse) SetHeaders(v map[string]*string) *DescribeLiveRecordVodConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) SetStatusCode(v int32) *DescribeLiveRecordVodConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) SetBody(v *DescribeLiveRecordVodConfigsResponseBody) *DescribeLiveRecordVodConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
