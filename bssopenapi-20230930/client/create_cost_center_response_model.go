// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCostCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCostCenterResponse
	GetStatusCode() *int32
	SetBody(v *CreateCostCenterResponseBody) *CreateCostCenterResponse
	GetBody() *CreateCostCenterResponseBody
}

type CreateCostCenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCostCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterResponse) GoString() string {
	return s.String()
}

func (s *CreateCostCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCostCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCostCenterResponse) GetBody() *CreateCostCenterResponseBody {
	return s.Body
}

func (s *CreateCostCenterResponse) SetHeaders(v map[string]*string) *CreateCostCenterResponse {
	s.Headers = v
	return s
}

func (s *CreateCostCenterResponse) SetStatusCode(v int32) *CreateCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCostCenterResponse) SetBody(v *CreateCostCenterResponseBody) *CreateCostCenterResponse {
	s.Body = v
	return s
}

func (s *CreateCostCenterResponse) Validate() error {
	return dara.Validate(s)
}
