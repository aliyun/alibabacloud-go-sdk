// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CostCenterDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CostCenterDeleteResponse
	GetStatusCode() *int32
	SetBody(v *CostCenterDeleteResponseBody) *CostCenterDeleteResponse
	GetBody() *CostCenterDeleteResponseBody
}

type CostCenterDeleteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CostCenterDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CostCenterDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s CostCenterDeleteResponse) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CostCenterDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CostCenterDeleteResponse) GetBody() *CostCenterDeleteResponseBody {
	return s.Body
}

func (s *CostCenterDeleteResponse) SetHeaders(v map[string]*string) *CostCenterDeleteResponse {
	s.Headers = v
	return s
}

func (s *CostCenterDeleteResponse) SetStatusCode(v int32) *CostCenterDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterDeleteResponse) SetBody(v *CostCenterDeleteResponseBody) *CostCenterDeleteResponse {
	s.Body = v
	return s
}

func (s *CostCenterDeleteResponse) Validate() error {
	return dara.Validate(s)
}
