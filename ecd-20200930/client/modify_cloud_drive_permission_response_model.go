// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDrivePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudDrivePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudDrivePermissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudDrivePermissionResponseBody) *ModifyCloudDrivePermissionResponse
	GetBody() *ModifyCloudDrivePermissionResponseBody
}

type ModifyCloudDrivePermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudDrivePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudDrivePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDrivePermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudDrivePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudDrivePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudDrivePermissionResponse) GetBody() *ModifyCloudDrivePermissionResponseBody {
	return s.Body
}

func (s *ModifyCloudDrivePermissionResponse) SetHeaders(v map[string]*string) *ModifyCloudDrivePermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudDrivePermissionResponse) SetStatusCode(v int32) *ModifyCloudDrivePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudDrivePermissionResponse) SetBody(v *ModifyCloudDrivePermissionResponseBody) *ModifyCloudDrivePermissionResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudDrivePermissionResponse) Validate() error {
	return dara.Validate(s)
}
