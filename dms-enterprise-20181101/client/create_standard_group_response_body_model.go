// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateStandardGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateStandardGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateStandardGroupResponseBody
	GetRequestId() *string
	SetStandardGroup(v *CreateStandardGroupResponseBodyStandardGroup) *CreateStandardGroupResponseBody
	GetStandardGroup() *CreateStandardGroupResponseBodyStandardGroup
	SetSuccess(v bool) *CreateStandardGroupResponseBody
	GetSuccess() *bool
}

type CreateStandardGroupResponseBody struct {
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
	// The information about the created security rule set.
	StandardGroup *CreateStandardGroupResponseBodyStandardGroup `json:"StandardGroup,omitempty" xml:"StandardGroup,omitempty" type:"Struct"`
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

func (s CreateStandardGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateStandardGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateStandardGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardGroupResponseBody) GetStandardGroup() *CreateStandardGroupResponseBodyStandardGroup {
	return s.StandardGroup
}

func (s *CreateStandardGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardGroupResponseBody) SetErrorCode(v string) *CreateStandardGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateStandardGroupResponseBody) SetErrorMessage(v string) *CreateStandardGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStandardGroupResponseBody) SetRequestId(v string) *CreateStandardGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardGroupResponseBody) SetStandardGroup(v *CreateStandardGroupResponseBodyStandardGroup) *CreateStandardGroupResponseBody {
	s.StandardGroup = v
	return s
}

func (s *CreateStandardGroupResponseBody) SetSuccess(v bool) *CreateStandardGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateStandardGroupResponseBodyStandardGroup struct {
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
	// The security rule set ID.
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
	// The ID of the user who creates the security rule set.
	//
	// example:
	//
	// 51****
	LastMenderId *int64 `json:"LastMenderId,omitempty" xml:"LastMenderId,omitempty"`
}

func (s CreateStandardGroupResponseBodyStandardGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardGroupResponseBodyStandardGroup) GoString() string {
	return s.String()
}

func (s *CreateStandardGroupResponseBodyStandardGroup) GetDbType() *string {
	return s.DbType
}

func (s *CreateStandardGroupResponseBodyStandardGroup) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardGroupResponseBodyStandardGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateStandardGroupResponseBodyStandardGroup) GetGroupMode() *string {
	return s.GroupMode
}

func (s *CreateStandardGroupResponseBodyStandardGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateStandardGroupResponseBodyStandardGroup) GetLastMenderId() *int64 {
	return s.LastMenderId
}

func (s *CreateStandardGroupResponseBodyStandardGroup) SetDbType(v string) *CreateStandardGroupResponseBodyStandardGroup {
	s.DbType = &v
	return s
}

func (s *CreateStandardGroupResponseBodyStandardGroup) SetDescription(v string) *CreateStandardGroupResponseBodyStandardGroup {
	s.Description = &v
	return s
}

func (s *CreateStandardGroupResponseBodyStandardGroup) SetGroupId(v int64) *CreateStandardGroupResponseBodyStandardGroup {
	s.GroupId = &v
	return s
}

func (s *CreateStandardGroupResponseBodyStandardGroup) SetGroupMode(v string) *CreateStandardGroupResponseBodyStandardGroup {
	s.GroupMode = &v
	return s
}

func (s *CreateStandardGroupResponseBodyStandardGroup) SetGroupName(v string) *CreateStandardGroupResponseBodyStandardGroup {
	s.GroupName = &v
	return s
}

func (s *CreateStandardGroupResponseBodyStandardGroup) SetLastMenderId(v int64) *CreateStandardGroupResponseBodyStandardGroup {
	s.LastMenderId = &v
	return s
}

func (s *CreateStandardGroupResponseBodyStandardGroup) Validate() error {
	return dara.Validate(s)
}
