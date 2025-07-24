// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionWithScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProcessDefinitionWithScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProcessDefinitionWithScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProcessDefinitionWithScheduleResponseBody) *UpdateProcessDefinitionWithScheduleResponse
	GetBody() *UpdateProcessDefinitionWithScheduleResponseBody
}

type UpdateProcessDefinitionWithScheduleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProcessDefinitionWithScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProcessDefinitionWithScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProcessDefinitionWithScheduleResponse) GetBody() *UpdateProcessDefinitionWithScheduleResponseBody {
	return s.Body
}

func (s *UpdateProcessDefinitionWithScheduleResponse) SetHeaders(v map[string]*string) *UpdateProcessDefinitionWithScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponse) SetStatusCode(v int32) *UpdateProcessDefinitionWithScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponse) SetBody(v *UpdateProcessDefinitionWithScheduleResponseBody) *UpdateProcessDefinitionWithScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponse) Validate() error {
	return dara.Validate(s)
}
