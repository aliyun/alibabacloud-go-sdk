// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDriveUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudDriveUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudDriveUsersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudDriveUsersResponseBody) *ModifyCloudDriveUsersResponse
	GetBody() *ModifyCloudDriveUsersResponseBody
}

type ModifyCloudDriveUsersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudDriveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudDriveUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDriveUsersResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudDriveUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudDriveUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudDriveUsersResponse) GetBody() *ModifyCloudDriveUsersResponseBody {
	return s.Body
}

func (s *ModifyCloudDriveUsersResponse) SetHeaders(v map[string]*string) *ModifyCloudDriveUsersResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudDriveUsersResponse) SetStatusCode(v int32) *ModifyCloudDriveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudDriveUsersResponse) SetBody(v *ModifyCloudDriveUsersResponseBody) *ModifyCloudDriveUsersResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudDriveUsersResponse) Validate() error {
	return dara.Validate(s)
}
