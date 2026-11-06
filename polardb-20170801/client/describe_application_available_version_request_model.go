// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationAvailableVersionRequest
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationAvailableVersionRequest
	GetApplicationType() *string
	SetComponents(v []*DescribeApplicationAvailableVersionRequestComponents) *DescribeApplicationAvailableVersionRequest
	GetComponents() []*DescribeApplicationAvailableVersionRequestComponents
	SetRegionId(v string) *DescribeApplicationAvailableVersionRequest
	GetRegionId() *string
}

type DescribeApplicationAvailableVersionRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The list of application subcomponents.
	Components []*DescribeApplicationAvailableVersionRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApplicationAvailableVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableVersionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationAvailableVersionRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationAvailableVersionRequest) GetComponents() []*DescribeApplicationAvailableVersionRequestComponents {
	return s.Components
}

func (s *DescribeApplicationAvailableVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationAvailableVersionRequest) SetApplicationId(v string) *DescribeApplicationAvailableVersionRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequest) SetApplicationType(v string) *DescribeApplicationAvailableVersionRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequest) SetComponents(v []*DescribeApplicationAvailableVersionRequestComponents) *DescribeApplicationAvailableVersionRequest {
	s.Components = v
	return s
}

func (s *DescribeApplicationAvailableVersionRequest) SetRegionId(v string) *DescribeApplicationAvailableVersionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationAvailableVersionRequestComponents struct {
	// The specifications of the application subcomponent.
	//
	// example:
	//
	// polar.app.g2.medium
	ComponentClass *string `json:"ComponentClass,omitempty" xml:"ComponentClass,omitempty"`
	// The maximum number of application subcomponents of the same specifications. Default value: the value of ComponentReplica. This parameter is supported only for raycluster applications.
	//
	// example:
	//
	// 1
	ComponentMaxReplica *int64 `json:"ComponentMaxReplica,omitempty" xml:"ComponentMaxReplica,omitempty"`
	// The number of replicas for the application subcomponent. Default value: 1.
	//
	// example:
	//
	// 1
	ComponentReplica *int64 `json:"ComponentReplica,omitempty" xml:"ComponentReplica,omitempty"`
	// The type of the application subcomponent.
	//
	// example:
	//
	// gateway
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
}

func (s DescribeApplicationAvailableVersionRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableVersionRequestComponents) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableVersionRequestComponents) GetComponentClass() *string {
	return s.ComponentClass
}

func (s *DescribeApplicationAvailableVersionRequestComponents) GetComponentMaxReplica() *int64 {
	return s.ComponentMaxReplica
}

func (s *DescribeApplicationAvailableVersionRequestComponents) GetComponentReplica() *int64 {
	return s.ComponentReplica
}

func (s *DescribeApplicationAvailableVersionRequestComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationAvailableVersionRequestComponents) SetComponentClass(v string) *DescribeApplicationAvailableVersionRequestComponents {
	s.ComponentClass = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequestComponents) SetComponentMaxReplica(v int64) *DescribeApplicationAvailableVersionRequestComponents {
	s.ComponentMaxReplica = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequestComponents) SetComponentReplica(v int64) *DescribeApplicationAvailableVersionRequestComponents {
	s.ComponentReplica = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequestComponents) SetComponentType(v string) *DescribeApplicationAvailableVersionRequestComponents {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationAvailableVersionRequestComponents) Validate() error {
	return dara.Validate(s)
}
