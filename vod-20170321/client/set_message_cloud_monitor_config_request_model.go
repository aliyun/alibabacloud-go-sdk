// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMessageCloudMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SetMessageCloudMonitorConfigRequest
	GetAppId() *string
	SetEventTypeList(v string) *SetMessageCloudMonitorConfigRequest
	GetEventTypeList() *string
	SetGroupId(v int64) *SetMessageCloudMonitorConfigRequest
	GetGroupId() *int64
	SetOwnerAccount(v string) *SetMessageCloudMonitorConfigRequest
	GetOwnerAccount() *string
}

type SetMessageCloudMonitorConfigRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	GroupId       *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s SetMessageCloudMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetMessageCloudMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *SetMessageCloudMonitorConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *SetMessageCloudMonitorConfigRequest) GetEventTypeList() *string {
	return s.EventTypeList
}

func (s *SetMessageCloudMonitorConfigRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *SetMessageCloudMonitorConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetMessageCloudMonitorConfigRequest) SetAppId(v string) *SetMessageCloudMonitorConfigRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageCloudMonitorConfigRequest) SetEventTypeList(v string) *SetMessageCloudMonitorConfigRequest {
	s.EventTypeList = &v
	return s
}

func (s *SetMessageCloudMonitorConfigRequest) SetGroupId(v int64) *SetMessageCloudMonitorConfigRequest {
	s.GroupId = &v
	return s
}

func (s *SetMessageCloudMonitorConfigRequest) SetOwnerAccount(v string) *SetMessageCloudMonitorConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetMessageCloudMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}
