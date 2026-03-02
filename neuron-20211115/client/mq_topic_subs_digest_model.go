// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqTopicSubsDigest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageModel(v string) *MqTopicSubsDigest
	GetMessageModel() *string
	SetServiceId(v string) *MqTopicSubsDigest
	GetServiceId() *string
	SetServiceName(v string) *MqTopicSubsDigest
	GetServiceName() *string
	SetSubsExpression(v string) *MqTopicSubsDigest
	GetSubsExpression() *string
}

type MqTopicSubsDigest struct {
	MessageModel   *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	ServiceId      *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SubsExpression *string `json:"SubsExpression,omitempty" xml:"SubsExpression,omitempty"`
}

func (s MqTopicSubsDigest) String() string {
	return dara.Prettify(s)
}

func (s MqTopicSubsDigest) GoString() string {
	return s.String()
}

func (s *MqTopicSubsDigest) GetMessageModel() *string {
	return s.MessageModel
}

func (s *MqTopicSubsDigest) GetServiceId() *string {
	return s.ServiceId
}

func (s *MqTopicSubsDigest) GetServiceName() *string {
	return s.ServiceName
}

func (s *MqTopicSubsDigest) GetSubsExpression() *string {
	return s.SubsExpression
}

func (s *MqTopicSubsDigest) SetMessageModel(v string) *MqTopicSubsDigest {
	s.MessageModel = &v
	return s
}

func (s *MqTopicSubsDigest) SetServiceId(v string) *MqTopicSubsDigest {
	s.ServiceId = &v
	return s
}

func (s *MqTopicSubsDigest) SetServiceName(v string) *MqTopicSubsDigest {
	s.ServiceName = &v
	return s
}

func (s *MqTopicSubsDigest) SetSubsExpression(v string) *MqTopicSubsDigest {
	s.SubsExpression = &v
	return s
}

func (s *MqTopicSubsDigest) Validate() error {
	return dara.Validate(s)
}
