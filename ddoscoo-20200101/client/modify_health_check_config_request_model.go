// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHealthCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardProtocol(v string) *ModifyHealthCheckConfigRequest
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *ModifyHealthCheckConfigRequest
	GetFrontendPort() *int32
	SetHealthCheck(v string) *ModifyHealthCheckConfigRequest
	GetHealthCheck() *string
	SetInstanceId(v string) *ModifyHealthCheckConfigRequest
	GetInstanceId() *string
}

type ModifyHealthCheckConfigRequest struct {
	// The forwarding protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// The forwarding port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The details of the health check configuration. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **Type**: the protocol type. This field is required and must be of the STRING type. Valid values: **tcp*	- (Layer 4) and **http*	- (Layer 7).
	//
	// 	- **Domain**: the domain name, which must be of the STRING type.
	//
	//     **
	//
	//     **Note**This parameter is returned only when the Layer 7 health check configuration is queried.
	//
	// 	- **Uri**: the check path, which must be of the STRING type.
	//
	//     **
	//
	//     **Note**This parameter is returned only when the Layer 7 health check configuration is queried.
	//
	// 	- **Timeout**: the response timeout period, which must be of the INTEGER type. Valid values: **1*	- to **30**. Unit: seconds.
	//
	// 	- **Port**: the port on which you want to perform the health check, which must be of the INTEGER type.
	//
	// 	- **Interval**: the health check interval, which must be of the INTEGER type. Valid values: **1*	- to **30**. Unit: seconds.
	//
	// 	- **Up**: the number of consecutive successful health checks that must occur before declaring a port healthy, which must be of the INTEGER type. Valid values: **1*	- to **10**.
	//
	// 	- **Down**: the number of consecutive failed health checks that must occur before declaring a port unhealthy, which must be of the INTEGER type. Valid values: **1*	- to **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Type":"tcp","Timeout":10,"Port":8080,"Interval":10,"Up":10,"Down":40}
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyHealthCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHealthCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigRequest) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *ModifyHealthCheckConfigRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *ModifyHealthCheckConfigRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *ModifyHealthCheckConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHealthCheckConfigRequest) SetForwardProtocol(v string) *ModifyHealthCheckConfigRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetFrontendPort(v int32) *ModifyHealthCheckConfigRequest {
	s.FrontendPort = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetHealthCheck(v string) *ModifyHealthCheckConfigRequest {
	s.HealthCheck = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetInstanceId(v string) *ModifyHealthCheckConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
