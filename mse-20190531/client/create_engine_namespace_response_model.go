// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEngineNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEngineNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEngineNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateEngineNamespaceResponseBody) *CreateEngineNamespaceResponse
	GetBody() *CreateEngineNamespaceResponseBody
}

type CreateEngineNamespaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEngineNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEngineNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEngineNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEngineNamespaceResponse) GetBody() *CreateEngineNamespaceResponseBody {
	return s.Body
}

func (s *CreateEngineNamespaceResponse) SetHeaders(v map[string]*string) *CreateEngineNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateEngineNamespaceResponse) SetStatusCode(v int32) *CreateEngineNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEngineNamespaceResponse) SetBody(v *CreateEngineNamespaceResponseBody) *CreateEngineNamespaceResponse {
	s.Body = v
	return s
}

func (s *CreateEngineNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
