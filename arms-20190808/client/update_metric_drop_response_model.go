// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetricDropResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetricDropResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetricDropResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetricDropResponseBody) *UpdateMetricDropResponse
	GetBody() *UpdateMetricDropResponseBody
}

type UpdateMetricDropResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetricDropResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetricDropResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetricDropResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetricDropResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetricDropResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetricDropResponse) GetBody() *UpdateMetricDropResponseBody {
	return s.Body
}

func (s *UpdateMetricDropResponse) SetHeaders(v map[string]*string) *UpdateMetricDropResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetricDropResponse) SetStatusCode(v int32) *UpdateMetricDropResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetricDropResponse) SetBody(v *UpdateMetricDropResponseBody) *UpdateMetricDropResponse {
	s.Body = v
	return s
}

func (s *UpdateMetricDropResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
