// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLaboratoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaboratories(v []*ListLaboratoriesResponseBodyLaboratories) *ListLaboratoriesResponseBody
	GetLaboratories() []*ListLaboratoriesResponseBodyLaboratories
	SetRequestId(v string) *ListLaboratoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListLaboratoriesResponseBody
	GetTotalCount() *int64
}

type ListLaboratoriesResponseBody struct {
	Laboratories []*ListLaboratoriesResponseBodyLaboratories `json:"Laboratories,omitempty" xml:"Laboratories,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 1C0898E5-9220-5443-B2D9-445FF0688215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLaboratoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLaboratoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesResponseBody) GetLaboratories() []*ListLaboratoriesResponseBodyLaboratories {
	return s.Laboratories
}

func (s *ListLaboratoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLaboratoriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLaboratoriesResponseBody) SetLaboratories(v []*ListLaboratoriesResponseBodyLaboratories) *ListLaboratoriesResponseBody {
	s.Laboratories = v
	return s
}

func (s *ListLaboratoriesResponseBody) SetRequestId(v string) *ListLaboratoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLaboratoriesResponseBody) SetTotalCount(v int64) *ListLaboratoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLaboratoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLaboratoriesResponseBodyLaboratories struct {
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
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s ListLaboratoriesResponseBodyLaboratories) String() string {
	return dara.Prettify(s)
}

func (s ListLaboratoriesResponseBodyLaboratories) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetBucketCount() *int32 {
	return s.BucketCount
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetBucketType() *string {
	return s.BucketType
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetBuckets() *string {
	return s.Buckets
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetCrowdId() *string {
	return s.CrowdId
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetDescription() *string {
	return s.Description
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetEnvironment() *string {
	return s.Environment
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetFilter() *string {
	return s.Filter
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetName() *string {
	return s.Name
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetSceneId() *string {
	return s.SceneId
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetStatus() *string {
	return s.Status
}

func (s *ListLaboratoriesResponseBodyLaboratories) GetType() *string {
	return s.Type
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetBucketCount(v int32) *ListLaboratoriesResponseBodyLaboratories {
	s.BucketCount = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetBucketType(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.BucketType = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetBuckets(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Buckets = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetCrowdId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.CrowdId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetDebugCrowdId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.DebugCrowdId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetDebugUsers(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.DebugUsers = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetDescription(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Description = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetEnvironment(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Environment = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetFilter(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Filter = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetLaboratoryId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.LaboratoryId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetName(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Name = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetSceneId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.SceneId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetStatus(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Status = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetType(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Type = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) Validate() error {
	return dara.Validate(s)
}
