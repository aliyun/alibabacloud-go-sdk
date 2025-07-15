// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudDriveUsersResponseBody
	GetRequestId() *string
}

type DeleteCloudDriveUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudDriveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudDriveUsersResponseBody) SetRequestId(v string) *DeleteCloudDriveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudDriveUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
