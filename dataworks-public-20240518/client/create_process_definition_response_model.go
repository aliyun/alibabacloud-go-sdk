// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProcessDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProcessDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *CreateProcessDefinitionResponseBody) *CreateProcessDefinitionResponse
	GetBody() *CreateProcessDefinitionResponseBody
}

type CreateProcessDefinitionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProcessDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProcessDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProcessDefinitionResponse) GetBody() *CreateProcessDefinitionResponseBody {
	return s.Body
}

func (s *CreateProcessDefinitionResponse) SetHeaders(v map[string]*string) *CreateProcessDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateProcessDefinitionResponse) SetStatusCode(v int32) *CreateProcessDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProcessDefinitionResponse) SetBody(v *CreateProcessDefinitionResponseBody) *CreateProcessDefinitionResponse {
	s.Body = v
	return s
}

func (s *CreateProcessDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
