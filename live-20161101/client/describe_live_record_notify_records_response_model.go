// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordNotifyRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRecordNotifyRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRecordNotifyRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRecordNotifyRecordsResponseBody) *DescribeLiveRecordNotifyRecordsResponse
	GetBody() *DescribeLiveRecordNotifyRecordsResponseBody
}

type DescribeLiveRecordNotifyRecordsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRecordNotifyRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRecordNotifyRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRecordNotifyRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRecordNotifyRecordsResponse) GetBody() *DescribeLiveRecordNotifyRecordsResponseBody {
	return s.Body
}

func (s *DescribeLiveRecordNotifyRecordsResponse) SetHeaders(v map[string]*string) *DescribeLiveRecordNotifyRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponse) SetStatusCode(v int32) *DescribeLiveRecordNotifyRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponse) SetBody(v *DescribeLiveRecordNotifyRecordsResponseBody) *DescribeLiveRecordNotifyRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
