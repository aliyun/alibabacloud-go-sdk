// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToggleAutoMinorVersionUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpgrade(v bool) *ToggleAutoMinorVersionUpgradeRequest
	GetAutoUpgrade() *bool
	SetInstanceId(v string) *ToggleAutoMinorVersionUpgradeRequest
	GetInstanceId() *string
}

type ToggleAutoMinorVersionUpgradeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"AutoUpgrade,omitempty" xml:"AutoUpgrade,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ToggleAutoMinorVersionUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s ToggleAutoMinorVersionUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ToggleAutoMinorVersionUpgradeRequest) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *ToggleAutoMinorVersionUpgradeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ToggleAutoMinorVersionUpgradeRequest) SetAutoUpgrade(v bool) *ToggleAutoMinorVersionUpgradeRequest {
	s.AutoUpgrade = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeRequest) SetInstanceId(v string) *ToggleAutoMinorVersionUpgradeRequest {
	s.InstanceId = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
