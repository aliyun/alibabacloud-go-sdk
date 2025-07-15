// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDriveGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudDriveGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudDriveGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudDriveGroupsResponseBody) *ModifyCloudDriveGroupsResponse
	GetBody() *ModifyCloudDriveGroupsResponseBody
}

type ModifyCloudDriveGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudDriveGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudDriveGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDriveGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudDriveGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudDriveGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudDriveGroupsResponse) GetBody() *ModifyCloudDriveGroupsResponseBody {
	return s.Body
}

func (s *ModifyCloudDriveGroupsResponse) SetHeaders(v map[string]*string) *ModifyCloudDriveGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudDriveGroupsResponse) SetStatusCode(v int32) *ModifyCloudDriveGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudDriveGroupsResponse) SetBody(v *ModifyCloudDriveGroupsResponseBody) *ModifyCloudDriveGroupsResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudDriveGroupsResponse) Validate() error {
	return dara.Validate(s)
}
