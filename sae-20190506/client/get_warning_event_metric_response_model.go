// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarningEventMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWarningEventMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWarningEventMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetWarningEventMetricResponseBody) *GetWarningEventMetricResponse
	GetBody() *GetWarningEventMetricResponseBody
}

type GetWarningEventMetricResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWarningEventMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWarningEventMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWarningEventMetricResponse) GoString() string {
	return s.String()
}

func (s *GetWarningEventMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWarningEventMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWarningEventMetricResponse) GetBody() *GetWarningEventMetricResponseBody {
	return s.Body
}

func (s *GetWarningEventMetricResponse) SetHeaders(v map[string]*string) *GetWarningEventMetricResponse {
	s.Headers = v
	return s
}

func (s *GetWarningEventMetricResponse) SetStatusCode(v int32) *GetWarningEventMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWarningEventMetricResponse) SetBody(v *GetWarningEventMetricResponseBody) *GetWarningEventMetricResponse {
	s.Body = v
	return s
}

func (s *GetWarningEventMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
