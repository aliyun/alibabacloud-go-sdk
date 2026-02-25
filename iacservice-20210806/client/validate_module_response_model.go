// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateModuleResponse
	GetStatusCode() *int32
	SetId(v string) *ValidateModuleResponse
	GetId() *string
	SetEvent(v string) *ValidateModuleResponse
	GetEvent() *string
	SetBody(v *ValidateModuleResponseBody) *ValidateModuleResponse
	GetBody() *ValidateModuleResponseBody
}

type ValidateModuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                     `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                     `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ValidateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateModuleResponse) GoString() string {
	return s.String()
}

func (s *ValidateModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateModuleResponse) GetId() *string {
	return s.Id
}

func (s *ValidateModuleResponse) GetEvent() *string {
	return s.Event
}

func (s *ValidateModuleResponse) GetBody() *ValidateModuleResponseBody {
	return s.Body
}

func (s *ValidateModuleResponse) SetHeaders(v map[string]*string) *ValidateModuleResponse {
	s.Headers = v
	return s
}

func (s *ValidateModuleResponse) SetStatusCode(v int32) *ValidateModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateModuleResponse) SetId(v string) *ValidateModuleResponse {
	s.Id = &v
	return s
}

func (s *ValidateModuleResponse) SetEvent(v string) *ValidateModuleResponse {
	s.Event = &v
	return s
}

func (s *ValidateModuleResponse) SetBody(v *ValidateModuleResponseBody) *ValidateModuleResponse {
	s.Body = v
	return s
}

func (s *ValidateModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
