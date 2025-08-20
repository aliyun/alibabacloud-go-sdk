// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNamespaceResponseBody) *UpdateNamespaceResponse
	GetBody() *UpdateNamespaceResponseBody
}

type UpdateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNamespaceResponse) GetBody() *UpdateNamespaceResponseBody {
	return s.Body
}

func (s *UpdateNamespaceResponse) SetHeaders(v map[string]*string) *UpdateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceResponse) SetStatusCode(v int32) *UpdateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceResponse) SetBody(v *UpdateNamespaceResponseBody) *UpdateNamespaceResponse {
	s.Body = v
	return s
}

func (s *UpdateNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
