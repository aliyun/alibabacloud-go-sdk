// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CostCenterModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CostCenterModifyResponse
	GetStatusCode() *int32
	SetBody(v *CostCenterModifyResponseBody) *CostCenterModifyResponse
	GetBody() *CostCenterModifyResponseBody
}

type CostCenterModifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CostCenterModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CostCenterModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s CostCenterModifyResponse) GoString() string {
	return s.String()
}

func (s *CostCenterModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CostCenterModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CostCenterModifyResponse) GetBody() *CostCenterModifyResponseBody {
	return s.Body
}

func (s *CostCenterModifyResponse) SetHeaders(v map[string]*string) *CostCenterModifyResponse {
	s.Headers = v
	return s
}

func (s *CostCenterModifyResponse) SetStatusCode(v int32) *CostCenterModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterModifyResponse) SetBody(v *CostCenterModifyResponseBody) *CostCenterModifyResponse {
	s.Body = v
	return s
}

func (s *CostCenterModifyResponse) Validate() error {
	return dara.Validate(s)
}
