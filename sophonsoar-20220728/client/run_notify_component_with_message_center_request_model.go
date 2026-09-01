// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithMessageCenterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *RunNotifyComponentWithMessageCenterRequest
	GetActionName() *string
	SetAliuid(v string) *RunNotifyComponentWithMessageCenterRequest
	GetAliuid() *string
	SetAssetId(v string) *RunNotifyComponentWithMessageCenterRequest
	GetAssetId() *string
	SetChannelTypeList(v []*string) *RunNotifyComponentWithMessageCenterRequest
	GetChannelTypeList() []*string
	SetComponentName(v string) *RunNotifyComponentWithMessageCenterRequest
	GetComponentName() *string
	SetEventId(v string) *RunNotifyComponentWithMessageCenterRequest
	GetEventId() *string
	SetLang(v string) *RunNotifyComponentWithMessageCenterRequest
	GetLang() *string
	SetNodeName(v string) *RunNotifyComponentWithMessageCenterRequest
	GetNodeName() *string
	SetParams(v string) *RunNotifyComponentWithMessageCenterRequest
	GetParams() *string
	SetPlaybookUuid(v string) *RunNotifyComponentWithMessageCenterRequest
	GetPlaybookUuid() *string
	SetRoleFor(v int64) *RunNotifyComponentWithMessageCenterRequest
	GetRoleFor() *int64
	SetRoleType(v string) *RunNotifyComponentWithMessageCenterRequest
	GetRoleType() *string
}

type RunNotifyComponentWithMessageCenterRequest struct {
	// The name of the playbook action.
	//
	// This parameter is required.
	//
	// example:
	//
	// notifyByMessageCenter
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The ID of the user who receives the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 146789xxxx733152
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The ID of the asset. This parameter is deprecated.
	//
	// example:
	//
	// 1
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// A collection of channel types. If you do not specify this parameter, messages are sent through all channels by default.
	ChannelTypeList []*string `json:"ChannelTypeList,omitempty" xml:"ChannelTypeList,omitempty" type:"Repeated"`
	// The name of the playbook component.
	//
	// This parameter is required.
	//
	// example:
	//
	// NotifyMessage
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The ID of the message event in Message Center. Valid values:
	//
	// - yundun_soar_incident_generate: An event is generated.
	//
	// - yundun_soar_alert_generate: An alert is generated.
	//
	// - yundun_soar_incident_update: An event is updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// yundun_soar_incident_generate
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the playbook node.
	//
	// This parameter is required.
	//
	// example:
	//
	// notify_message
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The template parameters for the message event.
	//
	// - Event generation: aliyunUID, incidentName, incidentID, startTime
	//
	// - Alert generation: aliyunUID, alertName, alertID, startTime
	//
	// - Event update: aliyunUID, incidentName, incidentID, startTime, endTime, status, level
	//
	// example:
	//
	// {"startTime":"test222","incidentName":"test123","incidentID":"teset123"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The UUID of the playbook.
	//
	// > Call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c88b5e-97ca-435d-8c20-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The user ID of the member. This parameter is used when an administrator calls the operation on behalf of a member.
	//
	// example:
	//
	// 1467894xxx733152
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0 (default): The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterRequest) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterRequest) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetActionName() *string {
	return s.ActionName
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetAliuid() *string {
	return s.Aliuid
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetChannelTypeList() []*string {
	return s.ChannelTypeList
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetEventId() *string {
	return s.EventId
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetLang() *string {
	return s.Lang
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetParams() *string {
	return s.Params
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *RunNotifyComponentWithMessageCenterRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetActionName(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.ActionName = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetAliuid(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.Aliuid = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetAssetId(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.AssetId = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetChannelTypeList(v []*string) *RunNotifyComponentWithMessageCenterRequest {
	s.ChannelTypeList = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetComponentName(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.ComponentName = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetEventId(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.EventId = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetLang(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.Lang = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetNodeName(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.NodeName = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetParams(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.Params = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetPlaybookUuid(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetRoleFor(v int64) *RunNotifyComponentWithMessageCenterRequest {
	s.RoleFor = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetRoleType(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.RoleType = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) Validate() error {
	return dara.Validate(s)
}
