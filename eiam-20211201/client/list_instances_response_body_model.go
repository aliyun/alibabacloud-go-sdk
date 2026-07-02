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
	// The list of instance information.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
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

type ListInstancesResponseBodyInstances struct {
	// The instance creation time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1550115455000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether cross-region replication is enabled. Valid values: enabled or disabled.
	//
	// example:
	//
	// enabled
	CrossRegionReplication *string `json:"CrossRegionReplication,omitempty" xml:"CrossRegionReplication,omitempty"`
	// The cross-region replication role. Valid values: primary (primary instance) or backup (backup instance).
	//
	// example:
	//
	// primary
	CrossRegionReplicationRole *string `json:"CrossRegionReplicationRole,omitempty" xml:"CrossRegionReplicationRole,omitempty"`
	// The default endpoint of the instance.
	DefaultEndpoint *ListInstancesResponseBodyInstancesDefaultEndpoint `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty" type:"Struct"`
	// The instance description.
	//
	// example:
	//
	// instance_for_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance failover activation status. Valid values: active (activated) or inactive (not activated).
	//
	// example:
	//
	// inactive
	InstanceFailoverStatus *string `json:"InstanceFailoverStatus,omitempty" xml:"InstanceFailoverStatus,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_eypq6ljgyeuwmlw672sulxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The service code of the Alibaba Cloud service that manages the instance.
	//
	// example:
	//
	// sase
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// The replication configuration. This parameter is returned only when CrossRegionReplication is set to enabled.
	ReplicationConfiguration *ListInstancesResponseBodyInstancesReplicationConfiguration `json:"ReplicationConfiguration,omitempty" xml:"ReplicationConfiguration,omitempty" type:"Struct"`
	// Indicates whether the instance is managed by an Alibaba Cloud service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The instance status. Valid values:
	//
	// - creating: Being created.
	//
	// - running: Running.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyInstances) GetCrossRegionReplication() *string {
	return s.CrossRegionReplication
}

func (s *ListInstancesResponseBodyInstances) GetCrossRegionReplicationRole() *string {
	return s.CrossRegionReplicationRole
}

func (s *ListInstancesResponseBodyInstances) GetDefaultEndpoint() *ListInstancesResponseBodyInstancesDefaultEndpoint {
	return s.DefaultEndpoint
}

func (s *ListInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyInstances) GetInstanceFailoverStatus() *string {
	return s.InstanceFailoverStatus
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetManagedServiceCode() *string {
	return s.ManagedServiceCode
}

func (s *ListInstancesResponseBodyInstances) GetReplicationConfiguration() *ListInstancesResponseBodyInstancesReplicationConfiguration {
	return s.ReplicationConfiguration
}

func (s *ListInstancesResponseBodyInstances) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCrossRegionReplication(v string) *ListInstancesResponseBodyInstances {
	s.CrossRegionReplication = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCrossRegionReplicationRole(v string) *ListInstancesResponseBodyInstances {
	s.CrossRegionReplicationRole = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDefaultEndpoint(v *ListInstancesResponseBodyInstancesDefaultEndpoint) *ListInstancesResponseBodyInstances {
	s.DefaultEndpoint = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceFailoverStatus(v string) *ListInstancesResponseBodyInstances {
	s.InstanceFailoverStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetManagedServiceCode(v string) *ListInstancesResponseBodyInstances {
	s.ManagedServiceCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetReplicationConfiguration(v *ListInstancesResponseBodyInstancesReplicationConfiguration) *ListInstancesResponseBodyInstances {
	s.ReplicationConfiguration = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetServiceManaged(v bool) *ListInstancesResponseBodyInstances {
	s.ServiceManaged = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	if s.DefaultEndpoint != nil {
		if err := s.DefaultEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.ReplicationConfiguration != nil {
		if err := s.ReplicationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesDefaultEndpoint struct {
	// The endpoint address of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The endpoint status. Valid values:
	//
	// - resolved: Resolved.
	//
	// - unresolved: Not resolved.
	//
	// example:
	//
	// resolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDefaultEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDefaultEndpoint) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) SetEndpoint(v string) *ListInstancesResponseBodyInstancesDefaultEndpoint {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) SetStatus(v string) *ListInstancesResponseBodyInstancesDefaultEndpoint {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesReplicationConfiguration struct {
	// The instance ID of the backup instance.
	//
	// example:
	//
	// idaas_xxxxxx
	BackupInstanceId *string `json:"BackupInstanceId,omitempty" xml:"BackupInstanceId,omitempty"`
	// The region ID of the backup instance.
	//
	// example:
	//
	// cn-beijing
	BackupInstanceRegionId *string `json:"BackupInstanceRegionId,omitempty" xml:"BackupInstanceRegionId,omitempty"`
	// The instance ID of the primary instance.
	//
	// example:
	//
	// idaas_xxxxxx
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// The region ID of the primary instance.
	//
	// example:
	//
	// cn-hangzhou
	PrimaryInstanceRegionId *string `json:"PrimaryInstanceRegionId,omitempty" xml:"PrimaryInstanceRegionId,omitempty"`
	// The time when the disaster recovery data replication was created, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1778499337000
	ReplicationCreateTime *int64 `json:"ReplicationCreateTime,omitempty" xml:"ReplicationCreateTime,omitempty"`
}

func (s ListInstancesResponseBodyInstancesReplicationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesReplicationConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) GetBackupInstanceId() *string {
	return s.BackupInstanceId
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) GetBackupInstanceRegionId() *string {
	return s.BackupInstanceRegionId
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) GetPrimaryInstanceId() *string {
	return s.PrimaryInstanceId
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) GetPrimaryInstanceRegionId() *string {
	return s.PrimaryInstanceRegionId
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) GetReplicationCreateTime() *int64 {
	return s.ReplicationCreateTime
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) SetBackupInstanceId(v string) *ListInstancesResponseBodyInstancesReplicationConfiguration {
	s.BackupInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) SetBackupInstanceRegionId(v string) *ListInstancesResponseBodyInstancesReplicationConfiguration {
	s.BackupInstanceRegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) SetPrimaryInstanceId(v string) *ListInstancesResponseBodyInstancesReplicationConfiguration {
	s.PrimaryInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) SetPrimaryInstanceRegionId(v string) *ListInstancesResponseBodyInstancesReplicationConfiguration {
	s.PrimaryInstanceRegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) SetReplicationCreateTime(v int64) *ListInstancesResponseBodyInstancesReplicationConfiguration {
	s.ReplicationCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesReplicationConfiguration) Validate() error {
	return dara.Validate(s)
}
