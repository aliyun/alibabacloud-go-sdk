// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupInsightEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LookupInsightEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LookupInsightEventsResponse
	GetStatusCode() *int32
	SetBody(v *LookupInsightEventsResponseBody) *LookupInsightEventsResponse
	GetBody() *LookupInsightEventsResponseBody
}

type LookupInsightEventsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LookupInsightEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LookupInsightEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s LookupInsightEventsResponse) GoString() string {
	return s.String()
}

func (s *LookupInsightEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LookupInsightEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LookupInsightEventsResponse) GetBody() *LookupInsightEventsResponseBody {
	return s.Body
}

func (s *LookupInsightEventsResponse) SetHeaders(v map[string]*string) *LookupInsightEventsResponse {
	s.Headers = v
	return s
}

func (s *LookupInsightEventsResponse) SetStatusCode(v int32) *LookupInsightEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupInsightEventsResponse) SetBody(v *LookupInsightEventsResponseBody) *LookupInsightEventsResponse {
	s.Body = v
	return s
}

func (s *LookupInsightEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
