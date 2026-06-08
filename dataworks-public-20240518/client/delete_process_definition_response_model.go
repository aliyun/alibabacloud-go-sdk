// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProcessDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProcessDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProcessDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProcessDefinitionResponseBody) *DeleteProcessDefinitionResponse
	GetBody() *DeleteProcessDefinitionResponseBody
}

type DeleteProcessDefinitionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProcessDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProcessDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteProcessDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProcessDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProcessDefinitionResponse) GetBody() *DeleteProcessDefinitionResponseBody {
	return s.Body
}

func (s *DeleteProcessDefinitionResponse) SetHeaders(v map[string]*string) *DeleteProcessDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteProcessDefinitionResponse) SetStatusCode(v int32) *DeleteProcessDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProcessDefinitionResponse) SetBody(v *DeleteProcessDefinitionResponseBody) *DeleteProcessDefinitionResponse {
	s.Body = v
	return s
}

func (s *DeleteProcessDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
