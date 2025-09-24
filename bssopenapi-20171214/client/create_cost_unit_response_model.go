// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCostUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCostUnitResponse
	GetStatusCode() *int32
	SetBody(v *CreateCostUnitResponseBody) *CreateCostUnitResponse
	GetBody() *CreateCostUnitResponseBody
}

type CreateCostUnitResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCostUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCostUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCostUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCostUnitResponse) GetBody() *CreateCostUnitResponseBody {
	return s.Body
}

func (s *CreateCostUnitResponse) SetHeaders(v map[string]*string) *CreateCostUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateCostUnitResponse) SetStatusCode(v int32) *CreateCostUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCostUnitResponse) SetBody(v *CreateCostUnitResponseBody) *CreateCostUnitResponse {
	s.Body = v
	return s
}

func (s *CreateCostUnitResponse) Validate() error {
	return dara.Validate(s)
}
