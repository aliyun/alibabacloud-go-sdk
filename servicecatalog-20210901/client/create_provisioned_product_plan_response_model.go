// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProvisionedProductPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProvisionedProductPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProvisionedProductPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateProvisionedProductPlanResponseBody) *CreateProvisionedProductPlanResponse
	GetBody() *CreateProvisionedProductPlanResponseBody
}

type CreateProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProvisionedProductPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProvisionedProductPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProvisionedProductPlanResponse) GetBody() *CreateProvisionedProductPlanResponseBody {
	return s.Body
}

func (s *CreateProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *CreateProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateProvisionedProductPlanResponse) SetStatusCode(v int32) *CreateProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProvisionedProductPlanResponse) SetBody(v *CreateProvisionedProductPlanResponseBody) *CreateProvisionedProductPlanResponse {
	s.Body = v
	return s
}

func (s *CreateProvisionedProductPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
