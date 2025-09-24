// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveDate(v string) *UpgradeResourcePackageRequest
	GetEffectiveDate() *string
	SetInstanceId(v string) *UpgradeResourcePackageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *UpgradeResourcePackageRequest
	GetOwnerId() *int64
	SetSpecification(v string) *UpgradeResourcePackageRequest
	GetSpecification() *string
}

type UpgradeResourcePackageRequest struct {
	// The time when the resource plan takes effect. If you leave this parameter empty, the resource plan immediately takes effect by default.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-02-02T12:00:00Z
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	// The ID of the resource plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSSBAG-cn-0xl*****x002
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The specifications to which you want to upgrade the resource plan.
	//
	// example:
	//
	// 51200
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s UpgradeResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageRequest) GetEffectiveDate() *string {
	return s.EffectiveDate
}

func (s *UpgradeResourcePackageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeResourcePackageRequest) GetSpecification() *string {
	return s.Specification
}

func (s *UpgradeResourcePackageRequest) SetEffectiveDate(v string) *UpgradeResourcePackageRequest {
	s.EffectiveDate = &v
	return s
}

func (s *UpgradeResourcePackageRequest) SetInstanceId(v string) *UpgradeResourcePackageRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeResourcePackageRequest) SetOwnerId(v int64) *UpgradeResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeResourcePackageRequest) SetSpecification(v string) *UpgradeResourcePackageRequest {
	s.Specification = &v
	return s
}

func (s *UpgradeResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
