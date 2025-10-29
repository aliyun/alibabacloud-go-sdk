// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsNotifyRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsNotifyRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsNotifyRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsNotifyRecordsResponseBody) *DescribeLiveStreamsNotifyRecordsResponse
	GetBody() *DescribeLiveStreamsNotifyRecordsResponseBody
}

type DescribeLiveStreamsNotifyRecordsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsNotifyRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsNotifyRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) GetBody() *DescribeLiveStreamsNotifyRecordsResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsNotifyRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) SetStatusCode(v int32) *DescribeLiveStreamsNotifyRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) SetBody(v *DescribeLiveStreamsNotifyRecordsResponseBody) *DescribeLiveStreamsNotifyRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
