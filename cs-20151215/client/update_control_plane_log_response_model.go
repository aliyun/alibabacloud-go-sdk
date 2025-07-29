// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPlaneLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateControlPlaneLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateControlPlaneLogResponse
	GetStatusCode() *int32
	SetBody(v *UpdateControlPlaneLogResponseBody) *UpdateControlPlaneLogResponse
	GetBody() *UpdateControlPlaneLogResponseBody
}

type UpdateControlPlaneLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateControlPlaneLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateControlPlaneLogResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPlaneLogResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateControlPlaneLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateControlPlaneLogResponse) GetBody() *UpdateControlPlaneLogResponseBody {
	return s.Body
}

func (s *UpdateControlPlaneLogResponse) SetHeaders(v map[string]*string) *UpdateControlPlaneLogResponse {
	s.Headers = v
	return s
}

func (s *UpdateControlPlaneLogResponse) SetStatusCode(v int32) *UpdateControlPlaneLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateControlPlaneLogResponse) SetBody(v *UpdateControlPlaneLogResponseBody) *UpdateControlPlaneLogResponse {
	s.Body = v
	return s
}

func (s *UpdateControlPlaneLogResponse) Validate() error {
	return dara.Validate(s)
}
