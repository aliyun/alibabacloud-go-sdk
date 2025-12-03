// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsMappingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *ListUserGroupsMappingsResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *ListUserGroupsMappingsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserGroupsMappingsResponseBody
	GetRequestId() *string
	SetUserGroupsMappings(v []*ListUserGroupsMappingsResponseBodyUserGroupsMappings) *ListUserGroupsMappingsResponseBody
	GetUserGroupsMappings() []*ListUserGroupsMappingsResponseBodyUserGroupsMappings
}

type ListUserGroupsMappingsResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// user1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroupsMappings []*ListUserGroupsMappingsResponseBodyUserGroupsMappings `json:"UserGroupsMappings,omitempty" xml:"UserGroupsMappings,omitempty" type:"Repeated"`
}

func (s ListUserGroupsMappingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListUserGroupsMappingsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserGroupsMappingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsMappingsResponseBody) GetUserGroupsMappings() []*ListUserGroupsMappingsResponseBodyUserGroupsMappings {
	return s.UserGroupsMappings
}

func (s *ListUserGroupsMappingsResponseBody) SetHasMore(v bool) *ListUserGroupsMappingsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) SetNextToken(v string) *ListUserGroupsMappingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) SetRequestId(v string) *ListUserGroupsMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) SetUserGroupsMappings(v []*ListUserGroupsMappingsResponseBodyUserGroupsMappings) *ListUserGroupsMappingsResponseBody {
	s.UserGroupsMappings = v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) Validate() error {
	if s.UserGroupsMappings != nil {
		for _, item := range s.UserGroupsMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsMappingsResponseBodyUserGroupsMappings struct {
	GroupNames []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUserGroupsMappingsResponseBodyUserGroupsMappings) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsMappingsResponseBodyUserGroupsMappings) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) GetGroupNames() []*string {
	return s.GroupNames
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) GetUserName() *string {
	return s.UserName
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) SetGroupNames(v []*string) *ListUserGroupsMappingsResponseBodyUserGroupsMappings {
	s.GroupNames = v
	return s
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) SetUserName(v string) *ListUserGroupsMappingsResponseBodyUserGroupsMappings {
	s.UserName = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) Validate() error {
	return dara.Validate(s)
}
