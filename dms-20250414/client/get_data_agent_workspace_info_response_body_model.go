// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentWorkspaceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataAgentWorkspaceInfoResponseBodyData) *GetDataAgentWorkspaceInfoResponseBody
	GetData() *GetDataAgentWorkspaceInfoResponseBodyData
	SetErrorCode(v string) *GetDataAgentWorkspaceInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataAgentWorkspaceInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataAgentWorkspaceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataAgentWorkspaceInfoResponseBody
	GetSuccess() *bool
}

type GetDataAgentWorkspaceInfoResponseBody struct {
	Data *GetDataAgentWorkspaceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetDataAgentWorkspaceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentWorkspaceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAgentWorkspaceInfoResponseBody) GetData() *GetDataAgentWorkspaceInfoResponseBodyData {
	return s.Data
}

func (s *GetDataAgentWorkspaceInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataAgentWorkspaceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataAgentWorkspaceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAgentWorkspaceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataAgentWorkspaceInfoResponseBody) SetData(v *GetDataAgentWorkspaceInfoResponseBodyData) *GetDataAgentWorkspaceInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBody) SetErrorCode(v string) *GetDataAgentWorkspaceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBody) SetErrorMessage(v string) *GetDataAgentWorkspaceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBody) SetRequestId(v string) *GetDataAgentWorkspaceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBody) SetSuccess(v bool) *GetDataAgentWorkspaceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataAgentWorkspaceInfoResponseBodyData struct {
	// example:
	//
	// 1765960516
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 20282*****7591
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1765961516
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 11
	TotalMember *string `json:"TotalMember,omitempty" xml:"TotalMember,omitempty"`
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

func (s GetDataAgentWorkspaceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentWorkspaceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetTotalMember() *string {
	return s.TotalMember
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetWorkspaceDesc() *string {
	return s.WorkspaceDesc
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) GetWorkspaceStatus() *string {
	return s.WorkspaceStatus
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetCreateTime(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetCreator(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetModifyTime(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetRoleName(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetTotalMember(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.TotalMember = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetWorkspaceDesc(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.WorkspaceDesc = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetWorkspaceId(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetWorkspaceName(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.WorkspaceName = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) SetWorkspaceStatus(v string) *GetDataAgentWorkspaceInfoResponseBodyData {
	s.WorkspaceStatus = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
