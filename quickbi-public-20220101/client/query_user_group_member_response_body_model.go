// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v []*QueryUserGroupMemberResponseBodyResult) *QueryUserGroupMemberResponseBody
	GetResult() []*QueryUserGroupMemberResponseBodyResult
	SetSuccess(v bool) *QueryUserGroupMemberResponseBody
	GetSuccess() *bool
}

type QueryUserGroupMemberResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 48C930FF-DFCF-5986-902B-E24C202E2443
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request for the user group member list.
	Result []*QueryUserGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserGroupMemberResponseBody) GetResult() []*QueryUserGroupMemberResponseBodyResult {
	return s.Result
}

func (s *QueryUserGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserGroupMemberResponseBody) SetRequestId(v string) *QueryUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserGroupMemberResponseBody) SetResult(v []*QueryUserGroupMemberResponseBodyResult) *QueryUserGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserGroupMemberResponseBody) SetSuccess(v bool) *QueryUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserGroupMemberResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserGroupMemberResponseBodyResult struct {
	// ID of the user group or the user group member.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether it is a user group. Possible values:
	//
	// - true: It is a user group.
	//
	// - false: It is a user.
	//
	// example:
	//
	// true
	IsUserGroup *bool `json:"IsUserGroup,omitempty" xml:"IsUserGroup,omitempty"`
	// Name or nickname of the user group or its member.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// ID of the parent user group.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	// Name of the parent user group.
	//
	// example:
	//
	// test
	ParentUserGroupName *string `json:"ParentUserGroupName,omitempty" xml:"ParentUserGroupName,omitempty"`
}

func (s QueryUserGroupMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *QueryUserGroupMemberResponseBodyResult) GetIsUserGroup() *bool {
	return s.IsUserGroup
}

func (s *QueryUserGroupMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *QueryUserGroupMemberResponseBodyResult) GetParentUserGroupId() *string {
	return s.ParentUserGroupId
}

func (s *QueryUserGroupMemberResponseBodyResult) GetParentUserGroupName() *string {
	return s.ParentUserGroupName
}

func (s *QueryUserGroupMemberResponseBodyResult) SetId(v string) *QueryUserGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetIsUserGroup(v bool) *QueryUserGroupMemberResponseBodyResult {
	s.IsUserGroup = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetName(v string) *QueryUserGroupMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetParentUserGroupId(v string) *QueryUserGroupMemberResponseBodyResult {
	s.ParentUserGroupId = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetParentUserGroupName(v string) *QueryUserGroupMemberResponseBodyResult {
	s.ParentUserGroupName = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
