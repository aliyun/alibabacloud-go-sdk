// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqTopic interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *MqTopic
	GetEnv() *string
	SetGmtCreate(v string) *MqTopic
	GetGmtCreate() *string
	SetGmtModified(v string) *MqTopic
	GetGmtModified() *string
	SetId(v int64) *MqTopic
	GetId() *int64
	SetMessageType(v string) *MqTopic
	GetMessageType() *string
	SetName(v string) *MqTopic
	GetName() *string
	SetPbcId(v string) *MqTopic
	GetPbcId() *string
	SetRemark(v string) *MqTopic
	GetRemark() *string
	SetRequestId(v string) *MqTopic
	GetRequestId() *string
}

type MqTopic struct {
	// This parameter is required.
	Env         *string `json:"env,omitempty" xml:"env,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	PbcId     *string `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	Remark    *string `json:"remark,omitempty" xml:"remark,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MqTopic) String() string {
	return dara.Prettify(s)
}

func (s MqTopic) GoString() string {
	return s.String()
}

func (s *MqTopic) GetEnv() *string {
	return s.Env
}

func (s *MqTopic) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MqTopic) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MqTopic) GetId() *int64 {
	return s.Id
}

func (s *MqTopic) GetMessageType() *string {
	return s.MessageType
}

func (s *MqTopic) GetName() *string {
	return s.Name
}

func (s *MqTopic) GetPbcId() *string {
	return s.PbcId
}

func (s *MqTopic) GetRemark() *string {
	return s.Remark
}

func (s *MqTopic) GetRequestId() *string {
	return s.RequestId
}

func (s *MqTopic) SetEnv(v string) *MqTopic {
	s.Env = &v
	return s
}

func (s *MqTopic) SetGmtCreate(v string) *MqTopic {
	s.GmtCreate = &v
	return s
}

func (s *MqTopic) SetGmtModified(v string) *MqTopic {
	s.GmtModified = &v
	return s
}

func (s *MqTopic) SetId(v int64) *MqTopic {
	s.Id = &v
	return s
}

func (s *MqTopic) SetMessageType(v string) *MqTopic {
	s.MessageType = &v
	return s
}

func (s *MqTopic) SetName(v string) *MqTopic {
	s.Name = &v
	return s
}

func (s *MqTopic) SetPbcId(v string) *MqTopic {
	s.PbcId = &v
	return s
}

func (s *MqTopic) SetRemark(v string) *MqTopic {
	s.Remark = &v
	return s
}

func (s *MqTopic) SetRequestId(v string) *MqTopic {
	s.RequestId = &v
	return s
}

func (s *MqTopic) Validate() error {
	return dara.Validate(s)
}
