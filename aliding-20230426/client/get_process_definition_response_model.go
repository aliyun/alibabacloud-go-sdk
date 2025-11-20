// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProcessDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProcessDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *GetProcessDefinitionResponseBody) *GetProcessDefinitionResponse
	GetBody() *GetProcessDefinitionResponseBody
}

type GetProcessDefinitionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProcessDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProcessDefinitionResponse) GetBody() *GetProcessDefinitionResponseBody {
	return s.Body
}

func (s *GetProcessDefinitionResponse) SetHeaders(v map[string]*string) *GetProcessDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetProcessDefinitionResponse) SetStatusCode(v int32) *GetProcessDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessDefinitionResponse) SetBody(v *GetProcessDefinitionResponseBody) *GetProcessDefinitionResponse {
	s.Body = v
	return s
}

func (s *GetProcessDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
