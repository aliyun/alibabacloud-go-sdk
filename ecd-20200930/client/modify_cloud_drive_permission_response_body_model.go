// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDrivePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudDrivePermissionResponseBody
	GetRequestId() *string
}

type ModifyCloudDrivePermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9197824D-AD4B-571F-94BB-C2E6D5855AB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudDrivePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDrivePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudDrivePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudDrivePermissionResponseBody) SetRequestId(v string) *ModifyCloudDrivePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudDrivePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
