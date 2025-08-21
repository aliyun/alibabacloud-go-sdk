// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTcpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEstablishedTimeout(v int32) *TcpConfig
	GetEstablishedTimeout() *int32
	SetPersistenceTimeout(v int32) *TcpConfig
	GetPersistenceTimeout() *int32
	SetScheduler(v string) *TcpConfig
	GetScheduler() *string
}

type TcpConfig struct {
	EstablishedTimeout *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	PersistenceTimeout *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	Scheduler          *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s TcpConfig) String() string {
	return dara.Prettify(s)
}

func (s TcpConfig) GoString() string {
	return s.String()
}

func (s *TcpConfig) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *TcpConfig) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *TcpConfig) GetScheduler() *string {
	return s.Scheduler
}

func (s *TcpConfig) SetEstablishedTimeout(v int32) *TcpConfig {
	s.EstablishedTimeout = &v
	return s
}

func (s *TcpConfig) SetPersistenceTimeout(v int32) *TcpConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *TcpConfig) SetScheduler(v string) *TcpConfig {
	s.Scheduler = &v
	return s
}

func (s *TcpConfig) Validate() error {
	return dara.Validate(s)
}
