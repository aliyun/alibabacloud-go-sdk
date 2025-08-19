// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMemberDisplayNameSyncStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMemberAccountDisplayNameSyncStatus(v string) *SetMemberDisplayNameSyncStatusResponseBody
	GetMemberAccountDisplayNameSyncStatus() *string
	SetRequestId(v string) *SetMemberDisplayNameSyncStatusResponseBody
	GetRequestId() *string
}

type SetMemberDisplayNameSyncStatusResponseBody struct {
	// The status of the Member Display Name Synchronization feature. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	MemberAccountDisplayNameSyncStatus *string `json:"MemberAccountDisplayNameSyncStatus,omitempty" xml:"MemberAccountDisplayNameSyncStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMemberDisplayNameSyncStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetMemberDisplayNameSyncStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetMemberDisplayNameSyncStatusResponseBody) GetMemberAccountDisplayNameSyncStatus() *string {
	return s.MemberAccountDisplayNameSyncStatus
}

func (s *SetMemberDisplayNameSyncStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetMemberDisplayNameSyncStatusResponseBody) SetMemberAccountDisplayNameSyncStatus(v string) *SetMemberDisplayNameSyncStatusResponseBody {
	s.MemberAccountDisplayNameSyncStatus = &v
	return s
}

func (s *SetMemberDisplayNameSyncStatusResponseBody) SetRequestId(v string) *SetMemberDisplayNameSyncStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMemberDisplayNameSyncStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
