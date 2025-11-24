// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPlaneLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateControlPlaneLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateControlPlaneLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateControlPlaneLogConfigResponseBody) *UpdateControlPlaneLogConfigResponse
	GetBody() *UpdateControlPlaneLogConfigResponseBody
}

type UpdateControlPlaneLogConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateControlPlaneLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateControlPlaneLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPlaneLogConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateControlPlaneLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateControlPlaneLogConfigResponse) GetBody() *UpdateControlPlaneLogConfigResponseBody {
	return s.Body
}

func (s *UpdateControlPlaneLogConfigResponse) SetHeaders(v map[string]*string) *UpdateControlPlaneLogConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateControlPlaneLogConfigResponse) SetStatusCode(v int32) *UpdateControlPlaneLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateControlPlaneLogConfigResponse) SetBody(v *UpdateControlPlaneLogConfigResponseBody) *UpdateControlPlaneLogConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateControlPlaneLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
