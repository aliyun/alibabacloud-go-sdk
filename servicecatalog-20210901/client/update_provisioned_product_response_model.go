// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProvisionedProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProvisionedProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProvisionedProductResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProvisionedProductResponseBody) *UpdateProvisionedProductResponse
	GetBody() *UpdateProvisionedProductResponseBody
}

type UpdateProvisionedProductResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProvisionedProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProvisionedProductResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProvisionedProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProvisionedProductResponse) GetBody() *UpdateProvisionedProductResponseBody {
	return s.Body
}

func (s *UpdateProvisionedProductResponse) SetHeaders(v map[string]*string) *UpdateProvisionedProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateProvisionedProductResponse) SetStatusCode(v int32) *UpdateProvisionedProductResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProvisionedProductResponse) SetBody(v *UpdateProvisionedProductResponseBody) *UpdateProvisionedProductResponse {
	s.Body = v
	return s
}

func (s *UpdateProvisionedProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
