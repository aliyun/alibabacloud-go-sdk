// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckList(v []*DescribeHealthCheckListResponseBodyHealthCheckList) *DescribeHealthCheckListResponseBody
	GetHealthCheckList() []*DescribeHealthCheckListResponseBodyHealthCheckList
	SetRequestId(v string) *DescribeHealthCheckListResponseBody
	GetRequestId() *string
}

type DescribeHealthCheckListResponseBody struct {
	// An array that consists of information about the health check configuration.
	HealthCheckList []*DescribeHealthCheckListResponseBodyHealthCheckList `json:"HealthCheckList,omitempty" xml:"HealthCheckList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 83B4AF42-E8EE-4DC9-BD73-87B7733A36F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBody) GetHealthCheckList() []*DescribeHealthCheckListResponseBodyHealthCheckList {
	return s.HealthCheckList
}

func (s *DescribeHealthCheckListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthCheckListResponseBody) SetHealthCheckList(v []*DescribeHealthCheckListResponseBodyHealthCheckList) *DescribeHealthCheckListResponseBody {
	s.HealthCheckList = v
	return s
}

func (s *DescribeHealthCheckListResponseBody) SetRequestId(v string) *DescribeHealthCheckListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBody) Validate() error {
	if s.HealthCheckList != nil {
		for _, item := range s.HealthCheckList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHealthCheckListResponseBodyHealthCheckList struct {
	// The forwarding port.
	//
	// example:
	//
	// 8080
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The health check configuration.
	HealthCheck *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The forwarding protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyHealthCheckList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyHealthCheckList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) GetHealthCheck() *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	return s.HealthCheck
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetFrontendPort(v int32) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetHealthCheck(v *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.HealthCheck = v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetInstanceId(v string) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetProtocol(v string) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) Validate() error {
	if s.HealthCheck != nil {
		if err := s.HealthCheck.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck struct {
	// The domain name.
	//
	// >  This parameter is returned only when the Layer 7 health check configuration is queried.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of consecutive failed health checks that must occur before a port is declared unhealthy. Valid values: **1*	- to **10**.
	//
	// example:
	//
	// 3
	Down *int32 `json:"Down,omitempty" xml:"Down,omitempty"`
	// The interval at which checks are performed. Valid values: **1*	- to **30**. Unit: seconds.
	//
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The port that was checked.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The response timeout period. Valid values: **1*	- to **30**. Unit: seconds.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**: The Layer 4 health check configuration was queried.
	//
	// 	- **http**: The Layer 7 health check configuration was queried.
	//
	// example:
	//
	// tcp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The number of consecutive successful health checks that must occur before a port is declared healthy. Valid values: **1*	- to **10**.
	//
	// example:
	//
	// 3
	Up *int32 `json:"Up,omitempty" xml:"Up,omitempty"`
	// The check path.
	//
	// >  This parameter is returned only when the Layer 7 health check configuration is queried.
	//
	// example:
	//
	// /abc
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetDown() *int32 {
	return s.Down
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetType() *string {
	return s.Type
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetUp() *int32 {
	return s.Up
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GetUri() *string {
	return s.Uri
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetDomain(v string) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetDown(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetInterval(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetPort(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetTimeout(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetType(v string) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetUp(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Up = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetUri(v string) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Uri = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) Validate() error {
	return dara.Validate(s)
}
