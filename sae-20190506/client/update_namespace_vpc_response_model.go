// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNamespaceVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNamespaceVpcResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNamespaceVpcResponseBody) *UpdateNamespaceVpcResponse
	GetBody() *UpdateNamespaceVpcResponseBody
}

type UpdateNamespaceVpcResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNamespaceVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNamespaceVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceVpcResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNamespaceVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNamespaceVpcResponse) GetBody() *UpdateNamespaceVpcResponseBody {
	return s.Body
}

func (s *UpdateNamespaceVpcResponse) SetHeaders(v map[string]*string) *UpdateNamespaceVpcResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceVpcResponse) SetStatusCode(v int32) *UpdateNamespaceVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceVpcResponse) SetBody(v *UpdateNamespaceVpcResponseBody) *UpdateNamespaceVpcResponse {
	s.Body = v
	return s
}

func (s *UpdateNamespaceVpcResponse) Validate() error {
	return dara.Validate(s)
}
