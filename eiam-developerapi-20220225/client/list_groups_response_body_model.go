// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListGroupsResponseBodyData) *ListGroupsResponseBody
	GetData() []*ListGroupsResponseBodyData
	SetMaxResults(v int32) *ListGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsResponseBody
	GetNextToken() *string
	SetTotalCount(v int64) *ListGroupsResponseBody
	GetTotalCount() *int64
}

type ListGroupsResponseBody struct {
	// The returned data.
	Data []*ListGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The start position of the query. If this parameter is left empty, the query starts from the beginning.
	//
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetData() []*ListGroupsResponseBodyData {
	return s.Data
}

func (s *ListGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsResponseBody) SetData(v []*ListGroupsResponseBodyData) *ListGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsResponseBody) SetMaxResults(v int32) *ListGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsResponseBody) SetNextToken(v string) *ListGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int64) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsResponseBodyData struct {
	// The time when the group was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The group description.
	//
	// example:
	//
	// description_demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The external ID of the group.
	//
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupExternalId *string `json:"groupExternalId,omitempty" xml:"groupExternalId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The source ID of the group.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"groupSourceId,omitempty" xml:"groupSourceId,omitempty"`
	// The source type of the group. Valid values: build_in, ding_talk, ad, and ldap.
	//
	// example:
	//
	// build_in
	GroupSourceType *string `json:"groupSourceType,omitempty" xml:"groupSourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The time when the group was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListGroupsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListGroupsResponseBodyData) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *ListGroupsResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsResponseBodyData) GetGroupSourceId() *string {
	return s.GroupSourceId
}

func (s *ListGroupsResponseBodyData) GetGroupSourceType() *string {
	return s.GroupSourceType
}

func (s *ListGroupsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListGroupsResponseBodyData) SetCreateTime(v int64) *ListGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetDescription(v string) *ListGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupExternalId(v string) *ListGroupsResponseBodyData {
	s.GroupExternalId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupId(v string) *ListGroupsResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupName(v string) *ListGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupSourceId(v string) *ListGroupsResponseBodyData {
	s.GroupSourceId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupSourceType(v string) *ListGroupsResponseBodyData {
	s.GroupSourceType = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetInstanceId(v string) *ListGroupsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetUpdateTime(v int64) *ListGroupsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
