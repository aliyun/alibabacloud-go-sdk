// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentSpaceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDataAgentSpaceInfoResponseBodyData) *UpdateDataAgentSpaceInfoResponseBody
	GetData() *UpdateDataAgentSpaceInfoResponseBodyData
	SetErrorCode(v string) *UpdateDataAgentSpaceInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataAgentSpaceInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataAgentSpaceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateDataAgentSpaceInfoResponseBody
	GetSuccess() *string
}

type UpdateDataAgentSpaceInfoResponseBody struct {
	Data *UpdateDataAgentSpaceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataAgentSpaceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentSpaceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentSpaceInfoResponseBody) GetData() *UpdateDataAgentSpaceInfoResponseBodyData {
	return s.Data
}

func (s *UpdateDataAgentSpaceInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataAgentSpaceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataAgentSpaceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataAgentSpaceInfoResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateDataAgentSpaceInfoResponseBody) SetData(v *UpdateDataAgentSpaceInfoResponseBodyData) *UpdateDataAgentSpaceInfoResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBody) SetErrorCode(v string) *UpdateDataAgentSpaceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBody) SetErrorMessage(v string) *UpdateDataAgentSpaceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBody) SetRequestId(v string) *UpdateDataAgentSpaceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBody) SetSuccess(v string) *UpdateDataAgentSpaceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataAgentSpaceInfoResponseBodyData struct {
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
	// space for test new
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1765962516
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// active
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 20
	TotalMember *string `json:"TotalMember,omitempty" xml:"TotalMember,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// yunqitest_v2
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// example:
	//
	// space for test new
	WorkspaceStatus *string `json:"WorkspaceStatus,omitempty" xml:"WorkspaceStatus,omitempty"`
}

func (s UpdateDataAgentSpaceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentSpaceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetTotalMember() *string {
	return s.TotalMember
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) GetWorkspaceStatus() *string {
	return s.WorkspaceStatus
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetCreateTime(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetCreator(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.Creator = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetDescription(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetModifyTime(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetRoleName(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetTotalMember(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.TotalMember = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetWorkspaceId(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetWorkspaceName(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) SetWorkspaceStatus(v string) *UpdateDataAgentSpaceInfoResponseBodyData {
	s.WorkspaceStatus = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
