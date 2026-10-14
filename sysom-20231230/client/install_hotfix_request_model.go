// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallHotfixRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCveId(v string) *InstallHotfixRequest
	GetCveId() *string
	SetInstances(v []*InstallHotfixRequestInstances) *InstallHotfixRequest
	GetInstances() []*InstallHotfixRequestInstances
}

type InstallHotfixRequest struct {
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
	Instances []*InstallHotfixRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s InstallHotfixRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallHotfixRequest) GoString() string {
	return s.String()
}

func (s *InstallHotfixRequest) GetCveId() *string {
	return s.CveId
}

func (s *InstallHotfixRequest) GetInstances() []*InstallHotfixRequestInstances {
	return s.Instances
}

func (s *InstallHotfixRequest) SetCveId(v string) *InstallHotfixRequest {
	s.CveId = &v
	return s
}

func (s *InstallHotfixRequest) SetInstances(v []*InstallHotfixRequestInstances) *InstallHotfixRequest {
	s.Instances = v
	return s
}

func (s *InstallHotfixRequest) Validate() error {
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

type InstallHotfixRequestInstances struct {
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

func (s InstallHotfixRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s InstallHotfixRequestInstances) GoString() string {
	return s.String()
}

func (s *InstallHotfixRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstallHotfixRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallHotfixRequestInstances) SetInstanceId(v string) *InstallHotfixRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *InstallHotfixRequestInstances) SetRegionId(v string) *InstallHotfixRequestInstances {
	s.RegionId = &v
	return s
}

func (s *InstallHotfixRequestInstances) Validate() error {
	return dara.Validate(s)
}
