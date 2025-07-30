// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFastMode(v bool) *UpgradeVersionRequest
	GetFastMode() *bool
	SetInstanceId(v string) *UpgradeVersionRequest
	GetInstanceId() *string
	SetMinor(v bool) *UpgradeVersionRequest
	GetMinor() *bool
	SetTargetVersion(v string) *UpgradeVersionRequest
	GetTargetVersion() *string
}

type UpgradeVersionRequest struct {
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the minor version is upgraded. Default value: true. Valid values:
	//
	// 	- true: The minor version is upgraded.
	//
	// 	- false: The major version is upgraded.
	//
	// example:
	//
	// true
	Minor *bool `json:"Minor,omitempty" xml:"Minor,omitempty"`
	// The version to which you want to upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.3.2-1.77-1.6.4
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s UpgradeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeVersionRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *UpgradeVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeVersionRequest) GetMinor() *bool {
	return s.Minor
}

func (s *UpgradeVersionRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UpgradeVersionRequest) SetFastMode(v bool) *UpgradeVersionRequest {
	s.FastMode = &v
	return s
}

func (s *UpgradeVersionRequest) SetInstanceId(v string) *UpgradeVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeVersionRequest) SetMinor(v bool) *UpgradeVersionRequest {
	s.Minor = &v
	return s
}

func (s *UpgradeVersionRequest) SetTargetVersion(v string) *UpgradeVersionRequest {
	s.TargetVersion = &v
	return s
}

func (s *UpgradeVersionRequest) Validate() error {
	return dara.Validate(s)
}
