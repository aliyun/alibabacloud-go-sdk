// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArmsTopNMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArmsTopNMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArmsTopNMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetArmsTopNMetricResponseBody) *GetArmsTopNMetricResponse
	GetBody() *GetArmsTopNMetricResponseBody
}

type GetArmsTopNMetricResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArmsTopNMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArmsTopNMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArmsTopNMetricResponse) GoString() string {
	return s.String()
}

func (s *GetArmsTopNMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArmsTopNMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArmsTopNMetricResponse) GetBody() *GetArmsTopNMetricResponseBody {
	return s.Body
}

func (s *GetArmsTopNMetricResponse) SetHeaders(v map[string]*string) *GetArmsTopNMetricResponse {
	s.Headers = v
	return s
}

func (s *GetArmsTopNMetricResponse) SetStatusCode(v int32) *GetArmsTopNMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArmsTopNMetricResponse) SetBody(v *GetArmsTopNMetricResponseBody) *GetArmsTopNMetricResponse {
	s.Body = v
	return s
}

func (s *GetArmsTopNMetricResponse) Validate() error {
	return dara.Validate(s)
}
