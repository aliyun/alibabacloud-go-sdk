// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentWorkspaceMemberRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) *UpdateDataAgentWorkspaceMemberRoleResponseBody
	GetData() *UpdateDataAgentWorkspaceMemberRoleResponseBodyData
	SetErrorCode(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataAgentWorkspaceMemberRoleResponseBody
	GetSuccess() *bool
}

type UpdateDataAgentWorkspaceMemberRoleResponseBody struct {
	Data *UpdateDataAgentWorkspaceMemberRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E0D2-*****-A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataAgentWorkspaceMemberRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentWorkspaceMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) GetData() *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	return s.Data
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) SetData(v *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) *UpdateDataAgentWorkspaceMemberRoleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) SetErrorCode(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) SetErrorMessage(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) SetRequestId(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) SetSuccess(v bool) *UpdateDataAgentWorkspaceMemberRoleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataAgentWorkspaceMemberRoleResponseBodyData struct {
	// example:
	//
	// 1765961516
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// example:
	//
	// 20282*****7591
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 10
	RunningTaskNumber *int64 `json:"RunningTaskNumber,omitempty" xml:"RunningTaskNumber,omitempty"`
	// example:
	//
	// 20
	TotalTaskNumber *int64 `json:"TotalTaskNumber,omitempty" xml:"TotalTaskNumber,omitempty"`
	// example:
	//
	// yunqitest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateDataAgentWorkspaceMemberRoleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GetRunningTaskNumber() *int64 {
	return s.RunningTaskNumber
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GetTotalTaskNumber() *int64 {
	return s.TotalTaskNumber
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) SetJoinTime(v int64) *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	s.JoinTime = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) SetMemberId(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) SetRoleName(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) SetRunningTaskNumber(v int64) *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	s.RunningTaskNumber = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) SetTotalTaskNumber(v int64) *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	s.TotalTaskNumber = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) SetUserName(v string) *UpdateDataAgentWorkspaceMemberRoleResponseBodyData {
	s.UserName = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
