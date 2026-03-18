// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationChartsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryObservationChartsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryObservationChartsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryObservationChartsResponseBody) *ModelRouterQueryObservationChartsResponse
	GetBody() *ModelRouterQueryObservationChartsResponseBody
}

type ModelRouterQueryObservationChartsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryObservationChartsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryObservationChartsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationChartsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationChartsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryObservationChartsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryObservationChartsResponse) GetBody() *ModelRouterQueryObservationChartsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryObservationChartsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryObservationChartsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryObservationChartsResponse) SetStatusCode(v int32) *ModelRouterQueryObservationChartsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationChartsResponse) SetBody(v *ModelRouterQueryObservationChartsResponseBody) *ModelRouterQueryObservationChartsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryObservationChartsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
