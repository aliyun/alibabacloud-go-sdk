// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourcePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushResourcePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushResourcePlanResponse
	GetStatusCode() *int32
	SetBody(v *PushResourcePlanResponseBody) *PushResourcePlanResponse
	GetBody() *PushResourcePlanResponseBody
}

type PushResourcePlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushResourcePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushResourcePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s PushResourcePlanResponse) GoString() string {
	return s.String()
}

func (s *PushResourcePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushResourcePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushResourcePlanResponse) GetBody() *PushResourcePlanResponseBody {
	return s.Body
}

func (s *PushResourcePlanResponse) SetHeaders(v map[string]*string) *PushResourcePlanResponse {
	s.Headers = v
	return s
}

func (s *PushResourcePlanResponse) SetStatusCode(v int32) *PushResourcePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResourcePlanResponse) SetBody(v *PushResourcePlanResponseBody) *PushResourcePlanResponse {
	s.Body = v
	return s
}

func (s *PushResourcePlanResponse) Validate() error {
	return dara.Validate(s)
}
