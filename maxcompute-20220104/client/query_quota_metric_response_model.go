// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryQuotaMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryQuotaMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryQuotaMetricResponseBody) *QueryQuotaMetricResponse
	GetBody() *QueryQuotaMetricResponseBody
}

type QueryQuotaMetricResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryQuotaMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryQuotaMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryQuotaMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryQuotaMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryQuotaMetricResponse) GetBody() *QueryQuotaMetricResponseBody {
	return s.Body
}

func (s *QueryQuotaMetricResponse) SetHeaders(v map[string]*string) *QueryQuotaMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryQuotaMetricResponse) SetStatusCode(v int32) *QueryQuotaMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQuotaMetricResponse) SetBody(v *QueryQuotaMetricResponseBody) *QueryQuotaMetricResponse {
	s.Body = v
	return s
}

func (s *QueryQuotaMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
