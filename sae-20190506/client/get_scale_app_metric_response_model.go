// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScaleAppMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScaleAppMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScaleAppMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetScaleAppMetricResponseBody) *GetScaleAppMetricResponse
	GetBody() *GetScaleAppMetricResponseBody
}

type GetScaleAppMetricResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScaleAppMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScaleAppMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScaleAppMetricResponse) GoString() string {
	return s.String()
}

func (s *GetScaleAppMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScaleAppMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScaleAppMetricResponse) GetBody() *GetScaleAppMetricResponseBody {
	return s.Body
}

func (s *GetScaleAppMetricResponse) SetHeaders(v map[string]*string) *GetScaleAppMetricResponse {
	s.Headers = v
	return s
}

func (s *GetScaleAppMetricResponse) SetStatusCode(v int32) *GetScaleAppMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScaleAppMetricResponse) SetBody(v *GetScaleAppMetricResponseBody) *GetScaleAppMetricResponse {
	s.Body = v
	return s
}

func (s *GetScaleAppMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
