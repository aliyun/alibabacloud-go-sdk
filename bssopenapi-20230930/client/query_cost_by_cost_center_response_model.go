// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostByCostCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostByCostCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostByCostCenterResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostByCostCenterResponseBody) *QueryCostByCostCenterResponse
	GetBody() *QueryCostByCostCenterResponseBody
}

type QueryCostByCostCenterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostByCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostByCostCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostByCostCenterResponse) GoString() string {
	return s.String()
}

func (s *QueryCostByCostCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostByCostCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostByCostCenterResponse) GetBody() *QueryCostByCostCenterResponseBody {
	return s.Body
}

func (s *QueryCostByCostCenterResponse) SetHeaders(v map[string]*string) *QueryCostByCostCenterResponse {
	s.Headers = v
	return s
}

func (s *QueryCostByCostCenterResponse) SetStatusCode(v int32) *QueryCostByCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostByCostCenterResponse) SetBody(v *QueryCostByCostCenterResponseBody) *QueryCostByCostCenterResponse {
	s.Body = v
	return s
}

func (s *QueryCostByCostCenterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
