// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryObservationLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryObservationLogsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryObservationLogsResponseBody) *ModelRouterQueryObservationLogsResponse
	GetBody() *ModelRouterQueryObservationLogsResponseBody
}

type ModelRouterQueryObservationLogsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryObservationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryObservationLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationLogsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryObservationLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryObservationLogsResponse) GetBody() *ModelRouterQueryObservationLogsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryObservationLogsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryObservationLogsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryObservationLogsResponse) SetStatusCode(v int32) *ModelRouterQueryObservationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponse) SetBody(v *ModelRouterQueryObservationLogsResponseBody) *ModelRouterQueryObservationLogsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryObservationLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
