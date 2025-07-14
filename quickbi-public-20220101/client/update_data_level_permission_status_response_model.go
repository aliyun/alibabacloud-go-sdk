// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLevelPermissionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataLevelPermissionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataLevelPermissionStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataLevelPermissionStatusResponseBody) *UpdateDataLevelPermissionStatusResponse
	GetBody() *UpdateDataLevelPermissionStatusResponseBody
}

type UpdateDataLevelPermissionStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataLevelPermissionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataLevelPermissionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLevelPermissionStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataLevelPermissionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataLevelPermissionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataLevelPermissionStatusResponse) GetBody() *UpdateDataLevelPermissionStatusResponseBody {
	return s.Body
}

func (s *UpdateDataLevelPermissionStatusResponse) SetHeaders(v map[string]*string) *UpdateDataLevelPermissionStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponse) SetStatusCode(v int32) *UpdateDataLevelPermissionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponse) SetBody(v *UpdateDataLevelPermissionStatusResponseBody) *UpdateDataLevelPermissionStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponse) Validate() error {
	return dara.Validate(s)
}
