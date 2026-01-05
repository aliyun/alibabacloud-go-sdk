// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlotMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySlotMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySlotMetricsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySlotMetricsResponseBody) *QuerySlotMetricsResponse
	GetBody() *QuerySlotMetricsResponseBody
}

type QuerySlotMetricsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySlotMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySlotMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySlotMetricsResponse) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySlotMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySlotMetricsResponse) GetBody() *QuerySlotMetricsResponseBody {
	return s.Body
}

func (s *QuerySlotMetricsResponse) SetHeaders(v map[string]*string) *QuerySlotMetricsResponse {
	s.Headers = v
	return s
}

func (s *QuerySlotMetricsResponse) SetStatusCode(v int32) *QuerySlotMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySlotMetricsResponse) SetBody(v *QuerySlotMetricsResponseBody) *QuerySlotMetricsResponse {
	s.Body = v
	return s
}

func (s *QuerySlotMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
