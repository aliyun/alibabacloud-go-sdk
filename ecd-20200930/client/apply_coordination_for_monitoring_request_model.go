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
	// The coordination policy.
	//
	// Set the value to FULL_CONTROL.
	//
	// 	- The value FULL_CONTROL specifies that the cloud desktop is shared and remote access to the cloud desktop is allowed.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// FULL_CONTROL
	CoordinatePolicyType *string `json:"CoordinatePolicyType,omitempty" xml:"CoordinatePolicyType,omitempty"`
	// The ID of the end user who initiates the stream collaboration. If the initiator is the administrator, do not specify this parameter.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The type of the initiator.
	//
	// Set the value to ADMIN_INITIATE.
	//
	// 	- The value ADMIN_INITIATE specifies that the administrator initiates the coordination request.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ADMIN_INITIATE
	InitiatorType *string `json:"InitiatorType,omitempty" xml:"InitiatorType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://next.api.aliyun.com/document/ecd/2020-09-30/DescribeRegions) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of cloud desktops that run the collaboration task at the same time.
	//
	// This parameter is required.
	ResourceCandidates []*ApplyCoordinationForMonitoringRequestResourceCandidates `json:"ResourceCandidates,omitempty" xml:"ResourceCandidates,omitempty" type:"Repeated"`
	// The universally unique identifier (UUID) of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 62f2f1f252f04e0e9d8bc****
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
	// The ID of the Alibaba Cloud account to which the current cloud desktop belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130247021517****
	OwnerAliUid *int64 `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
	// The ID of the current end user.
	//
	// example:
	//
	// alice
	OwnerEndUserId *string `json:"OwnerEndUserId,omitempty" xml:"OwnerEndUserId,omitempty"`
	// The ID of the cloud desktop.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-08zhejm3h7ilr****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the cloud desktop.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDesktop
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The properties of the cloud desktop.
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
	// Set the value to CLOUD_DESKTOP.
	//
	// 	- The value CLOUD_DESKTOP specifies that the resource is a cloud desktop.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
