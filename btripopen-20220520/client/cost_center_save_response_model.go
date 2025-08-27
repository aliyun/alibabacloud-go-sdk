// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CostCenterSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CostCenterSaveResponse
	GetStatusCode() *int32
	SetBody(v *CostCenterSaveResponseBody) *CostCenterSaveResponse
	GetBody() *CostCenterSaveResponseBody
}

type CostCenterSaveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CostCenterSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CostCenterSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s CostCenterSaveResponse) GoString() string {
	return s.String()
}

func (s *CostCenterSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CostCenterSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CostCenterSaveResponse) GetBody() *CostCenterSaveResponseBody {
	return s.Body
}

func (s *CostCenterSaveResponse) SetHeaders(v map[string]*string) *CostCenterSaveResponse {
	s.Headers = v
	return s
}

func (s *CostCenterSaveResponse) SetStatusCode(v int32) *CostCenterSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterSaveResponse) SetBody(v *CostCenterSaveResponseBody) *CostCenterSaveResponse {
	s.Body = v
	return s
}

func (s *CostCenterSaveResponse) Validate() error {
	return dara.Validate(s)
}
