// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProvisionedProductPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProvisionedProductPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProvisionedProductPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProvisionedProductPlanResponseBody) *UpdateProvisionedProductPlanResponse
	GetBody() *UpdateProvisionedProductPlanResponseBody
}

type UpdateProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProvisionedProductPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProvisionedProductPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProvisionedProductPlanResponse) GetBody() *UpdateProvisionedProductPlanResponseBody {
	return s.Body
}

func (s *UpdateProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *UpdateProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateProvisionedProductPlanResponse) SetStatusCode(v int32) *UpdateProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponse) SetBody(v *UpdateProvisionedProductPlanResponseBody) *UpdateProvisionedProductPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateProvisionedProductPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
