// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetTopicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAssetTopicsResponseBody
	GetCode() *string
	SetData(v *ListAssetTopicsResponseBodyData) *ListAssetTopicsResponseBody
	GetData() *ListAssetTopicsResponseBodyData
	SetHttpStatusCode(v int32) *ListAssetTopicsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAssetTopicsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAssetTopicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAssetTopicsResponseBody
	GetSuccess() *bool
}

type ListAssetTopicsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paginated result of asset topics.
	Data *ListAssetTopicsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAssetTopicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAssetTopicsResponseBody) GetData() *ListAssetTopicsResponseBodyData {
	return s.Data
}

func (s *ListAssetTopicsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAssetTopicsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAssetTopicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetTopicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAssetTopicsResponseBody) SetCode(v string) *ListAssetTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAssetTopicsResponseBody) SetData(v *ListAssetTopicsResponseBodyData) *ListAssetTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListAssetTopicsResponseBody) SetHttpStatusCode(v int32) *ListAssetTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAssetTopicsResponseBody) SetMessage(v string) *ListAssetTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAssetTopicsResponseBody) SetRequestId(v string) *ListAssetTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetTopicsResponseBody) SetSuccess(v bool) *ListAssetTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAssetTopicsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAssetTopicsResponseBodyData struct {
	// The list of topics.
	TopicList []*ListAssetTopicsResponseBodyDataTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// -599403204152
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAssetTopicsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponseBodyData) GetTopicList() []*ListAssetTopicsResponseBodyDataTopicList {
	return s.TopicList
}

func (s *ListAssetTopicsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAssetTopicsResponseBodyData) SetTopicList(v []*ListAssetTopicsResponseBodyDataTopicList) *ListAssetTopicsResponseBodyData {
	s.TopicList = v
	return s
}

func (s *ListAssetTopicsResponseBodyData) SetTotalCount(v int64) *ListAssetTopicsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAssetTopicsResponseBodyData) Validate() error {
	if s.TopicList != nil {
		for _, item := range s.TopicList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssetTopicsResponseBodyDataTopicList struct {
	// The asset type.
	//
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The last modified time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The topic administrators.
	Owners []*ListAssetTopicsResponseBodyDataTopicListOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	// The topic description.
	//
	// example:
	//
	// Aggregates assets related to core metrics
	TopicDescription *string `json:"TopicDescription,omitempty" xml:"TopicDescription,omitempty"`
	// The topic ID.
	//
	// example:
	//
	// 28440278777
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Core Metrics Topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The visibility scope. Valid values: PUBLIC, SPECIFIED.
	//
	// example:
	//
	// SPECIFIED
	VisibilityType *string `json:"VisibilityType,omitempty" xml:"VisibilityType,omitempty"`
	// The explicitly visible user groups. Returns null for PUBLIC topics.
	VisibleUserGroups []*ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups `json:"VisibleUserGroups,omitempty" xml:"VisibleUserGroups,omitempty" type:"Repeated"`
	// The explicitly visible users. Returns null for PUBLIC topics.
	VisibleUsers []*ListAssetTopicsResponseBodyDataTopicListVisibleUsers `json:"VisibleUsers,omitempty" xml:"VisibleUsers,omitempty" type:"Repeated"`
}

func (s ListAssetTopicsResponseBodyDataTopicList) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponseBodyDataTopicList) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetAssetType() *string {
	return s.AssetType
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetOwners() []*ListAssetTopicsResponseBodyDataTopicListOwners {
	return s.Owners
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetTopicDescription() *string {
	return s.TopicDescription
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetTopicName() *string {
	return s.TopicName
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetVisibilityType() *string {
	return s.VisibilityType
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetVisibleUserGroups() []*ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups {
	return s.VisibleUserGroups
}

func (s *ListAssetTopicsResponseBodyDataTopicList) GetVisibleUsers() []*ListAssetTopicsResponseBodyDataTopicListVisibleUsers {
	return s.VisibleUsers
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetAssetType(v string) *ListAssetTopicsResponseBodyDataTopicList {
	s.AssetType = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetModifyTime(v string) *ListAssetTopicsResponseBodyDataTopicList {
	s.ModifyTime = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetOwners(v []*ListAssetTopicsResponseBodyDataTopicListOwners) *ListAssetTopicsResponseBodyDataTopicList {
	s.Owners = v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetTopicDescription(v string) *ListAssetTopicsResponseBodyDataTopicList {
	s.TopicDescription = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetTopicId(v int64) *ListAssetTopicsResponseBodyDataTopicList {
	s.TopicId = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetTopicName(v string) *ListAssetTopicsResponseBodyDataTopicList {
	s.TopicName = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetVisibilityType(v string) *ListAssetTopicsResponseBodyDataTopicList {
	s.VisibilityType = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetVisibleUserGroups(v []*ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) *ListAssetTopicsResponseBodyDataTopicList {
	s.VisibleUserGroups = v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) SetVisibleUsers(v []*ListAssetTopicsResponseBodyDataTopicListVisibleUsers) *ListAssetTopicsResponseBodyDataTopicList {
	s.VisibleUsers = v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicList) Validate() error {
	if s.Owners != nil {
		for _, item := range s.Owners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VisibleUserGroups != nil {
		for _, item := range s.VisibleUserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VisibleUsers != nil {
		for _, item := range s.VisibleUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssetTopicsResponseBodyDataTopicListOwners struct {
	// The user ID.
	//
	// example:
	//
	// 30001011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// John
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListAssetTopicsResponseBodyDataTopicListOwners) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponseBodyDataTopicListOwners) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponseBodyDataTopicListOwners) GetUserId() *string {
	return s.UserId
}

func (s *ListAssetTopicsResponseBodyDataTopicListOwners) GetUserName() *string {
	return s.UserName
}

func (s *ListAssetTopicsResponseBodyDataTopicListOwners) SetUserId(v string) *ListAssetTopicsResponseBodyDataTopicListOwners {
	s.UserId = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicListOwners) SetUserName(v string) *ListAssetTopicsResponseBodyDataTopicListOwners {
	s.UserName = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicListOwners) Validate() error {
	return dara.Validate(s)
}

type ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups struct {
	// The user group ID.
	//
	// example:
	//
	// 20001
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// Data Governance Group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) SetUserGroupId(v string) *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) SetUserGroupName(v string) *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups {
	s.UserGroupName = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUserGroups) Validate() error {
	return dara.Validate(s)
}

type ListAssetTopicsResponseBodyDataTopicListVisibleUsers struct {
	// The user ID.
	//
	// example:
	//
	// 30001012
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// Jane
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListAssetTopicsResponseBodyDataTopicListVisibleUsers) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponseBodyDataTopicListVisibleUsers) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUsers) GetUserName() *string {
	return s.UserName
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUsers) SetUserId(v string) *ListAssetTopicsResponseBodyDataTopicListVisibleUsers {
	s.UserId = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUsers) SetUserName(v string) *ListAssetTopicsResponseBodyDataTopicListVisibleUsers {
	s.UserName = &v
	return s
}

func (s *ListAssetTopicsResponseBodyDataTopicListVisibleUsers) Validate() error {
	return dara.Validate(s)
}
