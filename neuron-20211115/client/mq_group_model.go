// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqGroup interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *MqGroup
	GetEnv() *string
	SetGmtCreate(v string) *MqGroup
	GetGmtCreate() *string
	SetGmtModified(v string) *MqGroup
	GetGmtModified() *string
	SetGroupId(v string) *MqGroup
	GetGroupId() *string
	SetId(v int64) *MqGroup
	GetId() *int64
	SetMessageType(v int32) *MqGroup
	GetMessageType() *int32
	SetProtocolType(v string) *MqGroup
	GetProtocolType() *string
	SetRemark(v string) *MqGroup
	GetRemark() *string
	SetRequestId(v string) *MqGroup
	GetRequestId() *string
	SetServiceId(v string) *MqGroup
	GetServiceId() *string
}

type MqGroup struct {
	Env          *string `json:"env,omitempty" xml:"env,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupId      *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	MessageType  *int32  `json:"messageType,omitempty" xml:"messageType,omitempty"`
	ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
	Remark       *string `json:"remark,omitempty" xml:"remark,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceId    *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s MqGroup) String() string {
	return dara.Prettify(s)
}

func (s MqGroup) GoString() string {
	return s.String()
}

func (s *MqGroup) GetEnv() *string {
	return s.Env
}

func (s *MqGroup) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MqGroup) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MqGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *MqGroup) GetId() *int64 {
	return s.Id
}

func (s *MqGroup) GetMessageType() *int32 {
	return s.MessageType
}

func (s *MqGroup) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *MqGroup) GetRemark() *string {
	return s.Remark
}

func (s *MqGroup) GetRequestId() *string {
	return s.RequestId
}

func (s *MqGroup) GetServiceId() *string {
	return s.ServiceId
}

func (s *MqGroup) SetEnv(v string) *MqGroup {
	s.Env = &v
	return s
}

func (s *MqGroup) SetGmtCreate(v string) *MqGroup {
	s.GmtCreate = &v
	return s
}

func (s *MqGroup) SetGmtModified(v string) *MqGroup {
	s.GmtModified = &v
	return s
}

func (s *MqGroup) SetGroupId(v string) *MqGroup {
	s.GroupId = &v
	return s
}

func (s *MqGroup) SetId(v int64) *MqGroup {
	s.Id = &v
	return s
}

func (s *MqGroup) SetMessageType(v int32) *MqGroup {
	s.MessageType = &v
	return s
}

func (s *MqGroup) SetProtocolType(v string) *MqGroup {
	s.ProtocolType = &v
	return s
}

func (s *MqGroup) SetRemark(v string) *MqGroup {
	s.Remark = &v
	return s
}

func (s *MqGroup) SetRequestId(v string) *MqGroup {
	s.RequestId = &v
	return s
}

func (s *MqGroup) SetServiceId(v string) *MqGroup {
	s.ServiceId = &v
	return s
}

func (s *MqGroup) Validate() error {
	return dara.Validate(s)
}
