// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCount(v int32) *CreateLaboratoryRequest
	GetBucketCount() *int32
	SetBucketType(v string) *CreateLaboratoryRequest
	GetBucketType() *string
	SetBuckets(v string) *CreateLaboratoryRequest
	GetBuckets() *string
	SetDebugCrowdId(v string) *CreateLaboratoryRequest
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *CreateLaboratoryRequest
	GetDebugUsers() *string
	SetDescription(v string) *CreateLaboratoryRequest
	GetDescription() *string
	SetEnvironment(v string) *CreateLaboratoryRequest
	GetEnvironment() *string
	SetFilter(v string) *CreateLaboratoryRequest
	GetFilter() *string
	SetInstanceId(v string) *CreateLaboratoryRequest
	GetInstanceId() *string
	SetName(v string) *CreateLaboratoryRequest
	GetName() *string
	SetSceneId(v string) *CreateLaboratoryRequest
	GetSceneId() *string
	SetType(v string) *CreateLaboratoryRequest
	GetType() *string
}

type CreateLaboratoryRequest struct {
	// example:
	//
	// 24
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UidHash
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
	// This parameter is required.
	//
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
	// pairec-cn-abcdefg1234
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
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *CreateLaboratoryRequest) GetBucketCount() *int32 {
	return s.BucketCount
}

func (s *CreateLaboratoryRequest) GetBucketType() *string {
	return s.BucketType
}

func (s *CreateLaboratoryRequest) GetBuckets() *string {
	return s.Buckets
}

func (s *CreateLaboratoryRequest) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *CreateLaboratoryRequest) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *CreateLaboratoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLaboratoryRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CreateLaboratoryRequest) GetFilter() *string {
	return s.Filter
}

func (s *CreateLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLaboratoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateLaboratoryRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateLaboratoryRequest) GetType() *string {
	return s.Type
}

func (s *CreateLaboratoryRequest) SetBucketCount(v int32) *CreateLaboratoryRequest {
	s.BucketCount = &v
	return s
}

func (s *CreateLaboratoryRequest) SetBucketType(v string) *CreateLaboratoryRequest {
	s.BucketType = &v
	return s
}

func (s *CreateLaboratoryRequest) SetBuckets(v string) *CreateLaboratoryRequest {
	s.Buckets = &v
	return s
}

func (s *CreateLaboratoryRequest) SetDebugCrowdId(v string) *CreateLaboratoryRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *CreateLaboratoryRequest) SetDebugUsers(v string) *CreateLaboratoryRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateLaboratoryRequest) SetDescription(v string) *CreateLaboratoryRequest {
	s.Description = &v
	return s
}

func (s *CreateLaboratoryRequest) SetEnvironment(v string) *CreateLaboratoryRequest {
	s.Environment = &v
	return s
}

func (s *CreateLaboratoryRequest) SetFilter(v string) *CreateLaboratoryRequest {
	s.Filter = &v
	return s
}

func (s *CreateLaboratoryRequest) SetInstanceId(v string) *CreateLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLaboratoryRequest) SetName(v string) *CreateLaboratoryRequest {
	s.Name = &v
	return s
}

func (s *CreateLaboratoryRequest) SetSceneId(v string) *CreateLaboratoryRequest {
	s.SceneId = &v
	return s
}

func (s *CreateLaboratoryRequest) SetType(v string) *CreateLaboratoryRequest {
	s.Type = &v
	return s
}

func (s *CreateLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
