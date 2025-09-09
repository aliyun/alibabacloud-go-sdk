// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEstimateCostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceEstimateCostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceEstimateCostResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceEstimateCostResponseBody) *GetServiceEstimateCostResponse
	GetBody() *GetServiceEstimateCostResponseBody
}

type GetServiceEstimateCostResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceEstimateCostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceEstimateCostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceEstimateCostResponse) GetBody() *GetServiceEstimateCostResponseBody {
	return s.Body
}

func (s *GetServiceEstimateCostResponse) SetHeaders(v map[string]*string) *GetServiceEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetServiceEstimateCostResponse) SetStatusCode(v int32) *GetServiceEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceEstimateCostResponse) SetBody(v *GetServiceEstimateCostResponseBody) *GetServiceEstimateCostResponse {
	s.Body = v
	return s
}

func (s *GetServiceEstimateCostResponse) Validate() error {
	return dara.Validate(s)
}
