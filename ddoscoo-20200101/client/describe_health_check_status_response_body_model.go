// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckStatus(v []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus) *DescribeHealthCheckStatusResponseBody
	GetHealthCheckStatus() []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus
	SetRequestId(v string) *DescribeHealthCheckStatusResponseBody
	GetRequestId() *string
}

type DescribeHealthCheckStatusResponseBody struct {
	// An array that consists of the details of the health status of the origin server.
	HealthCheckStatus []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus `json:"HealthCheckStatus,omitempty" xml:"HealthCheckStatus,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DE9FF9E1-569C-4B6C-AB6A-0F6D927BB27C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponseBody) GetHealthCheckStatus() []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	return s.HealthCheckStatus
}

func (s *DescribeHealthCheckStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthCheckStatusResponseBody) SetHealthCheckStatus(v []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus) *DescribeHealthCheckStatusResponseBody {
	s.HealthCheckStatus = v
	return s
}

func (s *DescribeHealthCheckStatusResponseBody) SetRequestId(v string) *DescribeHealthCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthCheckStatusResponseBodyHealthCheckStatus struct {
	// The forwarding port.
	//
	// example:
	//
	// 8080
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
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
	// An array that consists of the health states of the IP addresses of the origin server.
	RealServerStatusList []*DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList `json:"RealServerStatusList,omitempty" xml:"RealServerStatusList,omitempty" type:"Repeated"`
	// The health status of the origin server. Valid values:
	//
	// 	- **normal**: healthy
	//
	// 	- **abnormal**: unhealthy
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GetRealServerStatusList() []*DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList {
	return s.RealServerStatusList
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetFrontendPort(v int32) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetInstanceId(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetProtocol(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetRealServerStatusList(v []*DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.RealServerStatusList = v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetStatus(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList struct {
	// The IP address of the origin server.
	//
	// example:
	//
	// 192.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The health state of the IP address. Valid values:
	//
	// 	- **normal**: healthy
	//
	// 	- **abnormal**: unhealthy
	//
	// example:
	//
	// abnormal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) GetAddress() *string {
	return s.Address
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) SetAddress(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList {
	s.Address = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) SetStatus(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) Validate() error {
	return dara.Validate(s)
}
