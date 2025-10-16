// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudDriveUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudDriveUsersResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudDriveUsersResponseBody) *CreateCloudDriveUsersResponse
	GetBody() *CreateCloudDriveUsersResponseBody
}

type CreateCloudDriveUsersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudDriveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudDriveUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveUsersResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudDriveUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudDriveUsersResponse) GetBody() *CreateCloudDriveUsersResponseBody {
	return s.Body
}

func (s *CreateCloudDriveUsersResponse) SetHeaders(v map[string]*string) *CreateCloudDriveUsersResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudDriveUsersResponse) SetStatusCode(v int32) *CreateCloudDriveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudDriveUsersResponse) SetBody(v *CreateCloudDriveUsersResponseBody) *CreateCloudDriveUsersResponse {
	s.Body = v
	return s
}

func (s *CreateCloudDriveUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
