// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWarehouseAutoScaleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableWarehouseAutoScaleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableWarehouseAutoScaleResponse
	GetStatusCode() *int32
	SetBody(v *DisableWarehouseAutoScaleResponseBody) *DisableWarehouseAutoScaleResponse
	GetBody() *DisableWarehouseAutoScaleResponseBody
}

type DisableWarehouseAutoScaleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableWarehouseAutoScaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableWarehouseAutoScaleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableWarehouseAutoScaleResponse) GoString() string {
	return s.String()
}

func (s *DisableWarehouseAutoScaleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableWarehouseAutoScaleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableWarehouseAutoScaleResponse) GetBody() *DisableWarehouseAutoScaleResponseBody {
	return s.Body
}

func (s *DisableWarehouseAutoScaleResponse) SetHeaders(v map[string]*string) *DisableWarehouseAutoScaleResponse {
	s.Headers = v
	return s
}

func (s *DisableWarehouseAutoScaleResponse) SetStatusCode(v int32) *DisableWarehouseAutoScaleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableWarehouseAutoScaleResponse) SetBody(v *DisableWarehouseAutoScaleResponseBody) *DisableWarehouseAutoScaleResponse {
	s.Body = v
	return s
}

func (s *DisableWarehouseAutoScaleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
