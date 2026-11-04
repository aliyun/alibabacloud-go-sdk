// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceInternetProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotInstanceId(v string) *UpdateInstanceInternetProtocolRequest
	GetIotInstanceId() *string
	SetIpVersion(v string) *UpdateInstanceInternetProtocolRequest
	GetIpVersion() *string
}

type UpdateInstanceInternetProtocolRequest struct {
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If an ID value is available, you must specify it. Otherwise, the call fails.
	//
	// - If the <props="intl"><ph>**Instance overview*	- page or </ph>ID value is not available, you do not need to specify it.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The IP protocol version. Valid values:
	//
	// - IPv4 (default)
	//
	// - IPv6
	//
	// > If the IP protocol version is set to IPv6, streaming media-related addresses are changed to IPv4/IPv6 dual-stack addresses. Make sure that your devices and applications are configured for IPv6. Otherwise, exceptions such as playback screen loading failures may occur.
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s UpdateInstanceInternetProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInternetProtocolRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInternetProtocolRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *UpdateInstanceInternetProtocolRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *UpdateInstanceInternetProtocolRequest) SetIotInstanceId(v string) *UpdateInstanceInternetProtocolRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateInstanceInternetProtocolRequest) SetIpVersion(v string) *UpdateInstanceInternetProtocolRequest {
	s.IpVersion = &v
	return s
}

func (s *UpdateInstanceInternetProtocolRequest) Validate() error {
	return dara.Validate(s)
}
