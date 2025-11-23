// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListStandardGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListStandardGroupsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListStandardGroupsResponseBody
	GetRequestId() *string
	SetStandardGroupList(v []*ListStandardGroupsResponseBodyStandardGroupList) *ListStandardGroupsResponseBody
	GetStandardGroupList() []*ListStandardGroupsResponseBodyStandardGroupList
	SetSuccess(v bool) *ListStandardGroupsResponseBody
	GetSuccess() *bool
}

type ListStandardGroupsResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security rule sets returned.
	StandardGroupList []*ListStandardGroupsResponseBodyStandardGroupList `json:"StandardGroupList,omitempty" xml:"StandardGroupList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListStandardGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStandardGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStandardGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStandardGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListStandardGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStandardGroupsResponseBody) GetStandardGroupList() []*ListStandardGroupsResponseBodyStandardGroupList {
	return s.StandardGroupList
}

func (s *ListStandardGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStandardGroupsResponseBody) SetErrorCode(v string) *ListStandardGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStandardGroupsResponseBody) SetErrorMessage(v string) *ListStandardGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListStandardGroupsResponseBody) SetRequestId(v string) *ListStandardGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStandardGroupsResponseBody) SetStandardGroupList(v []*ListStandardGroupsResponseBodyStandardGroupList) *ListStandardGroupsResponseBody {
	s.StandardGroupList = v
	return s
}

func (s *ListStandardGroupsResponseBody) SetSuccess(v bool) *ListStandardGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStandardGroupsResponseBody) Validate() error {
	if s.StandardGroupList != nil {
		for _, item := range s.StandardGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStandardGroupsResponseBodyStandardGroupList struct {
	// The type of the database engine. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the security rule set.
	//
	// example:
	//
	// test_rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the security rule set.
	//
	// example:
	//
	// 41****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The control mode. Valid values:
	//
	// 	- **NONE_CONTROL**: Flexible Management
	//
	// 	- **STABLE**: Stable Change
	//
	// 	- **COMMON**: Security Collaboration
	//
	// example:
	//
	// COMMON
	GroupMode *string `json:"GroupMode,omitempty" xml:"GroupMode,omitempty"`
	// The name of the security rule set.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the user who queries the security sets.
	//
	// example:
	//
	// 51****
	LastMenderId *int64 `json:"LastMenderId,omitempty" xml:"LastMenderId,omitempty"`
}

func (s ListStandardGroupsResponseBodyStandardGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListStandardGroupsResponseBodyStandardGroupList) GoString() string {
	return s.String()
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) GetDbType() *string {
	return s.DbType
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) GetDescription() *string {
	return s.Description
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) GetGroupMode() *string {
	return s.GroupMode
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) GetLastMenderId() *int64 {
	return s.LastMenderId
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) SetDbType(v string) *ListStandardGroupsResponseBodyStandardGroupList {
	s.DbType = &v
	return s
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) SetDescription(v string) *ListStandardGroupsResponseBodyStandardGroupList {
	s.Description = &v
	return s
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) SetGroupId(v int64) *ListStandardGroupsResponseBodyStandardGroupList {
	s.GroupId = &v
	return s
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) SetGroupMode(v string) *ListStandardGroupsResponseBodyStandardGroupList {
	s.GroupMode = &v
	return s
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) SetGroupName(v string) *ListStandardGroupsResponseBodyStandardGroupList {
	s.GroupName = &v
	return s
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) SetLastMenderId(v int64) *ListStandardGroupsResponseBodyStandardGroupList {
	s.LastMenderId = &v
	return s
}

func (s *ListStandardGroupsResponseBodyStandardGroupList) Validate() error {
	return dara.Validate(s)
}
