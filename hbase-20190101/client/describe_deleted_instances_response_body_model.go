// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeletedInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeDeletedInstancesResponseBodyInstances) *DescribeDeletedInstancesResponseBody
	GetInstances() *DescribeDeletedInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeDeletedInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeletedInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDeletedInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDeletedInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeDeletedInstancesResponseBody struct {
	Instances *DescribeDeletedInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0CAC5702-C862-44C0-AD54-C9CE70F4B246
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeletedInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBody) GetInstances() *DescribeDeletedInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeDeletedInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeletedInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeletedInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeletedInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDeletedInstancesResponseBody) SetInstances(v *DescribeDeletedInstancesResponseBodyInstances) *DescribeDeletedInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetPageNumber(v int32) *DescribeDeletedInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetPageSize(v int32) *DescribeDeletedInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetRequestId(v string) *DescribeDeletedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetTotalCount(v int64) *DescribeDeletedInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeletedInstancesResponseBodyInstances struct {
	Instance []*DescribeDeletedInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDeletedInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBodyInstances) GetInstance() []*DescribeDeletedInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeDeletedInstancesResponseBodyInstances) SetInstance(v []*DescribeDeletedInstancesResponseBodyInstancesInstance) *DescribeDeletedInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeletedInstancesResponseBodyInstancesInstance struct {
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 2020-11-02T07:16:07Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2020-11-02T07:27:24Z
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// hb-bp10q7n2zdw12xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// e2e-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// null
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// null
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDeletedInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetDeleteTime() *string {
	return s.DeleteTime
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetModuleStackVersion() *string {
	return s.ModuleStackVersion
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetClusterType(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ClusterType = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetCreatedTime(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetDeleteTime(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.DeleteTime = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetEngine(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetMajorVersion(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.MajorVersion = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetModuleStackVersion(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetParentId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ParentId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
