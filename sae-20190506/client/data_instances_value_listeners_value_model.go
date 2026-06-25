// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataInstancesValueListenersValue interface {
	dara.Model
	String() string
	GoString() string
	SetProtocol(v string) *DataInstancesValueListenersValue
	GetProtocol() *string
	SetPort(v int32) *DataInstancesValueListenersValue
	GetPort() *int32
	SetStatus(v string) *DataInstancesValueListenersValue
	GetStatus() *string
	SetTargetPort(v int32) *DataInstancesValueListenersValue
	GetTargetPort() *int32
	SetCertIds(v string) *DataInstancesValueListenersValue
	GetCertIds() *string
}

type DataInstancesValueListenersValue struct {
	// The listener protocol.
	//
	// example:
	//
	// TCPSSL
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The state of the NLB listener. Valid values:
	//
	// - **Creating**: The listener is being created.
	//
	// - **Configuring**: The listener is being configured.
	//
	// - **Bounded**: The listener is running as expected.
	//
	// - **Unbinding**: The listener is being deleted.
	//
	// - **Failed**: The listener is unavailable.
	//
	// example:
	//
	// Bounded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The target port.
	//
	// example:
	//
	// 8080
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The server certificate ID.
	//
	// example:
	//
	// 123157******
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
}

func (s DataInstancesValueListenersValue) String() string {
	return dara.Prettify(s)
}

func (s DataInstancesValueListenersValue) GoString() string {
	return s.String()
}

func (s *DataInstancesValueListenersValue) GetProtocol() *string {
	return s.Protocol
}

func (s *DataInstancesValueListenersValue) GetPort() *int32 {
	return s.Port
}

func (s *DataInstancesValueListenersValue) GetStatus() *string {
	return s.Status
}

func (s *DataInstancesValueListenersValue) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *DataInstancesValueListenersValue) GetCertIds() *string {
	return s.CertIds
}

func (s *DataInstancesValueListenersValue) SetProtocol(v string) *DataInstancesValueListenersValue {
	s.Protocol = &v
	return s
}

func (s *DataInstancesValueListenersValue) SetPort(v int32) *DataInstancesValueListenersValue {
	s.Port = &v
	return s
}

func (s *DataInstancesValueListenersValue) SetStatus(v string) *DataInstancesValueListenersValue {
	s.Status = &v
	return s
}

func (s *DataInstancesValueListenersValue) SetTargetPort(v int32) *DataInstancesValueListenersValue {
	s.TargetPort = &v
	return s
}

func (s *DataInstancesValueListenersValue) SetCertIds(v string) *DataInstancesValueListenersValue {
	s.CertIds = &v
	return s
}

func (s *DataInstancesValueListenersValue) Validate() error {
	return dara.Validate(s)
}
