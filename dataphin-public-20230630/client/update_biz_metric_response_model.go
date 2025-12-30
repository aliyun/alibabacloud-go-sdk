// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBizMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBizMetricResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBizMetricResponseBody) *UpdateBizMetricResponse
	GetBody() *UpdateBizMetricResponseBody
}

type UpdateBizMetricResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBizMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBizMetricResponse) GetBody() *UpdateBizMetricResponseBody {
	return s.Body
}

func (s *UpdateBizMetricResponse) SetHeaders(v map[string]*string) *UpdateBizMetricResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizMetricResponse) SetStatusCode(v int32) *UpdateBizMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizMetricResponse) SetBody(v *UpdateBizMetricResponseBody) *UpdateBizMetricResponse {
	s.Body = v
	return s
}

func (s *UpdateBizMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
