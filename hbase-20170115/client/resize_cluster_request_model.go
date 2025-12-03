// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResizeClusterRequest
	GetClientToken() *string
	SetCloudType(v string) *ResizeClusterRequest
	GetCloudType() *string
	SetClusterId(v string) *ResizeClusterRequest
	GetClusterId() *string
	SetColdStorageSize(v string) *ResizeClusterRequest
	GetColdStorageSize() *string
	SetCoreDiskQuantity(v string) *ResizeClusterRequest
	GetCoreDiskQuantity() *string
	SetCoreDiskSize(v string) *ResizeClusterRequest
	GetCoreDiskSize() *string
	SetCoreDiskType(v string) *ResizeClusterRequest
	GetCoreDiskType() *string
	SetCoreInstanceQuantity(v string) *ResizeClusterRequest
	GetCoreInstanceQuantity() *string
	SetCoreInstanceType(v string) *ResizeClusterRequest
	GetCoreInstanceType() *string
	SetEngine(v string) *ResizeClusterRequest
	GetEngine() *string
	SetIsColdStorage(v string) *ResizeClusterRequest
	GetIsColdStorage() *string
	SetPayType(v string) *ResizeClusterRequest
	GetPayType() *string
	SetRegionId(v string) *ResizeClusterRequest
	GetRegionId() *string
	SetUpgradeType(v string) *ResizeClusterRequest
	GetUpgradeType() *string
	SetZoneId(v string) *ResizeClusterRequest
	GetZoneId() *string
}

type ResizeClusterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CloudType *string `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	// This parameter is required.
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ColdStorageSize  *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	CoreDiskQuantity *string `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize     *string `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType     *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// This parameter is required.
	CoreInstanceQuantity *string `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// This parameter is required.
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsColdStorage *string `json:"IsColdStorage,omitempty" xml:"IsColdStorage,omitempty"`
	PayType       *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ResizeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeClusterRequest) GoString() string {
	return s.String()
}

func (s *ResizeClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResizeClusterRequest) GetCloudType() *string {
	return s.CloudType
}

func (s *ResizeClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResizeClusterRequest) GetColdStorageSize() *string {
	return s.ColdStorageSize
}

func (s *ResizeClusterRequest) GetCoreDiskQuantity() *string {
	return s.CoreDiskQuantity
}

func (s *ResizeClusterRequest) GetCoreDiskSize() *string {
	return s.CoreDiskSize
}

func (s *ResizeClusterRequest) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *ResizeClusterRequest) GetCoreInstanceQuantity() *string {
	return s.CoreInstanceQuantity
}

func (s *ResizeClusterRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *ResizeClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *ResizeClusterRequest) GetIsColdStorage() *string {
	return s.IsColdStorage
}

func (s *ResizeClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *ResizeClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResizeClusterRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *ResizeClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ResizeClusterRequest) SetClientToken(v string) *ResizeClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *ResizeClusterRequest) SetCloudType(v string) *ResizeClusterRequest {
	s.CloudType = &v
	return s
}

func (s *ResizeClusterRequest) SetClusterId(v string) *ResizeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeClusterRequest) SetColdStorageSize(v string) *ResizeClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreDiskQuantity(v string) *ResizeClusterRequest {
	s.CoreDiskQuantity = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreDiskSize(v string) *ResizeClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreDiskType(v string) *ResizeClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreInstanceQuantity(v string) *ResizeClusterRequest {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreInstanceType(v string) *ResizeClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ResizeClusterRequest) SetEngine(v string) *ResizeClusterRequest {
	s.Engine = &v
	return s
}

func (s *ResizeClusterRequest) SetIsColdStorage(v string) *ResizeClusterRequest {
	s.IsColdStorage = &v
	return s
}

func (s *ResizeClusterRequest) SetPayType(v string) *ResizeClusterRequest {
	s.PayType = &v
	return s
}

func (s *ResizeClusterRequest) SetRegionId(v string) *ResizeClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeClusterRequest) SetUpgradeType(v string) *ResizeClusterRequest {
	s.UpgradeType = &v
	return s
}

func (s *ResizeClusterRequest) SetZoneId(v string) *ResizeClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *ResizeClusterRequest) Validate() error {
	return dara.Validate(s)
}
