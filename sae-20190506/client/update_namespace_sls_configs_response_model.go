// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceSlsConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNamespaceSlsConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNamespaceSlsConfigsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNamespaceSlsConfigsResponseBody) *UpdateNamespaceSlsConfigsResponse
	GetBody() *UpdateNamespaceSlsConfigsResponseBody
}

type UpdateNamespaceSlsConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNamespaceSlsConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNamespaceSlsConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceSlsConfigsResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceSlsConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNamespaceSlsConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNamespaceSlsConfigsResponse) GetBody() *UpdateNamespaceSlsConfigsResponseBody {
	return s.Body
}

func (s *UpdateNamespaceSlsConfigsResponse) SetHeaders(v map[string]*string) *UpdateNamespaceSlsConfigsResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponse) SetStatusCode(v int32) *UpdateNamespaceSlsConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponse) SetBody(v *UpdateNamespaceSlsConfigsResponseBody) *UpdateNamespaceSlsConfigsResponse {
	s.Body = v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
