// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateStandardGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateStandardGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateStandardGroupResponseBody
	GetRequestId() *string
	SetStandardGroup(v *UpdateStandardGroupResponseBodyStandardGroup) *UpdateStandardGroupResponseBody
	GetStandardGroup() *UpdateStandardGroupResponseBodyStandardGroup
	SetSuccess(v bool) *UpdateStandardGroupResponseBody
	GetSuccess() *bool
}

type UpdateStandardGroupResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the security rule set.
	StandardGroup *UpdateStandardGroupResponseBodyStandardGroup `json:"StandardGroup,omitempty" xml:"StandardGroup,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateStandardGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateStandardGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateStandardGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardGroupResponseBody) GetStandardGroup() *UpdateStandardGroupResponseBodyStandardGroup {
	return s.StandardGroup
}

func (s *UpdateStandardGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardGroupResponseBody) SetErrorCode(v string) *UpdateStandardGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateStandardGroupResponseBody) SetErrorMessage(v string) *UpdateStandardGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateStandardGroupResponseBody) SetRequestId(v string) *UpdateStandardGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardGroupResponseBody) SetStandardGroup(v *UpdateStandardGroupResponseBodyStandardGroup) *UpdateStandardGroupResponseBody {
	s.StandardGroup = v
	return s
}

func (s *UpdateStandardGroupResponseBody) SetSuccess(v bool) *UpdateStandardGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardGroupResponseBody) Validate() error {
	if s.StandardGroup != nil {
		if err := s.StandardGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardGroupResponseBodyStandardGroup struct {
	// The type of the database for which the security rules are used.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the security rule set.
	//
	// example:
	//
	// Production Environment test rules
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
	// poc_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the user who last modified the security rules.
	//
	// example:
	//
	// 51****
	LastMenderId *int64 `json:"LastMenderId,omitempty" xml:"LastMenderId,omitempty"`
}

func (s UpdateStandardGroupResponseBodyStandardGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardGroupResponseBodyStandardGroup) GoString() string {
	return s.String()
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) GetDbType() *string {
	return s.DbType
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) GetGroupMode() *string {
	return s.GroupMode
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) GetLastMenderId() *int64 {
	return s.LastMenderId
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) SetDbType(v string) *UpdateStandardGroupResponseBodyStandardGroup {
	s.DbType = &v
	return s
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) SetDescription(v string) *UpdateStandardGroupResponseBodyStandardGroup {
	s.Description = &v
	return s
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) SetGroupId(v int64) *UpdateStandardGroupResponseBodyStandardGroup {
	s.GroupId = &v
	return s
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) SetGroupMode(v string) *UpdateStandardGroupResponseBodyStandardGroup {
	s.GroupMode = &v
	return s
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) SetGroupName(v string) *UpdateStandardGroupResponseBodyStandardGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) SetLastMenderId(v int64) *UpdateStandardGroupResponseBodyStandardGroup {
	s.LastMenderId = &v
	return s
}

func (s *UpdateStandardGroupResponseBodyStandardGroup) Validate() error {
	return dara.Validate(s)
}
