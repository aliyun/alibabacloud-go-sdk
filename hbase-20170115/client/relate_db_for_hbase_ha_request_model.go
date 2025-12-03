// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelateDbForHBaseHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RelateDbForHBaseHaRequest
	GetClusterId() *string
	SetHaActive(v string) *RelateDbForHBaseHaRequest
	GetHaActive() *string
	SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest
	GetHaActiveClusterKey() *string
	SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest
	GetHaActiveDBType() *string
	SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest
	GetHaActiveHbaseFsDir() *string
	SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest
	GetHaActiveHdfsUri() *string
	SetHaActivePassword(v string) *RelateDbForHBaseHaRequest
	GetHaActivePassword() *string
	SetHaActiveUser(v string) *RelateDbForHBaseHaRequest
	GetHaActiveUser() *string
	SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest
	GetHaActiveVersion() *string
	SetHaMigrateType(v string) *RelateDbForHBaseHaRequest
	GetHaMigrateType() *string
	SetHaStandby(v string) *RelateDbForHBaseHaRequest
	GetHaStandby() *string
	SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyClusterKey() *string
	SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyDBType() *string
	SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyHbaseFsDir() *string
	SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyHdfsUri() *string
	SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyPassword() *string
	SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyUser() *string
	SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyVersion() *string
	SetHaTables(v string) *RelateDbForHBaseHaRequest
	GetHaTables() *string
	SetIsActiveStandard(v bool) *RelateDbForHBaseHaRequest
	GetIsActiveStandard() *bool
	SetIsStandbyStandard(v bool) *RelateDbForHBaseHaRequest
	GetIsStandbyStandard() *bool
	SetOwnerId(v int64) *RelateDbForHBaseHaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RelateDbForHBaseHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RelateDbForHBaseHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RelateDbForHBaseHaRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *RelateDbForHBaseHaRequest
	GetZoneId() *string
}

type RelateDbForHBaseHaRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	HaActive           *string `json:"HaActive,omitempty" xml:"HaActive,omitempty"`
	HaActiveClusterKey *string `json:"HaActiveClusterKey,omitempty" xml:"HaActiveClusterKey,omitempty"`
	// This parameter is required.
	HaActiveDBType     *string `json:"HaActiveDBType,omitempty" xml:"HaActiveDBType,omitempty"`
	HaActiveHbaseFsDir *string `json:"HaActiveHbaseFsDir,omitempty" xml:"HaActiveHbaseFsDir,omitempty"`
	HaActiveHdfsUri    *string `json:"HaActiveHdfsUri,omitempty" xml:"HaActiveHdfsUri,omitempty"`
	HaActivePassword   *string `json:"HaActivePassword,omitempty" xml:"HaActivePassword,omitempty"`
	HaActiveUser       *string `json:"HaActiveUser,omitempty" xml:"HaActiveUser,omitempty"`
	HaActiveVersion    *string `json:"HaActiveVersion,omitempty" xml:"HaActiveVersion,omitempty"`
	// This parameter is required.
	HaMigrateType *string `json:"HaMigrateType,omitempty" xml:"HaMigrateType,omitempty"`
	// This parameter is required.
	HaStandby           *string `json:"HaStandby,omitempty" xml:"HaStandby,omitempty"`
	HaStandbyClusterKey *string `json:"HaStandbyClusterKey,omitempty" xml:"HaStandbyClusterKey,omitempty"`
	// This parameter is required.
	HaStandbyDBType     *string `json:"HaStandbyDBType,omitempty" xml:"HaStandbyDBType,omitempty"`
	HaStandbyHbaseFsDir *string `json:"HaStandbyHbaseFsDir,omitempty" xml:"HaStandbyHbaseFsDir,omitempty"`
	HaStandbyHdfsUri    *string `json:"HaStandbyHdfsUri,omitempty" xml:"HaStandbyHdfsUri,omitempty"`
	HaStandbyPassword   *string `json:"HaStandbyPassword,omitempty" xml:"HaStandbyPassword,omitempty"`
	HaStandbyUser       *string `json:"HaStandbyUser,omitempty" xml:"HaStandbyUser,omitempty"`
	HaStandbyVersion    *string `json:"HaStandbyVersion,omitempty" xml:"HaStandbyVersion,omitempty"`
	HaTables            *string `json:"HaTables,omitempty" xml:"HaTables,omitempty"`
	// This parameter is required.
	IsActiveStandard *bool `json:"IsActiveStandard,omitempty" xml:"IsActiveStandard,omitempty"`
	// This parameter is required.
	IsStandbyStandard    *bool   `json:"IsStandbyStandard,omitempty" xml:"IsStandbyStandard,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RelateDbForHBaseHaRequest) String() string {
	return dara.Prettify(s)
}

func (s RelateDbForHBaseHaRequest) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RelateDbForHBaseHaRequest) GetHaActive() *string {
	return s.HaActive
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveClusterKey() *string {
	return s.HaActiveClusterKey
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveDBType() *string {
	return s.HaActiveDBType
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveHbaseFsDir() *string {
	return s.HaActiveHbaseFsDir
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveHdfsUri() *string {
	return s.HaActiveHdfsUri
}

func (s *RelateDbForHBaseHaRequest) GetHaActivePassword() *string {
	return s.HaActivePassword
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveUser() *string {
	return s.HaActiveUser
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveVersion() *string {
	return s.HaActiveVersion
}

func (s *RelateDbForHBaseHaRequest) GetHaMigrateType() *string {
	return s.HaMigrateType
}

func (s *RelateDbForHBaseHaRequest) GetHaStandby() *string {
	return s.HaStandby
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyClusterKey() *string {
	return s.HaStandbyClusterKey
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyDBType() *string {
	return s.HaStandbyDBType
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyHbaseFsDir() *string {
	return s.HaStandbyHbaseFsDir
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyHdfsUri() *string {
	return s.HaStandbyHdfsUri
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyPassword() *string {
	return s.HaStandbyPassword
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyUser() *string {
	return s.HaStandbyUser
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyVersion() *string {
	return s.HaStandbyVersion
}

func (s *RelateDbForHBaseHaRequest) GetHaTables() *string {
	return s.HaTables
}

func (s *RelateDbForHBaseHaRequest) GetIsActiveStandard() *bool {
	return s.IsActiveStandard
}

func (s *RelateDbForHBaseHaRequest) GetIsStandbyStandard() *bool {
	return s.IsStandbyStandard
}

func (s *RelateDbForHBaseHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RelateDbForHBaseHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RelateDbForHBaseHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RelateDbForHBaseHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RelateDbForHBaseHaRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RelateDbForHBaseHaRequest) SetClusterId(v string) *RelateDbForHBaseHaRequest {
	s.ClusterId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActive(v string) *RelateDbForHBaseHaRequest {
	s.HaActive = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActivePassword(v string) *RelateDbForHBaseHaRequest {
	s.HaActivePassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveUser(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaMigrateType(v string) *RelateDbForHBaseHaRequest {
	s.HaMigrateType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandby(v string) *RelateDbForHBaseHaRequest {
	s.HaStandby = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyPassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaTables(v string) *RelateDbForHBaseHaRequest {
	s.HaTables = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsActiveStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsActiveStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsStandbyStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsStandbyStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetOwnerId(v int64) *RelateDbForHBaseHaRequest {
	s.OwnerId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetRegionId(v string) *RelateDbForHBaseHaRequest {
	s.RegionId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetResourceOwnerAccount(v string) *RelateDbForHBaseHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetResourceOwnerId(v int64) *RelateDbForHBaseHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetZoneId(v string) *RelateDbForHBaseHaRequest {
	s.ZoneId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) Validate() error {
	return dara.Validate(s)
}
