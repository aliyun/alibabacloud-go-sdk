// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateEstimateCostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateEstimateCostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateEstimateCostResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateEstimateCostResponseBody) *GetTemplateEstimateCostResponse
	GetBody() *GetTemplateEstimateCostResponseBody
}

type GetTemplateEstimateCostResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateEstimateCostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateEstimateCostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateEstimateCostResponse) GetBody() *GetTemplateEstimateCostResponseBody {
	return s.Body
}

func (s *GetTemplateEstimateCostResponse) SetHeaders(v map[string]*string) *GetTemplateEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateEstimateCostResponse) SetStatusCode(v int32) *GetTemplateEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateEstimateCostResponse) SetBody(v *GetTemplateEstimateCostResponseBody) *GetTemplateEstimateCostResponse {
	s.Body = v
	return s
}

func (s *GetTemplateEstimateCostResponse) Validate() error {
	return dara.Validate(s)
}
