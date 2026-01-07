// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDataAgentWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddUserToDataAgentWorkspaceResponseBodyData) *AddUserToDataAgentWorkspaceResponseBody
	GetData() *AddUserToDataAgentWorkspaceResponseBodyData
	SetErrorCode(v string) *AddUserToDataAgentWorkspaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddUserToDataAgentWorkspaceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddUserToDataAgentWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddUserToDataAgentWorkspaceResponseBody
	GetSuccess() *bool
}

type AddUserToDataAgentWorkspaceResponseBody struct {
	Data *AddUserToDataAgentWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 67E910F2-***-695C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserToDataAgentWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDataAgentWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToDataAgentWorkspaceResponseBody) GetData() *AddUserToDataAgentWorkspaceResponseBodyData {
	return s.Data
}

func (s *AddUserToDataAgentWorkspaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddUserToDataAgentWorkspaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddUserToDataAgentWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToDataAgentWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserToDataAgentWorkspaceResponseBody) SetData(v *AddUserToDataAgentWorkspaceResponseBodyData) *AddUserToDataAgentWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBody) SetErrorCode(v string) *AddUserToDataAgentWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBody) SetErrorMessage(v string) *AddUserToDataAgentWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBody) SetRequestId(v string) *AddUserToDataAgentWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBody) SetSuccess(v bool) *AddUserToDataAgentWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddUserToDataAgentWorkspaceResponseBodyData struct {
	// example:
	//
	// 1765960516
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
	// 2
	RunningTaskNumber *int64 `json:"RunningTaskNumber,omitempty" xml:"RunningTaskNumber,omitempty"`
	// example:
	//
	// 5
	TotalTaskNumber *int64 `json:"TotalTaskNumber,omitempty" xml:"TotalTaskNumber,omitempty"`
	// example:
	//
	// agentTest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddUserToDataAgentWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDataAgentWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) GetMemberId() *string {
	return s.MemberId
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) GetRunningTaskNumber() *int64 {
	return s.RunningTaskNumber
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) GetTotalTaskNumber() *int64 {
	return s.TotalTaskNumber
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) SetJoinTime(v int64) *AddUserToDataAgentWorkspaceResponseBodyData {
	s.JoinTime = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) SetMemberId(v string) *AddUserToDataAgentWorkspaceResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) SetRoleName(v string) *AddUserToDataAgentWorkspaceResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) SetRunningTaskNumber(v int64) *AddUserToDataAgentWorkspaceResponseBodyData {
	s.RunningTaskNumber = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) SetTotalTaskNumber(v int64) *AddUserToDataAgentWorkspaceResponseBodyData {
	s.TotalTaskNumber = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) SetUserName(v string) *AddUserToDataAgentWorkspaceResponseBodyData {
	s.UserName = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
