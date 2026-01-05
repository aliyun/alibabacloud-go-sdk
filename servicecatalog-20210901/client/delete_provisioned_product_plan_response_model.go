// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProvisionedProductPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProvisionedProductPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProvisionedProductPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProvisionedProductPlanResponseBody) *DeleteProvisionedProductPlanResponse
	GetBody() *DeleteProvisionedProductPlanResponseBody
}

type DeleteProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProvisionedProductPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteProvisionedProductPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProvisionedProductPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProvisionedProductPlanResponse) GetBody() *DeleteProvisionedProductPlanResponseBody {
	return s.Body
}

func (s *DeleteProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *DeleteProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteProvisionedProductPlanResponse) SetStatusCode(v int32) *DeleteProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProvisionedProductPlanResponse) SetBody(v *DeleteProvisionedProductPlanResponseBody) *DeleteProvisionedProductPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteProvisionedProductPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
