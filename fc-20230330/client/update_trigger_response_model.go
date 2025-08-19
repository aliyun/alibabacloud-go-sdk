// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTriggerResponse
	GetStatusCode() *int32
	SetBody(v *Trigger) *UpdateTriggerResponse
	GetBody() *Trigger
}

type UpdateTriggerResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Trigger           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTriggerResponse) GetBody() *Trigger {
	return s.Body
}

func (s *UpdateTriggerResponse) SetHeaders(v map[string]*string) *UpdateTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTriggerResponse) SetStatusCode(v int32) *UpdateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTriggerResponse) SetBody(v *Trigger) *UpdateTriggerResponse {
	s.Body = v
	return s
}

func (s *UpdateTriggerResponse) Validate() error {
	return dara.Validate(s)
}
