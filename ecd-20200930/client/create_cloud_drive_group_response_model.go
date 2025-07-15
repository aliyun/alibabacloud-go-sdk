// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudDriveGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudDriveGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudDriveGroupResponseBody) *CreateCloudDriveGroupResponse
	GetBody() *CreateCloudDriveGroupResponseBody
}

type CreateCloudDriveGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudDriveGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudDriveGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudDriveGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudDriveGroupResponse) GetBody() *CreateCloudDriveGroupResponseBody {
	return s.Body
}

func (s *CreateCloudDriveGroupResponse) SetHeaders(v map[string]*string) *CreateCloudDriveGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudDriveGroupResponse) SetStatusCode(v int32) *CreateCloudDriveGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudDriveGroupResponse) SetBody(v *CreateCloudDriveGroupResponseBody) *CreateCloudDriveGroupResponse {
	s.Body = v
	return s
}

func (s *CreateCloudDriveGroupResponse) Validate() error {
	return dara.Validate(s)
}
