// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetChannelStrategy(v map[string]*string) *Policy
	GetChannelStrategy() map[string]*string
	SetExpireTime(v string) *Policy
	GetExpireTime() *string
	SetOuterBizNo(v string) *Policy
	GetOuterBizNo() *string
	SetSpeed(v int32) *Policy
	GetSpeed() *int32
	SetStartTime(v string) *Policy
	GetStartTime() *string
}

type Policy struct {
	ChannelStrategy map[string]*string `json:"channelStrategy,omitempty" xml:"channelStrategy,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	OuterBizNo *string `json:"outerBizNo,omitempty" xml:"outerBizNo,omitempty"`
	// example:
	//
	// 5000
	Speed *int32 `json:"speed,omitempty" xml:"speed,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s Policy) String() string {
	return dara.Prettify(s)
}

func (s Policy) GoString() string {
	return s.String()
}

func (s *Policy) GetChannelStrategy() map[string]*string {
	return s.ChannelStrategy
}

func (s *Policy) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *Policy) GetOuterBizNo() *string {
	return s.OuterBizNo
}

func (s *Policy) GetSpeed() *int32 {
	return s.Speed
}

func (s *Policy) GetStartTime() *string {
	return s.StartTime
}

func (s *Policy) SetChannelStrategy(v map[string]*string) *Policy {
	s.ChannelStrategy = v
	return s
}

func (s *Policy) SetExpireTime(v string) *Policy {
	s.ExpireTime = &v
	return s
}

func (s *Policy) SetOuterBizNo(v string) *Policy {
	s.OuterBizNo = &v
	return s
}

func (s *Policy) SetSpeed(v int32) *Policy {
	s.Speed = &v
	return s
}

func (s *Policy) SetStartTime(v string) *Policy {
	s.StartTime = &v
	return s
}

func (s *Policy) Validate() error {
	return dara.Validate(s)
}
