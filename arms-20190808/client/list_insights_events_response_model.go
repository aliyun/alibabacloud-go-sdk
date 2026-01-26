// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInsightsEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInsightsEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInsightsEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListInsightsEventsResponseBody) *ListInsightsEventsResponse
	GetBody() *ListInsightsEventsResponseBody
}

type ListInsightsEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInsightsEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInsightsEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInsightsEventsResponse) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInsightsEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInsightsEventsResponse) GetBody() *ListInsightsEventsResponseBody {
	return s.Body
}

func (s *ListInsightsEventsResponse) SetHeaders(v map[string]*string) *ListInsightsEventsResponse {
	s.Headers = v
	return s
}

func (s *ListInsightsEventsResponse) SetStatusCode(v int32) *ListInsightsEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInsightsEventsResponse) SetBody(v *ListInsightsEventsResponseBody) *ListInsightsEventsResponse {
	s.Body = v
	return s
}

func (s *ListInsightsEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
