// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpgradeInstanceRequest
	GetRegionId() *string
	SetType(v string) *UpgradeInstanceRequest
	GetType() *string
	SetUpgradeTime(v string) *UpgradeInstanceRequest
	GetUpgradeTime() *string
}

type UpgradeInstanceRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hot
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 2025-02-11 10:00:01
	UpgradeTime *string `json:"upgradeTime,omitempty" xml:"upgradeTime,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeInstanceRequest) GetType() *string {
	return s.Type
}

func (s *UpgradeInstanceRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeInstanceRequest) SetRegionId(v string) *UpgradeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetType(v string) *UpgradeInstanceRequest {
	s.Type = &v
	return s
}

func (s *UpgradeInstanceRequest) SetUpgradeTime(v string) *UpgradeInstanceRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
