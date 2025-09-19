// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspEventsResponseBody) *DescribeSuspEventsResponse
	GetBody() *DescribeSuspEventsResponseBody
}

type DescribeSuspEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspEventsResponse) GetBody() *DescribeSuspEventsResponseBody {
	return s.Body
}

func (s *DescribeSuspEventsResponse) SetHeaders(v map[string]*string) *DescribeSuspEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventsResponse) SetStatusCode(v int32) *DescribeSuspEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspEventsResponse) SetBody(v *DescribeSuspEventsResponseBody) *DescribeSuspEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspEventsResponse) Validate() error {
	return dara.Validate(s)
}
