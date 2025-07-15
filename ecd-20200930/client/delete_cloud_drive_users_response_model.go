// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudDriveUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudDriveUsersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudDriveUsersResponseBody) *DeleteCloudDriveUsersResponse
	GetBody() *DeleteCloudDriveUsersResponseBody
}

type DeleteCloudDriveUsersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudDriveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudDriveUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudDriveUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudDriveUsersResponse) GetBody() *DeleteCloudDriveUsersResponseBody {
	return s.Body
}

func (s *DeleteCloudDriveUsersResponse) SetHeaders(v map[string]*string) *DeleteCloudDriveUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudDriveUsersResponse) SetStatusCode(v int32) *DeleteCloudDriveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudDriveUsersResponse) SetBody(v *DeleteCloudDriveUsersResponseBody) *DeleteCloudDriveUsersResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudDriveUsersResponse) Validate() error {
	return dara.Validate(s)
}
