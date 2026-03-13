// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceDTSParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBrokerUrl(v string) *SourceDTSParameters
	GetBrokerUrl() *string
	SetInitCheckPoint(v int32) *SourceDTSParameters
	GetInitCheckPoint() *int32
	SetPassword(v string) *SourceDTSParameters
	GetPassword() *string
	SetRegionId(v string) *SourceDTSParameters
	GetRegionId() *string
	SetSid(v string) *SourceDTSParameters
	GetSid() *string
	SetTaskId(v string) *SourceDTSParameters
	GetTaskId() *string
	SetTopic(v string) *SourceDTSParameters
	GetTopic() *string
	SetUsername(v string) *SourceDTSParameters
	GetUsername() *string
}

type SourceDTSParameters struct {
	// The network address and port number of the change tracking instance.
	//
	// example:
	//
	// dts-cn-shanghai-vpc.com:18003
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The UNIX timestamp that is generated when the SDK client consumes the first data record.
	//
	// example:
	//
	// 1677340805
	InitCheckPoint *int32 `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The consumer group password.
	//
	// example:
	//
	// dtsTest123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region of the DTS instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// dtse34j22j025a****
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// e34z2gm325q****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the tracked topic of the change tracking instance.
	//
	// example:
	//
	// cn_shanghai_vpc_rm_uf6398ykj0218****_dts_trigger_upgrade_from_old_version2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The account of the consumer group.
	//
	// example:
	//
	// dts_trigger
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s SourceDTSParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceDTSParameters) GoString() string {
	return s.String()
}

func (s *SourceDTSParameters) GetBrokerUrl() *string {
	return s.BrokerUrl
}

func (s *SourceDTSParameters) GetInitCheckPoint() *int32 {
	return s.InitCheckPoint
}

func (s *SourceDTSParameters) GetPassword() *string {
	return s.Password
}

func (s *SourceDTSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceDTSParameters) GetSid() *string {
	return s.Sid
}

func (s *SourceDTSParameters) GetTaskId() *string {
	return s.TaskId
}

func (s *SourceDTSParameters) GetTopic() *string {
	return s.Topic
}

func (s *SourceDTSParameters) GetUsername() *string {
	return s.Username
}

func (s *SourceDTSParameters) SetBrokerUrl(v string) *SourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *SourceDTSParameters) SetInitCheckPoint(v int32) *SourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *SourceDTSParameters) SetPassword(v string) *SourceDTSParameters {
	s.Password = &v
	return s
}

func (s *SourceDTSParameters) SetRegionId(v string) *SourceDTSParameters {
	s.RegionId = &v
	return s
}

func (s *SourceDTSParameters) SetSid(v string) *SourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *SourceDTSParameters) SetTaskId(v string) *SourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *SourceDTSParameters) SetTopic(v string) *SourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *SourceDTSParameters) SetUsername(v string) *SourceDTSParameters {
	s.Username = &v
	return s
}

func (s *SourceDTSParameters) Validate() error {
	return dara.Validate(s)
}
