// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceSubscriptionEstimateCostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceInstanceSubscriptionEstimateCostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceInstanceSubscriptionEstimateCostResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceInstanceSubscriptionEstimateCostResponseBody) *GetServiceInstanceSubscriptionEstimateCostResponse
	GetBody() *GetServiceInstanceSubscriptionEstimateCostResponseBody
}

type GetServiceInstanceSubscriptionEstimateCostResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceSubscriptionEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) GetBody() *GetServiceInstanceSubscriptionEstimateCostResponseBody {
	return s.Body
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) SetHeaders(v map[string]*string) *GetServiceInstanceSubscriptionEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) SetStatusCode(v int32) *GetServiceInstanceSubscriptionEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) SetBody(v *GetServiceInstanceSubscriptionEstimateCostResponseBody) *GetServiceInstanceSubscriptionEstimateCostResponse {
	s.Body = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) Validate() error {
	return dara.Validate(s)
}
