// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LookupEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LookupEventsResponse
	GetStatusCode() *int32
	SetBody(v *LookupEventsResponseBody) *LookupEventsResponse
	GetBody() *LookupEventsResponseBody
}

type LookupEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LookupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LookupEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s LookupEventsResponse) GoString() string {
	return s.String()
}

func (s *LookupEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LookupEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LookupEventsResponse) GetBody() *LookupEventsResponseBody {
	return s.Body
}

func (s *LookupEventsResponse) SetHeaders(v map[string]*string) *LookupEventsResponse {
	s.Headers = v
	return s
}

func (s *LookupEventsResponse) SetStatusCode(v int32) *LookupEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupEventsResponse) SetBody(v *LookupEventsResponseBody) *LookupEventsResponse {
	s.Body = v
	return s
}

func (s *LookupEventsResponse) Validate() error {
	return dara.Validate(s)
}
