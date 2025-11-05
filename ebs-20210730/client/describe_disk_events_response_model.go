// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskEventsResponseBody) *DescribeDiskEventsResponse
	GetBody() *DescribeDiskEventsResponseBody
}

type DescribeDiskEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskEventsResponse) GetBody() *DescribeDiskEventsResponseBody {
	return s.Body
}

func (s *DescribeDiskEventsResponse) SetHeaders(v map[string]*string) *DescribeDiskEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskEventsResponse) SetStatusCode(v int32) *DescribeDiskEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskEventsResponse) SetBody(v *DescribeDiskEventsResponseBody) *DescribeDiskEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
