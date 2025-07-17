// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMetricByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMetricByPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryMetricByPageResponseBody) *QueryMetricByPageResponse
	GetBody() *QueryMetricByPageResponseBody
}

type QueryMetricByPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMetricByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMetricByPageResponse) GetBody() *QueryMetricByPageResponseBody {
	return s.Body
}

func (s *QueryMetricByPageResponse) SetHeaders(v map[string]*string) *QueryMetricByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricByPageResponse) SetStatusCode(v int32) *QueryMetricByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricByPageResponse) SetBody(v *QueryMetricByPageResponseBody) *QueryMetricByPageResponse {
	s.Body = v
	return s
}

func (s *QueryMetricByPageResponse) Validate() error {
	return dara.Validate(s)
}
