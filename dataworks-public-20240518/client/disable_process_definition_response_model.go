// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableProcessDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableProcessDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableProcessDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *DisableProcessDefinitionResponseBody) *DisableProcessDefinitionResponse
	GetBody() *DisableProcessDefinitionResponseBody
}

type DisableProcessDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableProcessDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableProcessDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DisableProcessDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableProcessDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableProcessDefinitionResponse) GetBody() *DisableProcessDefinitionResponseBody {
	return s.Body
}

func (s *DisableProcessDefinitionResponse) SetHeaders(v map[string]*string) *DisableProcessDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DisableProcessDefinitionResponse) SetStatusCode(v int32) *DisableProcessDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableProcessDefinitionResponse) SetBody(v *DisableProcessDefinitionResponseBody) *DisableProcessDefinitionResponse {
	s.Body = v
	return s
}

func (s *DisableProcessDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
