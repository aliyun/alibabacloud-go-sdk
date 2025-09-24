// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCostUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCostUnitResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCostUnitResponseBody) *DeleteCostUnitResponse
	GetBody() *DeleteCostUnitResponseBody
}

type DeleteCostUnitResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCostUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCostUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCostUnitResponse) GetBody() *DeleteCostUnitResponseBody {
	return s.Body
}

func (s *DeleteCostUnitResponse) SetHeaders(v map[string]*string) *DeleteCostUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteCostUnitResponse) SetStatusCode(v int32) *DeleteCostUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCostUnitResponse) SetBody(v *DeleteCostUnitResponseBody) *DeleteCostUnitResponse {
	s.Body = v
	return s
}

func (s *DeleteCostUnitResponse) Validate() error {
	return dara.Validate(s)
}
