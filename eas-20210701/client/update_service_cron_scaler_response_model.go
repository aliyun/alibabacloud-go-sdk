// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceCronScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceCronScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceCronScalerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceCronScalerResponseBody) *UpdateServiceCronScalerResponse
	GetBody() *UpdateServiceCronScalerResponseBody
}

type UpdateServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceCronScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceCronScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceCronScalerResponse) GetBody() *UpdateServiceCronScalerResponseBody {
	return s.Body
}

func (s *UpdateServiceCronScalerResponse) SetHeaders(v map[string]*string) *UpdateServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceCronScalerResponse) SetStatusCode(v int32) *UpdateServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceCronScalerResponse) SetBody(v *UpdateServiceCronScalerResponseBody) *UpdateServiceCronScalerResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceCronScalerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
