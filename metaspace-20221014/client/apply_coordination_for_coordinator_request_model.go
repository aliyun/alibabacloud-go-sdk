// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationForCoordinatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinationResourceCandidates(v []*ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) *ApplyCoordinationForCoordinatorRequest
	GetCoordinationResourceCandidates() []*ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates
	SetInitiatorType(v string) *ApplyCoordinationForCoordinatorRequest
	GetInitiatorType() *string
	SetUuid(v string) *ApplyCoordinationForCoordinatorRequest
	GetUuid() *string
}

type ApplyCoordinationForCoordinatorRequest struct {
	CoordinationResourceCandidates []*ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates `json:"CoordinationResourceCandidates,omitempty" xml:"CoordinationResourceCandidates,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ADMIN_INITIATE
	InitiatorType *string `json:"InitiatorType,omitempty" xml:"InitiatorType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-uuid-12345
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyCoordinationForCoordinatorRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForCoordinatorRequest) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForCoordinatorRequest) GetCoordinationResourceCandidates() []*ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	return s.CoordinationResourceCandidates
}

func (s *ApplyCoordinationForCoordinatorRequest) GetInitiatorType() *string {
	return s.InitiatorType
}

func (s *ApplyCoordinationForCoordinatorRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ApplyCoordinationForCoordinatorRequest) SetCoordinationResourceCandidates(v []*ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) *ApplyCoordinationForCoordinatorRequest {
	s.CoordinationResourceCandidates = v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequest) SetInitiatorType(v string) *ApplyCoordinationForCoordinatorRequest {
	s.InitiatorType = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequest) SetUuid(v string) *ApplyCoordinationForCoordinatorRequest {
	s.Uuid = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequest) Validate() error {
	return dara.Validate(s)
}

type ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates struct {
	// example:
	//
	// alice
	OwnerEndUserId *string `json:"OwnerEndUserId,omitempty" xml:"OwnerEndUserId,omitempty"`
	// example:
	//
	// 41fd1254d8f7****
	OwnerWyId *string `json:"OwnerWyId,omitempty" xml:"OwnerWyId,omitempty"`
	// example:
	//
	// ai-ij4a6kd4bn2****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// test-resource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// CloudApp
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GetOwnerEndUserId() *string {
	return s.OwnerEndUserId
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GetOwnerWyId() *string {
	return s.OwnerWyId
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GetResourceName() *string {
	return s.ResourceName
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) GetResourceType() *string {
	return s.ResourceType
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) SetOwnerEndUserId(v string) *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	s.OwnerEndUserId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) SetOwnerWyId(v string) *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	s.OwnerWyId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) SetResourceId(v string) *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	s.ResourceId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) SetResourceName(v string) *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	s.ResourceName = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) SetResourceRegionId(v string) *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	s.ResourceRegionId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) SetResourceType(v string) *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates {
	s.ResourceType = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorRequestCoordinationResourceCandidates) Validate() error {
	return dara.Validate(s)
}
