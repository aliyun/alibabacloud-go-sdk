// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveProvisionedProductPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveProvisionedProductPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveProvisionedProductPlanResponse
	GetStatusCode() *int32
	SetBody(v *ApproveProvisionedProductPlanResponseBody) *ApproveProvisionedProductPlanResponse
	GetBody() *ApproveProvisionedProductPlanResponseBody
}

type ApproveProvisionedProductPlanResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveProvisionedProductPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *ApproveProvisionedProductPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveProvisionedProductPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveProvisionedProductPlanResponse) GetBody() *ApproveProvisionedProductPlanResponseBody {
	return s.Body
}

func (s *ApproveProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *ApproveProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *ApproveProvisionedProductPlanResponse) SetStatusCode(v int32) *ApproveProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveProvisionedProductPlanResponse) SetBody(v *ApproveProvisionedProductPlanResponseBody) *ApproveProvisionedProductPlanResponse {
	s.Body = v
	return s
}

func (s *ApproveProvisionedProductPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
