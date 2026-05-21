// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *PrepareUpgradeRequest
	GetRegionId() *string
	SetDryRun(v bool) *PrepareUpgradeRequest
	GetDryRun() *bool
	SetVersion(v string) *PrepareUpgradeRequest
	GetVersion() *string
}

type PrepareUpgradeRequest struct {
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// r5.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PrepareUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s PrepareUpgradeRequest) GoString() string {
	return s.String()
}

func (s *PrepareUpgradeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PrepareUpgradeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *PrepareUpgradeRequest) GetVersion() *string {
	return s.Version
}

func (s *PrepareUpgradeRequest) SetRegionId(v string) *PrepareUpgradeRequest {
	s.RegionId = &v
	return s
}

func (s *PrepareUpgradeRequest) SetDryRun(v bool) *PrepareUpgradeRequest {
	s.DryRun = &v
	return s
}

func (s *PrepareUpgradeRequest) SetVersion(v string) *PrepareUpgradeRequest {
	s.Version = &v
	return s
}

func (s *PrepareUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
