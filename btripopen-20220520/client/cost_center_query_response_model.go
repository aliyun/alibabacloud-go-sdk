// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CostCenterQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CostCenterQueryResponse
	GetStatusCode() *int32
	SetBody(v *CostCenterQueryResponseBody) *CostCenterQueryResponse
	GetBody() *CostCenterQueryResponseBody
}

type CostCenterQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CostCenterQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CostCenterQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CostCenterQueryResponse) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CostCenterQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CostCenterQueryResponse) GetBody() *CostCenterQueryResponseBody {
	return s.Body
}

func (s *CostCenterQueryResponse) SetHeaders(v map[string]*string) *CostCenterQueryResponse {
	s.Headers = v
	return s
}

func (s *CostCenterQueryResponse) SetStatusCode(v int32) *CostCenterQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CostCenterQueryResponse) SetBody(v *CostCenterQueryResponseBody) *CostCenterQueryResponse {
	s.Body = v
	return s
}

func (s *CostCenterQueryResponse) Validate() error {
	return dara.Validate(s)
}
