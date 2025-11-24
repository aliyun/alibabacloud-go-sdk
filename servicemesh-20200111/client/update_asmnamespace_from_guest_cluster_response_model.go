// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMNamespaceFromGuestClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateASMNamespaceFromGuestClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateASMNamespaceFromGuestClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateASMNamespaceFromGuestClusterResponseBody) *UpdateASMNamespaceFromGuestClusterResponse
	GetBody() *UpdateASMNamespaceFromGuestClusterResponseBody
}

type UpdateASMNamespaceFromGuestClusterResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateASMNamespaceFromGuestClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateASMNamespaceFromGuestClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMNamespaceFromGuestClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) GetBody() *UpdateASMNamespaceFromGuestClusterResponseBody {
	return s.Body
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) SetHeaders(v map[string]*string) *UpdateASMNamespaceFromGuestClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) SetStatusCode(v int32) *UpdateASMNamespaceFromGuestClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) SetBody(v *UpdateASMNamespaceFromGuestClusterResponseBody) *UpdateASMNamespaceFromGuestClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
