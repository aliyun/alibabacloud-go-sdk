// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSEventsResponseBody) *DescribeDDoSEventsResponse
	GetBody() *DescribeDDoSEventsResponseBody
}

type DescribeDDoSEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSEventsResponse) GetBody() *DescribeDDoSEventsResponseBody {
	return s.Body
}

func (s *DescribeDDoSEventsResponse) SetHeaders(v map[string]*string) *DescribeDDoSEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSEventsResponse) SetStatusCode(v int32) *DescribeDDoSEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSEventsResponse) SetBody(v *DescribeDDoSEventsResponseBody) *DescribeDDoSEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
