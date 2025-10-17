// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionWithScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProcessDefinitionWithScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProcessDefinitionWithScheduleResponse
	GetStatusCode() *int32
	SetBody(v *CreateProcessDefinitionWithScheduleResponseBody) *CreateProcessDefinitionWithScheduleResponse
	GetBody() *CreateProcessDefinitionWithScheduleResponseBody
}

type CreateProcessDefinitionWithScheduleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProcessDefinitionWithScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProcessDefinitionWithScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProcessDefinitionWithScheduleResponse) GetBody() *CreateProcessDefinitionWithScheduleResponseBody {
	return s.Body
}

func (s *CreateProcessDefinitionWithScheduleResponse) SetHeaders(v map[string]*string) *CreateProcessDefinitionWithScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponse) SetStatusCode(v int32) *CreateProcessDefinitionWithScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponse) SetBody(v *CreateProcessDefinitionWithScheduleResponseBody) *CreateProcessDefinitionWithScheduleResponse {
	s.Body = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
