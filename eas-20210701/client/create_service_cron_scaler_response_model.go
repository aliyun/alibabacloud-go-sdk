// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCronScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceCronScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceCronScalerResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceCronScalerResponseBody) *CreateServiceCronScalerResponse
	GetBody() *CreateServiceCronScalerResponseBody
}

type CreateServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceCronScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceCronScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceCronScalerResponse) GetBody() *CreateServiceCronScalerResponseBody {
	return s.Body
}

func (s *CreateServiceCronScalerResponse) SetHeaders(v map[string]*string) *CreateServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceCronScalerResponse) SetStatusCode(v int32) *CreateServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceCronScalerResponse) SetBody(v *CreateServiceCronScalerResponseBody) *CreateServiceCronScalerResponse {
	s.Body = v
	return s
}

func (s *CreateServiceCronScalerResponse) Validate() error {
	return dara.Validate(s)
}
