// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMemberDeletionPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManagementAccountId(v string) *SetMemberDeletionPermissionResponseBody
	GetManagementAccountId() *string
	SetMemberDeletionStatus(v string) *SetMemberDeletionPermissionResponseBody
	GetMemberDeletionStatus() *string
	SetRequestId(v string) *SetMemberDeletionPermissionResponseBody
	GetRequestId() *string
	SetResourceDirectoryId(v string) *SetMemberDeletionPermissionResponseBody
	GetResourceDirectoryId() *string
}

type SetMemberDeletionPermissionResponseBody struct {
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 151266687691****
	ManagementAccountId *string `json:"ManagementAccountId,omitempty" xml:"ManagementAccountId,omitempty"`
	// The status of the member deletion feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// example:
	//
	// Enabled
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C55A4CAA-9039-1DDF-91CE-FCC134513D29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s SetMemberDeletionPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetMemberDeletionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionResponseBody) GetManagementAccountId() *string {
	return s.ManagementAccountId
}

func (s *SetMemberDeletionPermissionResponseBody) GetMemberDeletionStatus() *string {
	return s.MemberDeletionStatus
}

func (s *SetMemberDeletionPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetMemberDeletionPermissionResponseBody) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *SetMemberDeletionPermissionResponseBody) SetManagementAccountId(v string) *SetMemberDeletionPermissionResponseBody {
	s.ManagementAccountId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetMemberDeletionStatus(v string) *SetMemberDeletionPermissionResponseBody {
	s.MemberDeletionStatus = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetRequestId(v string) *SetMemberDeletionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetResourceDirectoryId(v string) *SetMemberDeletionPermissionResponseBody {
	s.ResourceDirectoryId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
