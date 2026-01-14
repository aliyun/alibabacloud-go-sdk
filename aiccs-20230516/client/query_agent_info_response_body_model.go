// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgentInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAgentInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *QueryAgentInfoResponseBody
	GetCode() *int64
	SetMessage(v string) *QueryAgentInfoResponseBody
	GetMessage() *string
	SetModel(v *QueryAgentInfoResponseBodyModel) *QueryAgentInfoResponseBody
	GetModel() *QueryAgentInfoResponseBodyModel
	SetRequestId(v string) *QueryAgentInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAgentInfoResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *QueryAgentInfoResponseBody
	GetTimestamp() *int64
}

type QueryAgentInfoResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 43
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *QueryAgentInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 17
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s QueryAgentInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAgentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAgentInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAgentInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *QueryAgentInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAgentInfoResponseBody) GetModel() *QueryAgentInfoResponseBodyModel {
	return s.Model
}

func (s *QueryAgentInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAgentInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAgentInfoResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QueryAgentInfoResponseBody) SetAccessDeniedDetail(v string) *QueryAgentInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAgentInfoResponseBody) SetCode(v int64) *QueryAgentInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAgentInfoResponseBody) SetMessage(v string) *QueryAgentInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAgentInfoResponseBody) SetModel(v *QueryAgentInfoResponseBodyModel) *QueryAgentInfoResponseBody {
	s.Model = v
	return s
}

func (s *QueryAgentInfoResponseBody) SetRequestId(v string) *QueryAgentInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAgentInfoResponseBody) SetSuccess(v bool) *QueryAgentInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAgentInfoResponseBody) SetTimestamp(v int64) *QueryAgentInfoResponseBody {
	s.Timestamp = &v
	return s
}

func (s *QueryAgentInfoResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAgentInfoResponseBodyModel struct {
	// 坐席账号
	//
	// example:
	//
	// a
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// 账号启用状态，0-停用，1-启用，默认1
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 坐席分机号
	//
	// example:
	//
	// 1
	AgentExtension *int64 `json:"AgentExtension,omitempty" xml:"AgentExtension,omitempty"`
	// 坐席组ID列表
	AgentGroupIds []*int64 `json:"AgentGroupIds,omitempty" xml:"AgentGroupIds,omitempty" type:"Repeated"`
	// 坐席组列表
	AgentGroupList []*QueryAgentInfoResponseBodyModelAgentGroupList `json:"AgentGroupList,omitempty" xml:"AgentGroupList,omitempty" type:"Repeated"`
	// 坐席ID
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席状态，1:在线；2:忙碌；3:小休；4:离线
	//
	// example:
	//
	// 1
	AgentStatus *int64 `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// c
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2026-11-11 11:11:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 分机密码
	//
	// example:
	//
	// a
	ExtensionPwd *string `json:"ExtensionPwd,omitempty" xml:"ExtensionPwd,omitempty"`
	// 分机注册地址
	//
	// example:
	//
	// b
	ExtensionServer *string `json:"ExtensionServer,omitempty" xml:"ExtensionServer,omitempty"`
	// 坐席名称
	//
	// example:
	//
	// b
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// WebSocket分机注册协议
	//
	// example:
	//
	// b
	WsProtocol *string `json:"WsProtocol,omitempty" xml:"WsProtocol,omitempty"`
	// WebSocket分机注册地址
	//
	// example:
	//
	// a
	WsRegisterAddress *string `json:"WsRegisterAddress,omitempty" xml:"WsRegisterAddress,omitempty"`
}

func (s QueryAgentInfoResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s QueryAgentInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryAgentInfoResponseBodyModel) GetAccount() *string {
	return s.Account
}

func (s *QueryAgentInfoResponseBodyModel) GetActive() *int64 {
	return s.Active
}

func (s *QueryAgentInfoResponseBodyModel) GetAgentExtension() *int64 {
	return s.AgentExtension
}

func (s *QueryAgentInfoResponseBodyModel) GetAgentGroupIds() []*int64 {
	return s.AgentGroupIds
}

func (s *QueryAgentInfoResponseBodyModel) GetAgentGroupList() []*QueryAgentInfoResponseBodyModelAgentGroupList {
	return s.AgentGroupList
}

func (s *QueryAgentInfoResponseBodyModel) GetAgentId() *int64 {
	return s.AgentId
}

func (s *QueryAgentInfoResponseBodyModel) GetAgentStatus() *int64 {
	return s.AgentStatus
}

func (s *QueryAgentInfoResponseBodyModel) GetAgentTag() *string {
	return s.AgentTag
}

func (s *QueryAgentInfoResponseBodyModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryAgentInfoResponseBodyModel) GetExtensionPwd() *string {
	return s.ExtensionPwd
}

func (s *QueryAgentInfoResponseBodyModel) GetExtensionServer() *string {
	return s.ExtensionServer
}

func (s *QueryAgentInfoResponseBodyModel) GetName() *string {
	return s.Name
}

func (s *QueryAgentInfoResponseBodyModel) GetWsProtocol() *string {
	return s.WsProtocol
}

func (s *QueryAgentInfoResponseBodyModel) GetWsRegisterAddress() *string {
	return s.WsRegisterAddress
}

func (s *QueryAgentInfoResponseBodyModel) SetAccount(v string) *QueryAgentInfoResponseBodyModel {
	s.Account = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetActive(v int64) *QueryAgentInfoResponseBodyModel {
	s.Active = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetAgentExtension(v int64) *QueryAgentInfoResponseBodyModel {
	s.AgentExtension = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetAgentGroupIds(v []*int64) *QueryAgentInfoResponseBodyModel {
	s.AgentGroupIds = v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetAgentGroupList(v []*QueryAgentInfoResponseBodyModelAgentGroupList) *QueryAgentInfoResponseBodyModel {
	s.AgentGroupList = v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetAgentId(v int64) *QueryAgentInfoResponseBodyModel {
	s.AgentId = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetAgentStatus(v int64) *QueryAgentInfoResponseBodyModel {
	s.AgentStatus = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetAgentTag(v string) *QueryAgentInfoResponseBodyModel {
	s.AgentTag = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetCreateTime(v string) *QueryAgentInfoResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetExtensionPwd(v string) *QueryAgentInfoResponseBodyModel {
	s.ExtensionPwd = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetExtensionServer(v string) *QueryAgentInfoResponseBodyModel {
	s.ExtensionServer = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetName(v string) *QueryAgentInfoResponseBodyModel {
	s.Name = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetWsProtocol(v string) *QueryAgentInfoResponseBodyModel {
	s.WsProtocol = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) SetWsRegisterAddress(v string) *QueryAgentInfoResponseBodyModel {
	s.WsRegisterAddress = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModel) Validate() error {
	if s.AgentGroupList != nil {
		for _, item := range s.AgentGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAgentInfoResponseBodyModelAgentGroupList struct {
	// 坐席组ID
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 坐席组名称
	//
	// example:
	//
	// a
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s QueryAgentInfoResponseBodyModelAgentGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryAgentInfoResponseBodyModelAgentGroupList) GoString() string {
	return s.String()
}

func (s *QueryAgentInfoResponseBodyModelAgentGroupList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryAgentInfoResponseBodyModelAgentGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryAgentInfoResponseBodyModelAgentGroupList) SetGroupId(v int64) *QueryAgentInfoResponseBodyModelAgentGroupList {
	s.GroupId = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModelAgentGroupList) SetGroupName(v string) *QueryAgentInfoResponseBodyModelAgentGroupList {
	s.GroupName = &v
	return s
}

func (s *QueryAgentInfoResponseBodyModelAgentGroupList) Validate() error {
	return dara.Validate(s)
}
