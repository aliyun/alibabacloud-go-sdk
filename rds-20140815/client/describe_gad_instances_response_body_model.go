// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstances(v []*DescribeGadInstancesResponseBodyGadInstances) *DescribeGadInstancesResponseBody
	GetGadInstances() []*DescribeGadInstancesResponseBodyGadInstances
	SetRequestId(v string) *DescribeGadInstancesResponseBody
	GetRequestId() *string
}

type DescribeGadInstancesResponseBody struct {
	GadInstances []*DescribeGadInstancesResponseBodyGadInstances `json:"GadInstances,omitempty" xml:"GadInstances,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGadInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBody) GetGadInstances() []*DescribeGadInstancesResponseBodyGadInstances {
	return s.GadInstances
}

func (s *DescribeGadInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGadInstancesResponseBody) SetGadInstances(v []*DescribeGadInstancesResponseBodyGadInstances) *DescribeGadInstancesResponseBody {
	s.GadInstances = v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetRequestId(v string) *DescribeGadInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) Validate() error {
	if s.GadInstances != nil {
		for _, item := range s.GadInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstancesResponseBodyGadInstances struct {
	CreationTime       *string                                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description        *string                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	GadInstanceMembers []*DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers `json:"GadInstanceMembers,omitempty" xml:"GadInstanceMembers,omitempty" type:"Repeated"`
	GadInstanceName    *string                                                           `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	ModificationTime   *string                                                           `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Service            *string                                                           `json:"Service,omitempty" xml:"Service,omitempty"`
	Status             *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstancesResponseBodyGadInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBodyGadInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetGadInstanceMembers() []*DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	return s.GadInstanceMembers
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetService() *string {
	return s.Service
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetCreationTime(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetDescription(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.Description = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetGadInstanceMembers(v []*DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) *DescribeGadInstancesResponseBodyGadInstances {
	s.GadInstanceMembers = v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetGadInstanceName(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.GadInstanceName = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetModificationTime(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.ModificationTime = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetService(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.Service = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetStatus(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.Status = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) Validate() error {
	if s.GadInstanceMembers != nil {
		for _, item := range s.GadInstanceMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers struct {
	DBInstanceID    *string `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	DtsInstance     *string `json:"DtsInstance,omitempty" xml:"DtsInstance,omitempty"`
	Engine          *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion   *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetDBInstanceID() *string {
	return s.DBInstanceID
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetDtsInstance() *string {
	return s.DtsInstance
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetEngine() *string {
	return s.Engine
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetRole() *string {
	return s.Role
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetDBInstanceID(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetDtsInstance(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.DtsInstance = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetEngine(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.Engine = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetEngineVersion(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.EngineVersion = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetRegionId(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetResourceGroupId(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetRole(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.Role = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetStatus(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.Status = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) Validate() error {
	return dara.Validate(s)
}
