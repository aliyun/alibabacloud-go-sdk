// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFreeUserEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFreeUserEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFreeUserEventsResponseBody) *DescribeFreeUserEventsResponse
	GetBody() *DescribeFreeUserEventsResponseBody
}

type DescribeFreeUserEventsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFreeUserEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFreeUserEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFreeUserEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFreeUserEventsResponse) GetBody() *DescribeFreeUserEventsResponseBody {
	return s.Body
}

func (s *DescribeFreeUserEventsResponse) SetHeaders(v map[string]*string) *DescribeFreeUserEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFreeUserEventsResponse) SetStatusCode(v int32) *DescribeFreeUserEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFreeUserEventsResponse) SetBody(v *DescribeFreeUserEventsResponseBody) *DescribeFreeUserEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeFreeUserEventsResponse) Validate() error {
	return dara.Validate(s)
}
