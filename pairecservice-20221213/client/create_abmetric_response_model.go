// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateABMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateABMetricResponse
	GetStatusCode() *int32
	SetBody(v *CreateABMetricResponseBody) *CreateABMetricResponse
	GetBody() *CreateABMetricResponseBody
}

type CreateABMetricResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateABMetricResponse) GoString() string {
	return s.String()
}

func (s *CreateABMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateABMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateABMetricResponse) GetBody() *CreateABMetricResponseBody {
	return s.Body
}

func (s *CreateABMetricResponse) SetHeaders(v map[string]*string) *CreateABMetricResponse {
	s.Headers = v
	return s
}

func (s *CreateABMetricResponse) SetStatusCode(v int32) *CreateABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABMetricResponse) SetBody(v *CreateABMetricResponseBody) *CreateABMetricResponse {
	s.Body = v
	return s
}

func (s *CreateABMetricResponse) Validate() error {
	return dara.Validate(s)
}
