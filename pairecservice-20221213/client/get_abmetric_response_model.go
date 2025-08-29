// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetABMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetABMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetABMetricResponseBody) *GetABMetricResponse
	GetBody() *GetABMetricResponseBody
}

type GetABMetricResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetABMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetABMetricResponse) GoString() string {
	return s.String()
}

func (s *GetABMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetABMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetABMetricResponse) GetBody() *GetABMetricResponseBody {
	return s.Body
}

func (s *GetABMetricResponse) SetHeaders(v map[string]*string) *GetABMetricResponse {
	s.Headers = v
	return s
}

func (s *GetABMetricResponse) SetStatusCode(v int32) *GetABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetABMetricResponse) SetBody(v *GetABMetricResponseBody) *GetABMetricResponse {
	s.Body = v
	return s
}

func (s *GetABMetricResponse) Validate() error {
	return dara.Validate(s)
}
