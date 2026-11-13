// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *GetMetricDataResponseBody) *GetMetricDataResponse
	GetBody() *GetMetricDataResponseBody
}

type GetMetricDataResponse struct {
	Headers    map[string]*string         `json:"headers" xml:"headers"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetricDataResponse) GoString() string {
	return s.String()
}

func (s *GetMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetricDataResponse) GetBody() *GetMetricDataResponseBody {
	return s.Body
}

func (s *GetMetricDataResponse) SetHeaders(v map[string]*string) *GetMetricDataResponse {
	s.Headers = v
	return s
}

func (s *GetMetricDataResponse) SetStatusCode(v int32) *GetMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricDataResponse) SetBody(v *GetMetricDataResponseBody) *GetMetricDataResponse {
	s.Body = v
	return s
}

func (s *GetMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
