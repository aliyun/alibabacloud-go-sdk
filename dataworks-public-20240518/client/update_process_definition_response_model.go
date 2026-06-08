// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProcessDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProcessDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProcessDefinitionResponseBody) *UpdateProcessDefinitionResponse
	GetBody() *UpdateProcessDefinitionResponseBody
}

type UpdateProcessDefinitionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProcessDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProcessDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProcessDefinitionResponse) GetBody() *UpdateProcessDefinitionResponseBody {
	return s.Body
}

func (s *UpdateProcessDefinitionResponse) SetHeaders(v map[string]*string) *UpdateProcessDefinitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProcessDefinitionResponse) SetStatusCode(v int32) *UpdateProcessDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProcessDefinitionResponse) SetBody(v *UpdateProcessDefinitionResponseBody) *UpdateProcessDefinitionResponse {
	s.Body = v
	return s
}

func (s *UpdateProcessDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
