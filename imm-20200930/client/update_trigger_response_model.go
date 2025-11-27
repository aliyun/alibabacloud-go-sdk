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
	SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse
	GetBody() *UpdateTriggerResponseBody
}

type UpdateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateTriggerResponse) GetBody() *UpdateTriggerResponseBody {
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

func (s *UpdateTriggerResponse) SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse {
	s.Body = v
	return s
}

func (s *UpdateTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
