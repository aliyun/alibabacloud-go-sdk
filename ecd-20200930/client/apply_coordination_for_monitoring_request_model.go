// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationForMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinatePolicyType(v string) *ApplyCoordinationForMonitoringRequest
	GetCoordinatePolicyType() *string
	SetEndUserId(v string) *ApplyCoordinationForMonitoringRequest
	GetEndUserId() *string
	SetInitiatorType(v string) *ApplyCoordinationForMonitoringRequest
	GetInitiatorType() *string
	SetRegionId(v string) *ApplyCoordinationForMonitoringRequest
	GetRegionId() *string
	SetResourceCandidates(v []*ApplyCoordinationForMonitoringRequestResourceCandidates) *ApplyCoordinationForMonitoringRequest
	GetResourceCandidates() []*ApplyCoordinationForMonitoringRequestResourceCandidates
	SetUuid(v string) *ApplyCoordinationForMonitoringRequest
	GetUuid() *string
}

type ApplyCoordinationForMonitoringRequest struct {
	// The access policy during the remote assistance procedure.
	//
	// This parameter is required.
	//
	// example:
	//
	// FULL_CONTROL
	CoordinatePolicyType *string `json:"CoordinatePolicyType,omitempty" xml:"CoordinatePolicyType,omitempty"`
	// The ID of the end user who initiates the coordination flow. This parameter is not required if the request is initiated by an administrator.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The initiator type.
	//
	// example:
	//
	// ADMIN_INITIATE
	InitiatorType *string `json:"InitiatorType,omitempty" xml:"InitiatorType,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of cloud computers that need to accept remote assistance.
	//
	// This parameter is required.
	ResourceCandidates []*ApplyCoordinationForMonitoringRequestResourceCandidates `json:"ResourceCandidates,omitempty" xml:"ResourceCandidates,omitempty" type:"Repeated"`
	// The UUID (unique identifier) of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3E14A18BD4D088504B9F8A8751AB****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyCoordinationForMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForMonitoringRequest) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForMonitoringRequest) GetCoordinatePolicyType() *string {
	return s.CoordinatePolicyType
}

func (s *ApplyCoordinationForMonitoringRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ApplyCoordinationForMonitoringRequest) GetInitiatorType() *string {
	return s.InitiatorType
}

func (s *ApplyCoordinationForMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyCoordinationForMonitoringRequest) GetResourceCandidates() []*ApplyCoordinationForMonitoringRequestResourceCandidates {
	return s.ResourceCandidates
}

func (s *ApplyCoordinationForMonitoringRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ApplyCoordinationForMonitoringRequest) SetCoordinatePolicyType(v string) *ApplyCoordinationForMonitoringRequest {
	s.CoordinatePolicyType = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequest) SetEndUserId(v string) *ApplyCoordinationForMonitoringRequest {
	s.EndUserId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequest) SetInitiatorType(v string) *ApplyCoordinationForMonitoringRequest {
	s.InitiatorType = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequest) SetRegionId(v string) *ApplyCoordinationForMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequest) SetResourceCandidates(v []*ApplyCoordinationForMonitoringRequestResourceCandidates) *ApplyCoordinationForMonitoringRequest {
	s.ResourceCandidates = v
	return s
}

func (s *ApplyCoordinationForMonitoringRequest) SetUuid(v string) *ApplyCoordinationForMonitoringRequest {
	s.Uuid = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequest) Validate() error {
	if s.ResourceCandidates != nil {
		for _, item := range s.ResourceCandidates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyCoordinationForMonitoringRequestResourceCandidates struct {
	// The Alibaba Cloud account ID of the cloud computer administrator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130247021517****
	OwnerAliUid *int64 `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
	// The username of the current user of the cloud computer.
	//
	// > This field is required.
	//
	// example:
	//
	// alice
	OwnerEndUserId *string `json:"OwnerEndUserId,omitempty" xml:"OwnerEndUserId,omitempty"`
	// The cloud computer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-08zhejm3h7ilr****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The cloud computer name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoComputer
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The properties of the cloud computer.
	//
	// example:
	//
	// TestProperty
	ResourceProperties *string `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty"`
	// The region where the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ApplyCoordinationForMonitoringRequestResourceCandidates) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForMonitoringRequestResourceCandidates) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetOwnerAliUid() *int64 {
	return s.OwnerAliUid
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetOwnerEndUserId() *string {
	return s.OwnerEndUserId
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetResourceName() *string {
	return s.ResourceName
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetResourceProperties() *string {
	return s.ResourceProperties
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) GetResourceType() *string {
	return s.ResourceType
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetOwnerAliUid(v int64) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.OwnerAliUid = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetOwnerEndUserId(v string) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.OwnerEndUserId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetResourceId(v string) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.ResourceId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetResourceName(v string) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.ResourceName = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetResourceProperties(v string) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.ResourceProperties = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetResourceRegionId(v string) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.ResourceRegionId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) SetResourceType(v string) *ApplyCoordinationForMonitoringRequestResourceCandidates {
	s.ResourceType = &v
	return s
}

func (s *ApplyCoordinationForMonitoringRequestResourceCandidates) Validate() error {
	return dara.Validate(s)
}
