// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductPlanApproversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProvisionedProductPlanApproversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProvisionedProductPlanApproversResponse
	GetStatusCode() *int32
	SetBody(v *ListProvisionedProductPlanApproversResponseBody) *ListProvisionedProductPlanApproversResponse
	GetBody() *ListProvisionedProductPlanApproversResponseBody
}

type ListProvisionedProductPlanApproversResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionedProductPlanApproversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionedProductPlanApproversResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlanApproversResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProvisionedProductPlanApproversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProvisionedProductPlanApproversResponse) GetBody() *ListProvisionedProductPlanApproversResponseBody {
	return s.Body
}

func (s *ListProvisionedProductPlanApproversResponse) SetHeaders(v map[string]*string) *ListProvisionedProductPlanApproversResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionedProductPlanApproversResponse) SetStatusCode(v int32) *ListProvisionedProductPlanApproversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionedProductPlanApproversResponse) SetBody(v *ListProvisionedProductPlanApproversResponseBody) *ListProvisionedProductPlanApproversResponse {
	s.Body = v
	return s
}

func (s *ListProvisionedProductPlanApproversResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
