// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataAgentWorkspaceResponseBodyData) *CreateDataAgentWorkspaceResponseBody
	GetData() *CreateDataAgentWorkspaceResponseBodyData
	SetErrorCode(v string) *CreateDataAgentWorkspaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataAgentWorkspaceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataAgentWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAgentWorkspaceResponseBody
	GetSuccess() *bool
}

type CreateDataAgentWorkspaceResponseBody struct {
	Data *CreateDataAgentWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateDataAgentWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentWorkspaceResponseBody) GetData() *CreateDataAgentWorkspaceResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentWorkspaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentWorkspaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataAgentWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAgentWorkspaceResponseBody) SetData(v *CreateDataAgentWorkspaceResponseBodyData) *CreateDataAgentWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBody) SetErrorCode(v string) *CreateDataAgentWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBody) SetErrorMessage(v string) *CreateDataAgentWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBody) SetRequestId(v string) *CreateDataAgentWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBody) SetSuccess(v bool) *CreateDataAgentWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentWorkspaceResponseBodyData struct {
	// example:
	//
	// 1765960516
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 20282*****7591
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1765961516
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// owner
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 11
	TotalMember *int64 `json:"TotalMember,omitempty" xml:"TotalMember,omitempty"`
	// example:
	//
	// space for test
	WorkspaceDesc *string `json:"WorkspaceDesc,omitempty" xml:"WorkspaceDesc,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// workspaceTest
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// example:
	//
	// active
	WorkspaceStatus *string `json:"WorkspaceStatus,omitempty" xml:"WorkspaceStatus,omitempty"`
}

func (s CreateDataAgentWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetTotalMember() *int64 {
	return s.TotalMember
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetWorkspaceDesc() *string {
	return s.WorkspaceDesc
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateDataAgentWorkspaceResponseBodyData) GetWorkspaceStatus() *string {
	return s.WorkspaceStatus
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetCreateTime(v int64) *CreateDataAgentWorkspaceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetCreator(v string) *CreateDataAgentWorkspaceResponseBodyData {
	s.Creator = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetModifyTime(v int64) *CreateDataAgentWorkspaceResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetRoleName(v string) *CreateDataAgentWorkspaceResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetTotalMember(v int64) *CreateDataAgentWorkspaceResponseBodyData {
	s.TotalMember = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetWorkspaceDesc(v string) *CreateDataAgentWorkspaceResponseBodyData {
	s.WorkspaceDesc = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetWorkspaceId(v string) *CreateDataAgentWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetWorkspaceName(v string) *CreateDataAgentWorkspaceResponseBodyData {
	s.WorkspaceName = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) SetWorkspaceStatus(v string) *CreateDataAgentWorkspaceResponseBodyData {
	s.WorkspaceStatus = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
