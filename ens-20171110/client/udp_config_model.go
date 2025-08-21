// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUdpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetHashKey(v string) *UdpConfig
	GetHashKey() *string
	SetScheduler(v string) *UdpConfig
	GetScheduler() *string
}

type UdpConfig struct {
	HashKey   *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s UdpConfig) String() string {
	return dara.Prettify(s)
}

func (s UdpConfig) GoString() string {
	return s.String()
}

func (s *UdpConfig) GetHashKey() *string {
	return s.HashKey
}

func (s *UdpConfig) GetScheduler() *string {
	return s.Scheduler
}

func (s *UdpConfig) SetHashKey(v string) *UdpConfig {
	s.HashKey = &v
	return s
}

func (s *UdpConfig) SetScheduler(v string) *UdpConfig {
	s.Scheduler = &v
	return s
}

func (s *UdpConfig) Validate() error {
	return dara.Validate(s)
}
