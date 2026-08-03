// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightsEventsCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInsightsEventsCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInsightsEventsCountResponse
	GetStatusCode() *int32
	SetBody(v *GetInsightsEventsCountResponseBody) *GetInsightsEventsCountResponse
	GetBody() *GetInsightsEventsCountResponseBody
}

type GetInsightsEventsCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInsightsEventsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInsightsEventsCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInsightsEventsCountResponse) GoString() string {
	return s.String()
}

func (s *GetInsightsEventsCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInsightsEventsCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInsightsEventsCountResponse) GetBody() *GetInsightsEventsCountResponseBody {
	return s.Body
}

func (s *GetInsightsEventsCountResponse) SetHeaders(v map[string]*string) *GetInsightsEventsCountResponse {
	s.Headers = v
	return s
}

func (s *GetInsightsEventsCountResponse) SetStatusCode(v int32) *GetInsightsEventsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInsightsEventsCountResponse) SetBody(v *GetInsightsEventsCountResponseBody) *GetInsightsEventsCountResponse {
	s.Body = v
	return s
}

func (s *GetInsightsEventsCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
