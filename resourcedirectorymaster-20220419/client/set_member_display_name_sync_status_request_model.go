// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMemberDisplayNameSyncStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *SetMemberDisplayNameSyncStatusRequest
	GetStatus() *string
}

type SetMemberDisplayNameSyncStatusRequest struct {
	// The status of the Member Display Name Synchronization feature. Valid values:
	//
	// 	- Enabled: The feature is enabled. This indicates that the display names specified by the management account for the members will be synchronized to the basic account information of the members. The display name information will be visible and perceptible to the members. If a notification is sent to a member, the display name of the member will be used as the appellation of the member.
	//
	// 	- Disabled: The feature is disabled. This indicates that the display names specified by the management account for the members are valid only in the resource directory and will not be updated to the basic account information of the members.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetMemberDisplayNameSyncStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetMemberDisplayNameSyncStatusRequest) GoString() string {
	return s.String()
}

func (s *SetMemberDisplayNameSyncStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetMemberDisplayNameSyncStatusRequest) SetStatus(v string) *SetMemberDisplayNameSyncStatusRequest {
	s.Status = &v
	return s
}

func (s *SetMemberDisplayNameSyncStatusRequest) Validate() error {
	return dara.Validate(s)
}
