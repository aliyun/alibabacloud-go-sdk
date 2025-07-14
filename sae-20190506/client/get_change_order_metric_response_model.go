// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeOrderMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChangeOrderMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChangeOrderMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetChangeOrderMetricResponseBody) *GetChangeOrderMetricResponse
	GetBody() *GetChangeOrderMetricResponseBody
}

type GetChangeOrderMetricResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChangeOrderMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChangeOrderMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderMetricResponse) GoString() string {
	return s.String()
}

func (s *GetChangeOrderMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChangeOrderMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChangeOrderMetricResponse) GetBody() *GetChangeOrderMetricResponseBody {
	return s.Body
}

func (s *GetChangeOrderMetricResponse) SetHeaders(v map[string]*string) *GetChangeOrderMetricResponse {
	s.Headers = v
	return s
}

func (s *GetChangeOrderMetricResponse) SetStatusCode(v int32) *GetChangeOrderMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChangeOrderMetricResponse) SetBody(v *GetChangeOrderMetricResponseBody) *GetChangeOrderMetricResponse {
	s.Body = v
	return s
}

func (s *GetChangeOrderMetricResponse) Validate() error {
	return dara.Validate(s)
}
