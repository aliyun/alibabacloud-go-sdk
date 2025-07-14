// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse
	GetBody() *CreateNamespaceResponseBody
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNamespaceResponse) GetBody() *CreateNamespaceResponseBody {
	return s.Body
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

func (s *CreateNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
