// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyCollectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollectors(v []*ListIntegrationPolicyCollectorsResponseBodyCollectors) *ListIntegrationPolicyCollectorsResponseBody
	GetCollectors() []*ListIntegrationPolicyCollectorsResponseBodyCollectors
	SetRequestId(v string) *ListIntegrationPolicyCollectorsResponseBody
	GetRequestId() *string
}

type ListIntegrationPolicyCollectorsResponseBody struct {
	// The list of collectors.
	Collectors []*ListIntegrationPolicyCollectorsResponseBodyCollectors `json:"collectors,omitempty" xml:"collectors,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIntegrationPolicyCollectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsResponseBody) GetCollectors() []*ListIntegrationPolicyCollectorsResponseBodyCollectors {
	return s.Collectors
}

func (s *ListIntegrationPolicyCollectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyCollectorsResponseBody) SetCollectors(v []*ListIntegrationPolicyCollectorsResponseBodyCollectors) *ListIntegrationPolicyCollectorsResponseBody {
	s.Collectors = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBody) SetRequestId(v string) *ListIntegrationPolicyCollectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBody) Validate() error {
	if s.Collectors != nil {
		for _, item := range s.Collectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyCollectorsResponseBodyCollectors struct {
	// The add-on details.
	AddonMeta *AddonMeta `json:"addonMeta,omitempty" xml:"addonMeta,omitempty"`
	// The collector name.
	//
	// example:
	//
	// collector-kkxx
	CollectorName *string `json:"collectorName,omitempty" xml:"collectorName,omitempty"`
	// The collector type.
	//
	// example:
	//
	// Exporter
	CollectorType *string `json:"collectorType,omitempty" xml:"collectorType,omitempty"`
	// The phase status.
	Conditions []*ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// Indicates whether the component is a managed component.
	//
	// example:
	//
	// false
	Managed *bool `json:"managed,omitempty" xml:"managed,omitempty"`
	// The name of the add-on release.
	//
	// example:
	//
	// ecs-loong-collector-i-f8z1176fg57rlwmc1rfi
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// The collector status.
	//
	// example:
	//
	// xx
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The component version.
	//
	// example:
	//
	// 2.6.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The list of workloads.
	Workloads []*ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads `json:"workloads,omitempty" xml:"workloads,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectors) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectors) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetAddonMeta() *AddonMeta {
	return s.AddonMeta
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetCollectorName() *string {
	return s.CollectorName
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetCollectorType() *string {
	return s.CollectorType
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetConditions() []*ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	return s.Conditions
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetManaged() *bool {
	return s.Managed
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetState() *string {
	return s.State
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetVersion() *string {
	return s.Version
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) GetWorkloads() []*ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	return s.Workloads
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetAddonMeta(v *AddonMeta) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.AddonMeta = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetCollectorName(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.CollectorName = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetCollectorType(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.CollectorType = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetConditions(v []*ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.Conditions = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetManaged(v bool) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.Managed = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetReleaseName(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.ReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetState(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.State = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetVersion(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.Version = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) SetWorkloads(v []*ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) *ListIntegrationPolicyCollectorsResponseBodyCollectors {
	s.Workloads = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectors) Validate() error {
	if s.AddonMeta != nil {
		if err := s.AddonMeta.Validate(); err != nil {
			return err
		}
	}
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Workloads != nil {
		for _, item := range s.Workloads {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions struct {
	// The time of the first transition.
	//
	// example:
	//
	// 2024-08-27T13:59:23+08:00
	FirstTransitionTime *string `json:"firstTransitionTime,omitempty" xml:"firstTransitionTime,omitempty"`
	// The time of the last transition.
	//
	// example:
	//
	// 2024-08-27T13:59:23+08:00
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
	// The details.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// Probe
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The phase status.
	//
	// example:
	//
	// True
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The phase type.
	//
	// example:
	//
	// Ready
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GetMessage() *string {
	return s.Message
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GetReason() *string {
	return s.Reason
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GetStatus() *string {
	return s.Status
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) GetType() *string {
	return s.Type
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) SetFirstTransitionTime(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) SetLastTransitionTime(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) SetMessage(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	s.Message = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) SetReason(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	s.Reason = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) SetStatus(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	s.Status = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) SetType(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions {
	s.Type = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsConditions) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads struct {
	// The host IP address.
	//
	// example:
	//
	// 10.10.10.10
	HostIp *string `json:"hostIp,omitempty" xml:"hostIp,omitempty"`
	// The IP address of the workload.
	//
	// example:
	//
	// 11.193.82.198
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// Indicates whether the component is a managed component.
	//
	// example:
	//
	// false
	Managed *bool `json:"managed,omitempty" xml:"managed,omitempty"`
	// The policy management information.
	ManagedInfo *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo `json:"managedInfo,omitempty" xml:"managedInfo,omitempty" type:"Struct"`
	// The details.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The workload name.
	//
	// example:
	//
	// exporter-xxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The workload namespace.
	//
	// example:
	//
	// prod-db
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The kind of the parent reference.
	//
	// example:
	//
	// Deployment
	OwnerReferenceKind *string `json:"ownerReferenceKind,omitempty" xml:"ownerReferenceKind,omitempty"`
	// The name of the parent reference.
	//
	// example:
	//
	// exporter
	OwnerReferenceName *string `json:"ownerReferenceName,omitempty" xml:"ownerReferenceName,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2024-08-27T13:59:23+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status.
	//
	// example:
	//
	// {}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The workload version.
	//
	// example:
	//
	// v2.4.4
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetHostIp() *string {
	return s.HostIp
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetIp() *string {
	return s.Ip
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetManaged() *bool {
	return s.Managed
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetManagedInfo() *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo {
	return s.ManagedInfo
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetMessage() *string {
	return s.Message
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetOwnerReferenceKind() *string {
	return s.OwnerReferenceKind
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetOwnerReferenceName() *string {
	return s.OwnerReferenceName
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetStartTime() *string {
	return s.StartTime
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetStatus() *string {
	return s.Status
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) GetVersion() *string {
	return s.Version
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetHostIp(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.HostIp = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetIp(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Ip = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetManaged(v bool) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Managed = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetManagedInfo(v *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.ManagedInfo = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetMessage(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Message = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetName(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetNamespace(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetOwnerReferenceKind(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.OwnerReferenceKind = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetOwnerReferenceName(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.OwnerReferenceName = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetStartTime(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.StartTime = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetStatus(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Status = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) SetVersion(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads {
	s.Version = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloads) Validate() error {
	if s.ManagedInfo != nil {
		if err := s.ManagedInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo struct {
	// The security group ID.
	//
	// example:
	//
	// sg-xxxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxxx
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) SetSecurityGroupId(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) SetVswitchId(v string) *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo {
	s.VswitchId = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponseBodyCollectorsWorkloadsManagedInfo) Validate() error {
	return dara.Validate(s)
}
