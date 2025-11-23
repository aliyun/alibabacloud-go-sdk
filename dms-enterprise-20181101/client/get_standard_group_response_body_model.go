// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetStandardGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetStandardGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetStandardGroupResponseBody
	GetRequestId() *string
	SetStandardGroup(v *GetStandardGroupResponseBodyStandardGroup) *GetStandardGroupResponseBody
	GetStandardGroup() *GetStandardGroupResponseBodyStandardGroup
	SetSuccess(v bool) *GetStandardGroupResponseBody
	GetSuccess() *bool
}

type GetStandardGroupResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// BF7E9543-F431-566A-991A-B5C493EA36C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the security rule set.
	StandardGroup *GetStandardGroupResponseBodyStandardGroup `json:"StandardGroup,omitempty" xml:"StandardGroup,omitempty" type:"Struct"`
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

func (s GetStandardGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStandardGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStandardGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardGroupResponseBody) GetStandardGroup() *GetStandardGroupResponseBodyStandardGroup {
	return s.StandardGroup
}

func (s *GetStandardGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardGroupResponseBody) SetErrorCode(v string) *GetStandardGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStandardGroupResponseBody) SetErrorMessage(v string) *GetStandardGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStandardGroupResponseBody) SetRequestId(v string) *GetStandardGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardGroupResponseBody) SetStandardGroup(v *GetStandardGroupResponseBodyStandardGroup) *GetStandardGroupResponseBody {
	s.StandardGroup = v
	return s
}

func (s *GetStandardGroupResponseBody) SetSuccess(v bool) *GetStandardGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardGroupResponseBody) Validate() error {
	if s.StandardGroup != nil {
		if err := s.StandardGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardGroupResponseBodyStandardGroup struct {
	// The engine type.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the security rule set.
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
	// poc_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the user who last modified the security rules.
	//
	// example:
	//
	// 51****
	LastMenderId *int64 `json:"LastMenderId,omitempty" xml:"LastMenderId,omitempty"`
}

func (s GetStandardGroupResponseBodyStandardGroup) String() string {
	return dara.Prettify(s)
}

func (s GetStandardGroupResponseBodyStandardGroup) GoString() string {
	return s.String()
}

func (s *GetStandardGroupResponseBodyStandardGroup) GetDbType() *string {
	return s.DbType
}

func (s *GetStandardGroupResponseBodyStandardGroup) GetDescription() *string {
	return s.Description
}

func (s *GetStandardGroupResponseBodyStandardGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetStandardGroupResponseBodyStandardGroup) GetGroupMode() *string {
	return s.GroupMode
}

func (s *GetStandardGroupResponseBodyStandardGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetStandardGroupResponseBodyStandardGroup) GetLastMenderId() *int64 {
	return s.LastMenderId
}

func (s *GetStandardGroupResponseBodyStandardGroup) SetDbType(v string) *GetStandardGroupResponseBodyStandardGroup {
	s.DbType = &v
	return s
}

func (s *GetStandardGroupResponseBodyStandardGroup) SetDescription(v string) *GetStandardGroupResponseBodyStandardGroup {
	s.Description = &v
	return s
}

func (s *GetStandardGroupResponseBodyStandardGroup) SetGroupId(v int64) *GetStandardGroupResponseBodyStandardGroup {
	s.GroupId = &v
	return s
}

func (s *GetStandardGroupResponseBodyStandardGroup) SetGroupMode(v string) *GetStandardGroupResponseBodyStandardGroup {
	s.GroupMode = &v
	return s
}

func (s *GetStandardGroupResponseBodyStandardGroup) SetGroupName(v string) *GetStandardGroupResponseBodyStandardGroup {
	s.GroupName = &v
	return s
}

func (s *GetStandardGroupResponseBodyStandardGroup) SetLastMenderId(v int64) *GetStandardGroupResponseBodyStandardGroup {
	s.LastMenderId = &v
	return s
}

func (s *GetStandardGroupResponseBodyStandardGroup) Validate() error {
	return dara.Validate(s)
}
