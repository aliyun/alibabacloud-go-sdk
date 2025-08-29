// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCount(v int32) *GetLaboratoryResponseBody
	GetBucketCount() *int32
	SetBucketType(v string) *GetLaboratoryResponseBody
	GetBucketType() *string
	SetBuckets(v string) *GetLaboratoryResponseBody
	GetBuckets() *string
	SetCrowdId(v string) *GetLaboratoryResponseBody
	GetCrowdId() *string
	SetDebugCrowdId(v string) *GetLaboratoryResponseBody
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *GetLaboratoryResponseBody
	GetDebugUsers() *string
	SetDescription(v string) *GetLaboratoryResponseBody
	GetDescription() *string
	SetEnvironment(v string) *GetLaboratoryResponseBody
	GetEnvironment() *string
	SetFilter(v string) *GetLaboratoryResponseBody
	GetFilter() *string
	SetName(v string) *GetLaboratoryResponseBody
	GetName() *string
	SetRequestId(v string) *GetLaboratoryResponseBody
	GetRequestId() *string
	SetSceneId(v string) *GetLaboratoryResponseBody
	GetSceneId() *string
	SetStatus(v string) *GetLaboratoryResponseBody
	GetStatus() *string
	SetType(v string) *GetLaboratoryResponseBody
	GetType() *string
}

type GetLaboratoryResponseBody struct {
	// example:
	//
	// 100
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
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
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// user1,user2,user3
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
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1C0898E5-9220-5443-B2D9-445FF0688215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLaboratoryResponseBody) GetBucketCount() *int32 {
	return s.BucketCount
}

func (s *GetLaboratoryResponseBody) GetBucketType() *string {
	return s.BucketType
}

func (s *GetLaboratoryResponseBody) GetBuckets() *string {
	return s.Buckets
}

func (s *GetLaboratoryResponseBody) GetCrowdId() *string {
	return s.CrowdId
}

func (s *GetLaboratoryResponseBody) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *GetLaboratoryResponseBody) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *GetLaboratoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetLaboratoryResponseBody) GetEnvironment() *string {
	return s.Environment
}

func (s *GetLaboratoryResponseBody) GetFilter() *string {
	return s.Filter
}

func (s *GetLaboratoryResponseBody) GetName() *string {
	return s.Name
}

func (s *GetLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLaboratoryResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetLaboratoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetLaboratoryResponseBody) GetType() *string {
	return s.Type
}

func (s *GetLaboratoryResponseBody) SetBucketCount(v int32) *GetLaboratoryResponseBody {
	s.BucketCount = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetBucketType(v string) *GetLaboratoryResponseBody {
	s.BucketType = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetBuckets(v string) *GetLaboratoryResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetCrowdId(v string) *GetLaboratoryResponseBody {
	s.CrowdId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetDebugCrowdId(v string) *GetLaboratoryResponseBody {
	s.DebugCrowdId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetDebugUsers(v string) *GetLaboratoryResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetDescription(v string) *GetLaboratoryResponseBody {
	s.Description = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetEnvironment(v string) *GetLaboratoryResponseBody {
	s.Environment = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetFilter(v string) *GetLaboratoryResponseBody {
	s.Filter = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetName(v string) *GetLaboratoryResponseBody {
	s.Name = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetRequestId(v string) *GetLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetSceneId(v string) *GetLaboratoryResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetStatus(v string) *GetLaboratoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetType(v string) *GetLaboratoryResponseBody {
	s.Type = &v
	return s
}

func (s *GetLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
