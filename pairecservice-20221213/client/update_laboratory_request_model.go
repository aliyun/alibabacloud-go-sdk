// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCount(v int32) *UpdateLaboratoryRequest
	GetBucketCount() *int32
	SetBucketType(v string) *UpdateLaboratoryRequest
	GetBucketType() *string
	SetBuckets(v string) *UpdateLaboratoryRequest
	GetBuckets() *string
	SetDebugCrowdId(v string) *UpdateLaboratoryRequest
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *UpdateLaboratoryRequest
	GetDebugUsers() *string
	SetDescription(v string) *UpdateLaboratoryRequest
	GetDescription() *string
	SetEnvironment(v string) *UpdateLaboratoryRequest
	GetEnvironment() *string
	SetFilter(v string) *UpdateLaboratoryRequest
	GetFilter() *string
	SetInstanceId(v string) *UpdateLaboratoryRequest
	GetInstanceId() *string
	SetName(v string) *UpdateLaboratoryRequest
	GetName() *string
	SetType(v string) *UpdateLaboratoryRequest
	GetType() *string
}

type UpdateLaboratoryRequest struct {
	// example:
	//
	// 24
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Filter
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// example:
	//
	// 1,2,3,10-20
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// filter=xxx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateLaboratoryRequest) GetBucketCount() *int32 {
	return s.BucketCount
}

func (s *UpdateLaboratoryRequest) GetBucketType() *string {
	return s.BucketType
}

func (s *UpdateLaboratoryRequest) GetBuckets() *string {
	return s.Buckets
}

func (s *UpdateLaboratoryRequest) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *UpdateLaboratoryRequest) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *UpdateLaboratoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLaboratoryRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *UpdateLaboratoryRequest) GetFilter() *string {
	return s.Filter
}

func (s *UpdateLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLaboratoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLaboratoryRequest) GetType() *string {
	return s.Type
}

func (s *UpdateLaboratoryRequest) SetBucketCount(v int32) *UpdateLaboratoryRequest {
	s.BucketCount = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetBucketType(v string) *UpdateLaboratoryRequest {
	s.BucketType = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetBuckets(v string) *UpdateLaboratoryRequest {
	s.Buckets = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetDebugCrowdId(v string) *UpdateLaboratoryRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetDebugUsers(v string) *UpdateLaboratoryRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetDescription(v string) *UpdateLaboratoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetEnvironment(v string) *UpdateLaboratoryRequest {
	s.Environment = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetFilter(v string) *UpdateLaboratoryRequest {
	s.Filter = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetInstanceId(v string) *UpdateLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetName(v string) *UpdateLaboratoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetType(v string) *UpdateLaboratoryRequest {
	s.Type = &v
	return s
}

func (s *UpdateLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
