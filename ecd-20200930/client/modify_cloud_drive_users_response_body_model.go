// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDriveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudDriveUsersResponseBody
	GetRequestId() *string
}

type ModifyCloudDriveUsersResponseBody struct {
	// example:
	//
	// D2E005C4-8CA3-5F1D-9917-E75BE3BF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudDriveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDriveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudDriveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudDriveUsersResponseBody) SetRequestId(v string) *ModifyCloudDriveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudDriveUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
