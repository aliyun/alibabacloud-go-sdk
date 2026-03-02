// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqGroupCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *MqGroupCreateCmd
	GetEnv() *string
	SetLaneId(v int64) *MqGroupCreateCmd
	GetLaneId() *int64
	SetLaneName(v string) *MqGroupCreateCmd
	GetLaneName() *string
	SetMessageType(v int32) *MqGroupCreateCmd
	GetMessageType() *int32
	SetProtocolType(v string) *MqGroupCreateCmd
	GetProtocolType() *string
	SetRemark(v string) *MqGroupCreateCmd
	GetRemark() *string
	SetServiceId(v string) *MqGroupCreateCmd
	GetServiceId() *string
}

type MqGroupCreateCmd struct {
	Env          *string `json:"env,omitempty" xml:"env,omitempty"`
	LaneId       *int64  `json:"laneId,omitempty" xml:"laneId,omitempty"`
	LaneName     *string `json:"laneName,omitempty" xml:"laneName,omitempty"`
	MessageType  *int32  `json:"messageType,omitempty" xml:"messageType,omitempty"`
	ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
	Remark       *string `json:"remark,omitempty" xml:"remark,omitempty"`
	ServiceId    *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s MqGroupCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s MqGroupCreateCmd) GoString() string {
	return s.String()
}

func (s *MqGroupCreateCmd) GetEnv() *string {
	return s.Env
}

func (s *MqGroupCreateCmd) GetLaneId() *int64 {
	return s.LaneId
}

func (s *MqGroupCreateCmd) GetLaneName() *string {
	return s.LaneName
}

func (s *MqGroupCreateCmd) GetMessageType() *int32 {
	return s.MessageType
}

func (s *MqGroupCreateCmd) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *MqGroupCreateCmd) GetRemark() *string {
	return s.Remark
}

func (s *MqGroupCreateCmd) GetServiceId() *string {
	return s.ServiceId
}

func (s *MqGroupCreateCmd) SetEnv(v string) *MqGroupCreateCmd {
	s.Env = &v
	return s
}

func (s *MqGroupCreateCmd) SetLaneId(v int64) *MqGroupCreateCmd {
	s.LaneId = &v
	return s
}

func (s *MqGroupCreateCmd) SetLaneName(v string) *MqGroupCreateCmd {
	s.LaneName = &v
	return s
}

func (s *MqGroupCreateCmd) SetMessageType(v int32) *MqGroupCreateCmd {
	s.MessageType = &v
	return s
}

func (s *MqGroupCreateCmd) SetProtocolType(v string) *MqGroupCreateCmd {
	s.ProtocolType = &v
	return s
}

func (s *MqGroupCreateCmd) SetRemark(v string) *MqGroupCreateCmd {
	s.Remark = &v
	return s
}

func (s *MqGroupCreateCmd) SetServiceId(v string) *MqGroupCreateCmd {
	s.ServiceId = &v
	return s
}

func (s *MqGroupCreateCmd) Validate() error {
	return dara.Validate(s)
}
