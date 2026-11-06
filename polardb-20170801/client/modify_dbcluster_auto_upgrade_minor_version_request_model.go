// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAutoUpgradeMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpgradeMinorVersion(v string) *ModifyDBClusterAutoUpgradeMinorVersionRequest
	GetAutoUpgradeMinorVersion() *string
	SetDBClusterId(v string) *ModifyDBClusterAutoUpgradeMinorVersionRequest
	GetDBClusterId() *string
}

type ModifyDBClusterAutoUpgradeMinorVersionRequest struct {
	// The minor engine version update policy. Valid values: Auto: automatic update. Manual: manual update.
	//
	// This parameter is required.
	//
	// example:
	//
	// Manual
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ModifyDBClusterAutoUpgradeMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAutoUpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionRequest) GetAutoUpgradeMinorVersion() *string {
	return s.AutoUpgradeMinorVersion
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionRequest) SetAutoUpgradeMinorVersion(v string) *ModifyDBClusterAutoUpgradeMinorVersionRequest {
	s.AutoUpgradeMinorVersion = &v
	return s
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionRequest) SetDBClusterId(v string) *ModifyDBClusterAutoUpgradeMinorVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAutoUpgradeMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
