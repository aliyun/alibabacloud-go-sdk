// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordIndexFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamRecordIndexFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamRecordIndexFileResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamRecordIndexFileResponseBody) *DescribeLiveStreamRecordIndexFileResponse
	GetBody() *DescribeLiveStreamRecordIndexFileResponseBody
}

type DescribeLiveStreamRecordIndexFileResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamRecordIndexFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamRecordIndexFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamRecordIndexFileResponse) GetBody() *DescribeLiveStreamRecordIndexFileResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamRecordIndexFileResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamRecordIndexFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponse) SetStatusCode(v int32) *DescribeLiveStreamRecordIndexFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponse) SetBody(v *DescribeLiveStreamRecordIndexFileResponseBody) *DescribeLiveStreamRecordIndexFileResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
