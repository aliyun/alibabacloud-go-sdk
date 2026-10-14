// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallHotfixRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCveId(v string) *UninstallHotfixRequest
	GetCveId() *string
	SetInstances(v []*UninstallHotfixRequestInstances) *UninstallHotfixRequest
	GetInstances() []*UninstallHotfixRequestInstances
}

type UninstallHotfixRequest struct {
	// CVE ID
	//
	// This parameter is required.
	//
	// example:
	//
	// CVE-2026-12345
	CveId *string `json:"cveId,omitempty" xml:"cveId,omitempty"`
	// The target instances.
	//
	// This parameter is required.
	Instances []*UninstallHotfixRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UninstallHotfixRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallHotfixRequest) GoString() string {
	return s.String()
}

func (s *UninstallHotfixRequest) GetCveId() *string {
	return s.CveId
}

func (s *UninstallHotfixRequest) GetInstances() []*UninstallHotfixRequestInstances {
	return s.Instances
}

func (s *UninstallHotfixRequest) SetCveId(v string) *UninstallHotfixRequest {
	s.CveId = &v
	return s
}

func (s *UninstallHotfixRequest) SetInstances(v []*UninstallHotfixRequestInstances) *UninstallHotfixRequest {
	s.Instances = v
	return s
}

func (s *UninstallHotfixRequest) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UninstallHotfixRequestInstances struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UninstallHotfixRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UninstallHotfixRequestInstances) GoString() string {
	return s.String()
}

func (s *UninstallHotfixRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UninstallHotfixRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *UninstallHotfixRequestInstances) SetInstanceId(v string) *UninstallHotfixRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *UninstallHotfixRequestInstances) SetRegionId(v string) *UninstallHotfixRequestInstances {
	s.RegionId = &v
	return s
}

func (s *UninstallHotfixRequestInstances) Validate() error {
	return dara.Validate(s)
}
