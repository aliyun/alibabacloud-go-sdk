// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEngineNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEngineNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEngineNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEngineNamespaceResponseBody) *UpdateEngineNamespaceResponse
	GetBody() *UpdateEngineNamespaceResponseBody
}

type UpdateEngineNamespaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEngineNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEngineNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEngineNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEngineNamespaceResponse) GetBody() *UpdateEngineNamespaceResponseBody {
	return s.Body
}

func (s *UpdateEngineNamespaceResponse) SetHeaders(v map[string]*string) *UpdateEngineNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateEngineNamespaceResponse) SetStatusCode(v int32) *UpdateEngineNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEngineNamespaceResponse) SetBody(v *UpdateEngineNamespaceResponseBody) *UpdateEngineNamespaceResponse {
	s.Body = v
	return s
}

func (s *UpdateEngineNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
