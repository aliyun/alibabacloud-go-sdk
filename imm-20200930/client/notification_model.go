// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotification interface {
	dara.Model
	String() string
	GoString() string
	SetExtendedMessageURI(v string) *Notification
	GetExtendedMessageURI() *string
	SetMNS(v *MNS) *Notification
	GetMNS() *MNS
	SetRocketMQ(v *RocketMQ) *Notification
	GetRocketMQ() *RocketMQ
}

type Notification struct {
	ExtendedMessageURI *string   `json:"ExtendedMessageURI,omitempty" xml:"ExtendedMessageURI,omitempty"`
	MNS                *MNS      `json:"MNS,omitempty" xml:"MNS,omitempty"`
	RocketMQ           *RocketMQ `json:"RocketMQ,omitempty" xml:"RocketMQ,omitempty"`
}

func (s Notification) String() string {
	return dara.Prettify(s)
}

func (s Notification) GoString() string {
	return s.String()
}

func (s *Notification) GetExtendedMessageURI() *string {
	return s.ExtendedMessageURI
}

func (s *Notification) GetMNS() *MNS {
	return s.MNS
}

func (s *Notification) GetRocketMQ() *RocketMQ {
	return s.RocketMQ
}

func (s *Notification) SetExtendedMessageURI(v string) *Notification {
	s.ExtendedMessageURI = &v
	return s
}

func (s *Notification) SetMNS(v *MNS) *Notification {
	s.MNS = v
	return s
}

func (s *Notification) SetRocketMQ(v *RocketMQ) *Notification {
	s.RocketMQ = v
	return s
}

func (s *Notification) Validate() error {
	if s.MNS != nil {
		if err := s.MNS.Validate(); err != nil {
			return err
		}
	}
	if s.RocketMQ != nil {
		if err := s.RocketMQ.Validate(); err != nil {
			return err
		}
	}
	return nil
}
