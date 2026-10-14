// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotfixStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCveId(v string) *QueryHotfixStatusRequest
	GetCveId() *string
	SetInstances(v []*QueryHotfixStatusRequestInstances) *QueryHotfixStatusRequest
	GetInstances() []*QueryHotfixStatusRequestInstances
}

type QueryHotfixStatusRequest struct {
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
	Instances []*QueryHotfixStatusRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s QueryHotfixStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHotfixStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryHotfixStatusRequest) GetCveId() *string {
	return s.CveId
}

func (s *QueryHotfixStatusRequest) GetInstances() []*QueryHotfixStatusRequestInstances {
	return s.Instances
}

func (s *QueryHotfixStatusRequest) SetCveId(v string) *QueryHotfixStatusRequest {
	s.CveId = &v
	return s
}

func (s *QueryHotfixStatusRequest) SetInstances(v []*QueryHotfixStatusRequestInstances) *QueryHotfixStatusRequest {
	s.Instances = v
	return s
}

func (s *QueryHotfixStatusRequest) Validate() error {
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

type QueryHotfixStatusRequestInstances struct {
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

func (s QueryHotfixStatusRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s QueryHotfixStatusRequestInstances) GoString() string {
	return s.String()
}

func (s *QueryHotfixStatusRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryHotfixStatusRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryHotfixStatusRequestInstances) SetInstanceId(v string) *QueryHotfixStatusRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *QueryHotfixStatusRequestInstances) SetRegionId(v string) *QueryHotfixStatusRequestInstances {
	s.RegionId = &v
	return s
}

func (s *QueryHotfixStatusRequestInstances) Validate() error {
	return dara.Validate(s)
}
