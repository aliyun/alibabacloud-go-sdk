// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCostCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCostCenterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCostCenterResponseBody) *DeleteCostCenterResponse
	GetBody() *DeleteCostCenterResponseBody
}

type DeleteCostCenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCostCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCostCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCostCenterResponse) GetBody() *DeleteCostCenterResponseBody {
	return s.Body
}

func (s *DeleteCostCenterResponse) SetHeaders(v map[string]*string) *DeleteCostCenterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCostCenterResponse) SetStatusCode(v int32) *DeleteCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCostCenterResponse) SetBody(v *DeleteCostCenterResponseBody) *DeleteCostCenterResponse {
	s.Body = v
	return s
}

func (s *DeleteCostCenterResponse) Validate() error {
	return dara.Validate(s)
}
