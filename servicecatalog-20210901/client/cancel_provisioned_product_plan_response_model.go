// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelProvisionedProductPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelProvisionedProductPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelProvisionedProductPlanResponse
	GetStatusCode() *int32
	SetBody(v *CancelProvisionedProductPlanResponseBody) *CancelProvisionedProductPlanResponse
	GetBody() *CancelProvisionedProductPlanResponseBody
}

type CancelProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelProvisionedProductPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *CancelProvisionedProductPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelProvisionedProductPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelProvisionedProductPlanResponse) GetBody() *CancelProvisionedProductPlanResponseBody {
	return s.Body
}

func (s *CancelProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *CancelProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *CancelProvisionedProductPlanResponse) SetStatusCode(v int32) *CancelProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelProvisionedProductPlanResponse) SetBody(v *CancelProvisionedProductPlanResponseBody) *CancelProvisionedProductPlanResponse {
	s.Body = v
	return s
}

func (s *CancelProvisionedProductPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
