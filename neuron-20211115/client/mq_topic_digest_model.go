// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqTopicDigest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *MqTopicDigest
	GetAlias() *string
	SetCreateTime(v string) *MqTopicDigest
	GetCreateTime() *string
	SetId(v int64) *MqTopicDigest
	GetId() *int64
	SetMessageType(v string) *MqTopicDigest
	GetMessageType() *string
	SetName(v string) *MqTopicDigest
	GetName() *string
	SetRemark(v string) *MqTopicDigest
	GetRemark() *string
}

type MqTopicDigest struct {
	Alias       *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s MqTopicDigest) String() string {
	return dara.Prettify(s)
}

func (s MqTopicDigest) GoString() string {
	return s.String()
}

func (s *MqTopicDigest) GetAlias() *string {
	return s.Alias
}

func (s *MqTopicDigest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MqTopicDigest) GetId() *int64 {
	return s.Id
}

func (s *MqTopicDigest) GetMessageType() *string {
	return s.MessageType
}

func (s *MqTopicDigest) GetName() *string {
	return s.Name
}

func (s *MqTopicDigest) GetRemark() *string {
	return s.Remark
}

func (s *MqTopicDigest) SetAlias(v string) *MqTopicDigest {
	s.Alias = &v
	return s
}

func (s *MqTopicDigest) SetCreateTime(v string) *MqTopicDigest {
	s.CreateTime = &v
	return s
}

func (s *MqTopicDigest) SetId(v int64) *MqTopicDigest {
	s.Id = &v
	return s
}

func (s *MqTopicDigest) SetMessageType(v string) *MqTopicDigest {
	s.MessageType = &v
	return s
}

func (s *MqTopicDigest) SetName(v string) *MqTopicDigest {
	s.Name = &v
	return s
}

func (s *MqTopicDigest) SetRemark(v string) *MqTopicDigest {
	s.Remark = &v
	return s
}

func (s *MqTopicDigest) Validate() error {
	return dara.Validate(s)
}
