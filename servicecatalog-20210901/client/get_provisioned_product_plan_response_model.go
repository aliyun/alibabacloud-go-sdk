// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionedProductPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProvisionedProductPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProvisionedProductPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetProvisionedProductPlanResponseBody) *GetProvisionedProductPlanResponse
	GetBody() *GetProvisionedProductPlanResponseBody
}

type GetProvisionedProductPlanResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProvisionedProductPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProvisionedProductPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProvisionedProductPlanResponse) GetBody() *GetProvisionedProductPlanResponseBody {
	return s.Body
}

func (s *GetProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *GetProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionedProductPlanResponse) SetStatusCode(v int32) *GetProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProvisionedProductPlanResponse) SetBody(v *GetProvisionedProductPlanResponseBody) *GetProvisionedProductPlanResponse {
	s.Body = v
	return s
}

func (s *GetProvisionedProductPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
