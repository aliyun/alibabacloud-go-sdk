// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserGroupInfoResponseBody
	GetRequestId() *string
	SetResult(v []*GetUserGroupInfoResponseBodyResult) *GetUserGroupInfoResponseBody
	GetResult() []*GetUserGroupInfoResponseBodyResult
	SetSuccess(v bool) *GetUserGroupInfoResponseBody
	GetSuccess() *bool
}

type GetUserGroupInfoResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D7980306-1F08-5A88-9FE7-ECB8B9C4C0F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// User group information.
	Result []*GetUserGroupInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s GetUserGroupInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserGroupInfoResponseBody) GetResult() []*GetUserGroupInfoResponseBodyResult {
	return s.Result
}

func (s *GetUserGroupInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserGroupInfoResponseBody) SetRequestId(v string) *GetUserGroupInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupInfoResponseBody) SetResult(v []*GetUserGroupInfoResponseBodyResult) *GetUserGroupInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetUserGroupInfoResponseBody) SetSuccess(v bool) *GetUserGroupInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserGroupInfoResponseBody) Validate() error {
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

type GetUserGroupInfoResponseBodyResult struct {
	// Creation time of the user group.
	//
	// example:
	//
	// 2021-03-15 17:13:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator of the sub-user group. This is the UserID in Quick BI, not the UID in Alibaba Cloud.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Directory level of the user group.
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// Last modified time of the user group.
	//
	// example:
	//
	// 2021-03-15 20:36:40
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Modifier of the user group. This is the UserID in Quick BI, not the UID in Alibaba Cloud.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// Parent user group ID.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	// Description of the user group.
	//
	// example:
	//
	// test
	UsergroupDesc *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	// User group ID.
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****
	UsergroupId *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	// Name of the user group.
	//
	// example:
	//
	// test
	UsergroupName *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
}

func (s GetUserGroupInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserGroupInfoResponseBodyResult) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetUserGroupInfoResponseBodyResult) GetIdentifiedPath() *string {
	return s.IdentifiedPath
}

func (s *GetUserGroupInfoResponseBodyResult) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetUserGroupInfoResponseBodyResult) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetUserGroupInfoResponseBodyResult) GetParentUsergroupId() *string {
	return s.ParentUsergroupId
}

func (s *GetUserGroupInfoResponseBodyResult) GetUsergroupDesc() *string {
	return s.UsergroupDesc
}

func (s *GetUserGroupInfoResponseBodyResult) GetUsergroupId() *string {
	return s.UsergroupId
}

func (s *GetUserGroupInfoResponseBodyResult) GetUsergroupName() *string {
	return s.UsergroupName
}

func (s *GetUserGroupInfoResponseBodyResult) SetCreateTime(v string) *GetUserGroupInfoResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetCreateUser(v string) *GetUserGroupInfoResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetIdentifiedPath(v string) *GetUserGroupInfoResponseBodyResult {
	s.IdentifiedPath = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetModifiedTime(v string) *GetUserGroupInfoResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetModifyUser(v string) *GetUserGroupInfoResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetParentUsergroupId(v string) *GetUserGroupInfoResponseBodyResult {
	s.ParentUsergroupId = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetUsergroupDesc(v string) *GetUserGroupInfoResponseBodyResult {
	s.UsergroupDesc = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetUsergroupId(v string) *GetUserGroupInfoResponseBodyResult {
	s.UsergroupId = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetUsergroupName(v string) *GetUserGroupInfoResponseBodyResult {
	s.UsergroupName = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
