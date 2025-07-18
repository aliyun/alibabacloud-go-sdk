// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstancePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDBInstancePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDBInstancePlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDBInstancePlanResponseBody) *UpdateDBInstancePlanResponse
	GetBody() *UpdateDBInstancePlanResponseBody
}

type UpdateDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDBInstancePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDBInstancePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDBInstancePlanResponse) GetBody() *UpdateDBInstancePlanResponseBody {
	return s.Body
}

func (s *UpdateDBInstancePlanResponse) SetHeaders(v map[string]*string) *UpdateDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstancePlanResponse) SetStatusCode(v int32) *UpdateDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstancePlanResponse) SetBody(v *UpdateDBInstancePlanResponseBody) *UpdateDBInstancePlanResponse {
	s.Body = v
	return s
}

func (s *UpdateDBInstancePlanResponse) Validate() error {
	return dara.Validate(s)
}
