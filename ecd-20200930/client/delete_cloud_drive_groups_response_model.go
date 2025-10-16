// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudDriveGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudDriveGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudDriveGroupsResponseBody) *DeleteCloudDriveGroupsResponse
	GetBody() *DeleteCloudDriveGroupsResponseBody
}

type DeleteCloudDriveGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudDriveGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudDriveGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudDriveGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudDriveGroupsResponse) GetBody() *DeleteCloudDriveGroupsResponseBody {
	return s.Body
}

func (s *DeleteCloudDriveGroupsResponse) SetHeaders(v map[string]*string) *DeleteCloudDriveGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudDriveGroupsResponse) SetStatusCode(v int32) *DeleteCloudDriveGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudDriveGroupsResponse) SetBody(v *DeleteCloudDriveGroupsResponseBody) *DeleteCloudDriveGroupsResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudDriveGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
