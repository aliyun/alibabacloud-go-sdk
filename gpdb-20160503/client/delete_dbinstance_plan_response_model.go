// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstancePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstancePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstancePlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstancePlanResponseBody) *DeleteDBInstancePlanResponse
	GetBody() *DeleteDBInstancePlanResponseBody
}

type DeleteDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstancePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstancePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstancePlanResponse) GetBody() *DeleteDBInstancePlanResponseBody {
	return s.Body
}

func (s *DeleteDBInstancePlanResponse) SetHeaders(v map[string]*string) *DeleteDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstancePlanResponse) SetStatusCode(v int32) *DeleteDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstancePlanResponse) SetBody(v *DeleteDBInstancePlanResponseBody) *DeleteDBInstancePlanResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstancePlanResponse) Validate() error {
	return dara.Validate(s)
}
