// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody
	GetInstance() *GetInstanceResponseBodyInstance
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
}

type GetInstanceResponseBody struct {
	// The instance information.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetInstance() *GetInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyInstance struct {
	// The time when the instance was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1550115455000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// enabled
	CrossRegionReplication *string `json:"CrossRegionReplication,omitempty" xml:"CrossRegionReplication,omitempty"`
	// example:
	//
	// primary
	CrossRegionReplicationRole *string `json:"CrossRegionReplicationRole,omitempty" xml:"CrossRegionReplicationRole,omitempty"`
	// The default domain name of the instance. This field is no longer maintained. Use the DomainConfig fields or refer to the query domain name list operation instead.
	DefaultEndpoint *GetInstanceResponseBodyInstanceDefaultEndpoint `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty" type:"Struct"`
	// The description of the instance.
	//
	// example:
	//
	// instance_for_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name configuration of the instance.
	DomainConfig *GetInstanceResponseBodyInstanceDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
	// The public egress CIDR blocks of the instance. For example, during Active Directory (AD) account synchronization, the EIAM instance accesses your AD server through these public CIDR blocks.
	EgressAddresses []*string `json:"EgressAddresses,omitempty" xml:"EgressAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// inactive
	InstanceFailoverStatus *string `json:"InstanceFailoverStatus,omitempty" xml:"InstanceFailoverStatus,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_abt3pfwojojcq323si6g5xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ServiceCode of the Alibaba Cloud service that manages the instance.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// sase
	ManagedServiceCode       *string                                                  `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	ReplicationConfiguration *GetInstanceResponseBodyInstanceReplicationConfiguration `json:"ReplicationConfiguration,omitempty" xml:"ReplicationConfiguration,omitempty" type:"Struct"`
	// Indicates whether the instance is managed by an Alibaba Cloud service.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// false
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

func (s GetInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyInstance) GetCrossRegionReplication() *string {
	return s.CrossRegionReplication
}

func (s *GetInstanceResponseBodyInstance) GetCrossRegionReplicationRole() *string {
	return s.CrossRegionReplicationRole
}

func (s *GetInstanceResponseBodyInstance) GetDefaultEndpoint() *GetInstanceResponseBodyInstanceDefaultEndpoint {
	return s.DefaultEndpoint
}

func (s *GetInstanceResponseBodyInstance) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceResponseBodyInstance) GetDomainConfig() *GetInstanceResponseBodyInstanceDomainConfig {
	return s.DomainConfig
}

func (s *GetInstanceResponseBodyInstance) GetEgressAddresses() []*string {
	return s.EgressAddresses
}

func (s *GetInstanceResponseBodyInstance) GetInstanceFailoverStatus() *string {
	return s.InstanceFailoverStatus
}

func (s *GetInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyInstance) GetManagedServiceCode() *string {
	return s.ManagedServiceCode
}

func (s *GetInstanceResponseBodyInstance) GetReplicationConfiguration() *GetInstanceResponseBodyInstanceReplicationConfiguration {
	return s.ReplicationConfiguration
}

func (s *GetInstanceResponseBodyInstance) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *GetInstanceResponseBodyInstance) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyInstance) SetCreateTime(v int64) *GetInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCrossRegionReplication(v string) *GetInstanceResponseBodyInstance {
	s.CrossRegionReplication = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCrossRegionReplicationRole(v string) *GetInstanceResponseBodyInstance {
	s.CrossRegionReplicationRole = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDefaultEndpoint(v *GetInstanceResponseBodyInstanceDefaultEndpoint) *GetInstanceResponseBodyInstance {
	s.DefaultEndpoint = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDescription(v string) *GetInstanceResponseBodyInstance {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDomainConfig(v *GetInstanceResponseBodyInstanceDomainConfig) *GetInstanceResponseBodyInstance {
	s.DomainConfig = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEgressAddresses(v []*string) *GetInstanceResponseBodyInstance {
	s.EgressAddresses = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceFailoverStatus(v string) *GetInstanceResponseBodyInstance {
	s.InstanceFailoverStatus = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetManagedServiceCode(v string) *GetInstanceResponseBodyInstance {
	s.ManagedServiceCode = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetReplicationConfiguration(v *GetInstanceResponseBodyInstanceReplicationConfiguration) *GetInstanceResponseBodyInstance {
	s.ReplicationConfiguration = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetServiceManaged(v bool) *GetInstanceResponseBodyInstance {
	s.ServiceManaged = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStatus(v string) *GetInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) Validate() error {
	if s.DefaultEndpoint != nil {
		if err := s.DefaultEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.DomainConfig != nil {
		if err := s.DomainConfig.Validate(); err != nil {
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

type GetInstanceResponseBodyInstanceDefaultEndpoint struct {
	// The domain name of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The status of the instance domain name. Valid values:
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

func (s GetInstanceResponseBodyInstanceDefaultEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceDefaultEndpoint) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) SetEndpoint(v string) *GetInstanceResponseBodyInstanceDefaultEndpoint {
	s.Endpoint = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) SetStatus(v string) *GetInstanceResponseBodyInstanceDefaultEndpoint {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceDomainConfig struct {
	// The default domain name of the instance.
	//
	// example:
	//
	// login.example.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The initialization domain name of the instance.
	//
	// example:
	//
	// rx72nxxx.example.com
	InitDomain *string `json:"InitDomain,omitempty" xml:"InitDomain,omitempty"`
	// The automatic redirect status of the initialization domain name. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// disabled
	InitDomainAutoRedirectStatus *string `json:"InitDomainAutoRedirectStatus,omitempty" xml:"InitDomainAutoRedirectStatus,omitempty"`
}

func (s GetInstanceResponseBodyInstanceDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceDomainConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) GetDefaultDomain() *string {
	return s.DefaultDomain
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) GetInitDomain() *string {
	return s.InitDomain
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) GetInitDomainAutoRedirectStatus() *string {
	return s.InitDomainAutoRedirectStatus
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetDefaultDomain(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.DefaultDomain = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetInitDomain(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.InitDomain = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetInitDomainAutoRedirectStatus(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.InitDomainAutoRedirectStatus = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceReplicationConfiguration struct {
	// example:
	//
	// idaas_xxxx
	BackupInstanceId *string `json:"BackupInstanceId,omitempty" xml:"BackupInstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	BackupInstanceRegionId *string `json:"BackupInstanceRegionId,omitempty" xml:"BackupInstanceRegionId,omitempty"`
	// example:
	//
	// idaas_xxxx
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	PrimaryInstanceRegionId *string `json:"PrimaryInstanceRegionId,omitempty" xml:"PrimaryInstanceRegionId,omitempty"`
	// example:
	//
	// 1778499337000
	ReplicationCreateTime *int64 `json:"ReplicationCreateTime,omitempty" xml:"ReplicationCreateTime,omitempty"`
}

func (s GetInstanceResponseBodyInstanceReplicationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceReplicationConfiguration) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) GetBackupInstanceId() *string {
	return s.BackupInstanceId
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) GetBackupInstanceRegionId() *string {
	return s.BackupInstanceRegionId
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) GetPrimaryInstanceId() *string {
	return s.PrimaryInstanceId
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) GetPrimaryInstanceRegionId() *string {
	return s.PrimaryInstanceRegionId
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) GetReplicationCreateTime() *int64 {
	return s.ReplicationCreateTime
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) SetBackupInstanceId(v string) *GetInstanceResponseBodyInstanceReplicationConfiguration {
	s.BackupInstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) SetBackupInstanceRegionId(v string) *GetInstanceResponseBodyInstanceReplicationConfiguration {
	s.BackupInstanceRegionId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) SetPrimaryInstanceId(v string) *GetInstanceResponseBodyInstanceReplicationConfiguration {
	s.PrimaryInstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) SetPrimaryInstanceRegionId(v string) *GetInstanceResponseBodyInstanceReplicationConfiguration {
	s.PrimaryInstanceRegionId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) SetReplicationCreateTime(v int64) *GetInstanceResponseBodyInstanceReplicationConfiguration {
	s.ReplicationCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceReplicationConfiguration) Validate() error {
	return dara.Validate(s)
}
