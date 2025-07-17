// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListInstancesResponseBody
	GetTotalCount() *int64
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 2CA11923-2A3D-5E5A-8314-E699D2DD15DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstances struct {
	FeatureDBInfo *ListInstancesResponseBodyInstancesFeatureDBInfo `json:"FeatureDBInfo,omitempty" xml:"FeatureDBInfo,omitempty" type:"Struct"`
	// Deprecated
	FeatureDBInstanceInfo *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo `json:"FeatureDBInstanceInfo,omitempty" xml:"FeatureDBInstanceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// featureStore-cn-7mz2xfu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetFeatureDBInfo() *ListInstancesResponseBodyInstancesFeatureDBInfo {
	return s.FeatureDBInfo
}

func (s *ListInstancesResponseBodyInstances) GetFeatureDBInstanceInfo() *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo {
	return s.FeatureDBInstanceInfo
}

func (s *ListInstancesResponseBodyInstances) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstances) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstances) SetFeatureDBInfo(v *ListInstancesResponseBodyInstancesFeatureDBInfo) *ListInstancesResponseBodyInstances {
	s.FeatureDBInfo = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetFeatureDBInstanceInfo(v *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) *ListInstancesResponseBodyInstances {
	s.FeatureDBInstanceInfo = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetType(v string) *ListInstancesResponseBodyInstances {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesFeatureDBInfo struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesFeatureDBInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesFeatureDBInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInfo) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInfo) SetStatus(v string) *ListInstancesResponseBodyInstancesFeatureDBInfo {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInfo) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesFeatureDBInstanceInfo struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) SetStatus(v string) *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) Validate() error {
	return dara.Validate(s)
}
