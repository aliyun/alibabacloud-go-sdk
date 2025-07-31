// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBizUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBizUnitResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBizUnitResponseBody) *UpdateBizUnitResponse
	GetBody() *UpdateBizUnitResponseBody
}

type UpdateBizUnitResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizUnitResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBizUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBizUnitResponse) GetBody() *UpdateBizUnitResponseBody {
	return s.Body
}

func (s *UpdateBizUnitResponse) SetHeaders(v map[string]*string) *UpdateBizUnitResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizUnitResponse) SetStatusCode(v int32) *UpdateBizUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizUnitResponse) SetBody(v *UpdateBizUnitResponseBody) *UpdateBizUnitResponse {
	s.Body = v
	return s
}

func (s *UpdateBizUnitResponse) Validate() error {
	return dara.Validate(s)
}
