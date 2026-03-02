// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainMqConsoleLinkCmd interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ObtainMqConsoleLinkCmd
	GetEnv() *string
	SetGroupType(v string) *ObtainMqConsoleLinkCmd
	GetGroupType() *string
	SetMqGroupId(v string) *ObtainMqConsoleLinkCmd
	GetMqGroupId() *string
	SetPbcId(v int64) *ObtainMqConsoleLinkCmd
	GetPbcId() *int64
	SetServiceName(v string) *ObtainMqConsoleLinkCmd
	GetServiceName() *string
	SetTopicName(v string) *ObtainMqConsoleLinkCmd
	GetTopicName() *string
}

type ObtainMqConsoleLinkCmd struct {
	Env         *string `json:"env,omitempty" xml:"env,omitempty"`
	GroupType   *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	MqGroupId   *string `json:"mqGroupId,omitempty" xml:"mqGroupId,omitempty"`
	PbcId       *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TopicName   *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ObtainMqConsoleLinkCmd) String() string {
	return dara.Prettify(s)
}

func (s ObtainMqConsoleLinkCmd) GoString() string {
	return s.String()
}

func (s *ObtainMqConsoleLinkCmd) GetEnv() *string {
	return s.Env
}

func (s *ObtainMqConsoleLinkCmd) GetGroupType() *string {
	return s.GroupType
}

func (s *ObtainMqConsoleLinkCmd) GetMqGroupId() *string {
	return s.MqGroupId
}

func (s *ObtainMqConsoleLinkCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ObtainMqConsoleLinkCmd) GetServiceName() *string {
	return s.ServiceName
}

func (s *ObtainMqConsoleLinkCmd) GetTopicName() *string {
	return s.TopicName
}

func (s *ObtainMqConsoleLinkCmd) SetEnv(v string) *ObtainMqConsoleLinkCmd {
	s.Env = &v
	return s
}

func (s *ObtainMqConsoleLinkCmd) SetGroupType(v string) *ObtainMqConsoleLinkCmd {
	s.GroupType = &v
	return s
}

func (s *ObtainMqConsoleLinkCmd) SetMqGroupId(v string) *ObtainMqConsoleLinkCmd {
	s.MqGroupId = &v
	return s
}

func (s *ObtainMqConsoleLinkCmd) SetPbcId(v int64) *ObtainMqConsoleLinkCmd {
	s.PbcId = &v
	return s
}

func (s *ObtainMqConsoleLinkCmd) SetServiceName(v string) *ObtainMqConsoleLinkCmd {
	s.ServiceName = &v
	return s
}

func (s *ObtainMqConsoleLinkCmd) SetTopicName(v string) *ObtainMqConsoleLinkCmd {
	s.TopicName = &v
	return s
}

func (s *ObtainMqConsoleLinkCmd) Validate() error {
	return dara.Validate(s)
}
