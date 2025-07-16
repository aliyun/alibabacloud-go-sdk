// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureEntities(v []*ListFeatureEntitiesResponseBodyFeatureEntities) *ListFeatureEntitiesResponseBody
	GetFeatureEntities() []*ListFeatureEntitiesResponseBodyFeatureEntities
	SetRequestId(v string) *ListFeatureEntitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFeatureEntitiesResponseBody
	GetTotalCount() *int32
}

type ListFeatureEntitiesResponseBody struct {
	FeatureEntities []*ListFeatureEntitiesResponseBodyFeatureEntities `json:"FeatureEntities,omitempty" xml:"FeatureEntities,omitempty" type:"Repeated"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesResponseBody) GetFeatureEntities() []*ListFeatureEntitiesResponseBodyFeatureEntities {
	return s.FeatureEntities
}

func (s *ListFeatureEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureEntitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFeatureEntitiesResponseBody) SetFeatureEntities(v []*ListFeatureEntitiesResponseBodyFeatureEntities) *ListFeatureEntitiesResponseBody {
	s.FeatureEntities = v
	return s
}

func (s *ListFeatureEntitiesResponseBody) SetRequestId(v string) *ListFeatureEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBody) SetTotalCount(v int32) *ListFeatureEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFeatureEntitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFeatureEntitiesResponseBodyFeatureEntities struct {
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456789****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListFeatureEntitiesResponseBodyFeatureEntities) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureEntitiesResponseBodyFeatureEntities) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetFeatureEntityId() *string {
	return s.FeatureEntityId
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetJoinId() *string {
	return s.JoinId
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetName() *string {
	return s.Name
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetOwner() *string {
	return s.Owner
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetFeatureEntityId(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.FeatureEntityId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetGmtCreateTime(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.GmtCreateTime = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetJoinId(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.JoinId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetName(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.Name = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetOwner(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.Owner = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetProjectId(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetProjectName(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.ProjectName = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) Validate() error {
	return dara.Validate(s)
}
