// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckStatusListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckStatusList(v []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) *DescribeHealthCheckStatusListResponseBody
	GetHealthCheckStatusList() []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList
	SetRequestId(v string) *DescribeHealthCheckStatusListResponseBody
	GetRequestId() *string
}

type DescribeHealthCheckStatusListResponseBody struct {
	HealthCheckStatusList []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList `json:"HealthCheckStatusList,omitempty" xml:"HealthCheckStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBody) GetHealthCheckStatusList() []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	return s.HealthCheckStatusList
}

func (s *DescribeHealthCheckStatusListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthCheckStatusListResponseBody) SetHealthCheckStatusList(v []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) *DescribeHealthCheckStatusListResponseBody {
	s.HealthCheckStatusList = v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBody) SetRequestId(v string) *DescribeHealthCheckStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList struct {
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// tcp
	Protocol             *string                                                                               `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServerStatusList []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList `json:"RealServerStatusList,omitempty" xml:"RealServerStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GetRealServerStatusList() []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	return s.RealServerStatusList
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetFrontendPort(v int32) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetInstanceId(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetProtocol(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetRealServerStatusList(v []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.RealServerStatusList = v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList struct {
	// example:
	//
	// 1.1.1.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) GetAddress() *string {
	return s.Address
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) SetAddress(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	s.Address = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) Validate() error {
	return dara.Validate(s)
}
