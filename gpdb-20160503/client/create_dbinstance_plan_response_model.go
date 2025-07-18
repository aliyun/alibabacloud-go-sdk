// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstancePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstancePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstancePlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstancePlanResponseBody) *CreateDBInstancePlanResponse
	GetBody() *CreateDBInstancePlanResponseBody
}

type CreateDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstancePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstancePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstancePlanResponse) GetBody() *CreateDBInstancePlanResponseBody {
	return s.Body
}

func (s *CreateDBInstancePlanResponse) SetHeaders(v map[string]*string) *CreateDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstancePlanResponse) SetStatusCode(v int32) *CreateDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstancePlanResponse) SetBody(v *CreateDBInstancePlanResponseBody) *CreateDBInstancePlanResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstancePlanResponse) Validate() error {
	return dara.Validate(s)
}
