// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMappCenterWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMappCenterWorkspaceResult(v *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) *ListMappCenterWorkspacesResponseBody
	GetListMappCenterWorkspaceResult() *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult
	SetRequestId(v string) *ListMappCenterWorkspacesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMappCenterWorkspacesResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMappCenterWorkspacesResponseBody
	GetResultMessage() *string
}

type ListMappCenterWorkspacesResponseBody struct {
	ListMappCenterWorkspaceResult *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult `json:"ListMappCenterWorkspaceResult,omitempty" xml:"ListMappCenterWorkspaceResult,omitempty" type:"Struct"`
	RequestId                     *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                    *string                                                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage                 *string                                                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMappCenterWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMappCenterWorkspacesResponseBody) GetListMappCenterWorkspaceResult() *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult {
	return s.ListMappCenterWorkspaceResult
}

func (s *ListMappCenterWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMappCenterWorkspacesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMappCenterWorkspacesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMappCenterWorkspacesResponseBody) SetListMappCenterWorkspaceResult(v *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) *ListMappCenterWorkspacesResponseBody {
	s.ListMappCenterWorkspaceResult = v
	return s
}

func (s *ListMappCenterWorkspacesResponseBody) SetRequestId(v string) *ListMappCenterWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBody) SetResultCode(v string) *ListMappCenterWorkspacesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBody) SetResultMessage(v string) *ListMappCenterWorkspacesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult struct {
	MappCenterWorkspaceList []*ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList `json:"MappCenterWorkspaceList,omitempty" xml:"MappCenterWorkspaceList,omitempty" type:"Repeated"`
	ResultMsg               *string                                                                                     `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success                 *bool                                                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	UserId                  *string                                                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) GoString() string {
	return s.String()
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) GetMappCenterWorkspaceList() []*ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	return s.MappCenterWorkspaceList
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) GetUserId() *string {
	return s.UserId
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) SetMappCenterWorkspaceList(v []*ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult {
	s.MappCenterWorkspaceList = v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) SetResultMsg(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) SetSuccess(v bool) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult {
	s.Success = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) SetUserId(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult {
	s.UserId = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResult) Validate() error {
	return dara.Validate(s)
}

type ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList struct {
	CompatibleId *string `json:"CompatibleId,omitempty" xml:"CompatibleId,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid          *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	Zones        *string `json:"Zones,omitempty" xml:"Zones,omitempty"`
}

func (s ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GoString() string {
	return s.String()
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetCompatibleId() *string {
	return s.CompatibleId
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetId() *string {
	return s.Id
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetRegion() *string {
	return s.Region
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetStatus() *string {
	return s.Status
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetType() *string {
	return s.Type
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetUid() *int64 {
	return s.Uid
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) GetZones() *string {
	return s.Zones
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetCompatibleId(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.CompatibleId = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetCreateTime(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.CreateTime = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetDisplayName(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.DisplayName = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetId(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.Id = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetRegion(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.Region = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetStatus(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.Status = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetTenantId(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.TenantId = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetType(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.Type = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetUid(v int64) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.Uid = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetUpdateTime(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.UpdateTime = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetWorkspaceId(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.WorkspaceId = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) SetZones(v string) *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList {
	s.Zones = &v
	return s
}

func (s *ListMappCenterWorkspacesResponseBodyListMappCenterWorkspaceResultMappCenterWorkspaceList) Validate() error {
	return dara.Validate(s)
}
