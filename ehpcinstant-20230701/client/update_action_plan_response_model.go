// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateActionPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateActionPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateActionPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateActionPlanResponseBody) *UpdateActionPlanResponse
	GetBody() *UpdateActionPlanResponseBody
}

type UpdateActionPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateActionPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateActionPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateActionPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateActionPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateActionPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateActionPlanResponse) GetBody() *UpdateActionPlanResponseBody {
	return s.Body
}

func (s *UpdateActionPlanResponse) SetHeaders(v map[string]*string) *UpdateActionPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateActionPlanResponse) SetStatusCode(v int32) *UpdateActionPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateActionPlanResponse) SetBody(v *UpdateActionPlanResponseBody) *UpdateActionPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateActionPlanResponse) Validate() error {
	return dara.Validate(s)
}
