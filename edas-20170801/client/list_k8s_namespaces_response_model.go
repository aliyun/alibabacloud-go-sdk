// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sNamespacesResponseBody) *ListK8sNamespacesResponse
	GetBody() *ListK8sNamespacesResponseBody
}

type ListK8sNamespacesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListK8sNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sNamespacesResponse) GetBody() *ListK8sNamespacesResponseBody {
	return s.Body
}

func (s *ListK8sNamespacesResponse) SetHeaders(v map[string]*string) *ListK8sNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListK8sNamespacesResponse) SetStatusCode(v int32) *ListK8sNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sNamespacesResponse) SetBody(v *ListK8sNamespacesResponseBody) *ListK8sNamespacesResponse {
	s.Body = v
	return s
}

func (s *ListK8sNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
