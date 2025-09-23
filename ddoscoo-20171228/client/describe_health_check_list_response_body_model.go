// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*DescribeHealthCheckListResponseBodyListeners) *DescribeHealthCheckListResponseBody
	GetListeners() []*DescribeHealthCheckListResponseBodyListeners
	SetRequestId(v string) *DescribeHealthCheckListResponseBody
	GetRequestId() *string
}

type DescribeHealthCheckListResponseBody struct {
	Listeners []*DescribeHealthCheckListResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBody) GetListeners() []*DescribeHealthCheckListResponseBodyListeners {
	return s.Listeners
}

func (s *DescribeHealthCheckListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthCheckListResponseBody) SetListeners(v []*DescribeHealthCheckListResponseBodyListeners) *DescribeHealthCheckListResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeHealthCheckListResponseBody) SetRequestId(v string) *DescribeHealthCheckListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthCheckListResponseBodyListeners struct {
	// example:
	//
	// 233
	FrontendPort *int32                                                   `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	HealthCheck  *DescribeHealthCheckListResponseBodyListenersHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyListeners) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeHealthCheckListResponseBodyListeners) GetHealthCheck() *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	return s.HealthCheck
}

func (s *DescribeHealthCheckListResponseBodyListeners) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHealthCheckListResponseBodyListeners) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetFrontendPort(v int32) *DescribeHealthCheckListResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetHealthCheck(v *DescribeHealthCheckListResponseBodyListenersHealthCheck) *DescribeHealthCheckListResponseBodyListeners {
	s.HealthCheck = v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetInstanceId(v string) *DescribeHealthCheckListResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetProtocol(v string) *DescribeHealthCheckListResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthCheckListResponseBodyListenersHealthCheck struct {
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 500
	Down *int32 `json:"Down,omitempty" xml:"Down,omitempty"`
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 233
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// tcp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1000
	Up *int32 `json:"Up,omitempty" xml:"Up,omitempty"`
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyListenersHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyListenersHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetDown() *int32 {
	return s.Down
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetType() *string {
	return s.Type
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetUp() *int32 {
	return s.Up
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) GetUri() *string {
	return s.Uri
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetDomain(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetDown(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetInterval(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetPort(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetTimeout(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetType(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetUp(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Up = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetUri(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Uri = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) Validate() error {
	return dara.Validate(s)
}
