// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceScopeSidecarConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNamespaceScopeSidecarConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNamespaceScopeSidecarConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNamespaceScopeSidecarConfigResponseBody) *UpdateNamespaceScopeSidecarConfigResponse
	GetBody() *UpdateNamespaceScopeSidecarConfigResponseBody
}

type UpdateNamespaceScopeSidecarConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNamespaceScopeSidecarConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) GetBody() *UpdateNamespaceScopeSidecarConfigResponseBody {
	return s.Body
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetHeaders(v map[string]*string) *UpdateNamespaceScopeSidecarConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetStatusCode(v int32) *UpdateNamespaceScopeSidecarConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetBody(v *UpdateNamespaceScopeSidecarConfigResponseBody) *UpdateNamespaceScopeSidecarConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
