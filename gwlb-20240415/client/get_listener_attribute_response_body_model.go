// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListenerDescription(v string) *GetListenerAttributeResponseBody
	GetListenerDescription() *string
	SetListenerId(v string) *GetListenerAttributeResponseBody
	GetListenerId() *string
	SetListenerStatus(v string) *GetListenerAttributeResponseBody
	GetListenerStatus() *string
	SetLoadBalancerId(v string) *GetListenerAttributeResponseBody
	GetLoadBalancerId() *string
	SetRegionId(v string) *GetListenerAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetListenerAttributeResponseBody
	GetRequestId() *string
	SetServerGroupId(v string) *GetListenerAttributeResponseBody
	GetServerGroupId() *string
	SetTags(v []*GetListenerAttributeResponseBodyTags) *GetListenerAttributeResponseBody
	GetTags() []*GetListenerAttributeResponseBodyTags
	SetTcpIdleTimeout(v int32) *GetListenerAttributeResponseBody
	GetTcpIdleTimeout() *int32
}

type GetListenerAttributeResponseBody struct {
	// The listener description.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// listener_description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsn-3kbj3587mqhm3p****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listener status. Valid values:
	//
	// 	- **Provisioning**: The listener is being created.
	//
	// 	- **Running**: The listener is running.
	//
	// 	- **Configuring**: The listener is being configured.
	//
	// 	- **Deleting**: The listener is being deleted.
	//
	// example:
	//
	// Provisioning
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-te609d6696632f76****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The region ID of the GWLB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75CC3312-7757-5EE1-90D8-49CEA66608AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-sp8d2r6y7t0xtl****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The tags.
	Tags []*GetListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of an idle TCP connection. Unit: seconds.
	//
	// example:
	//
	// 350
	TcpIdleTimeout *int32 `json:"TcpIdleTimeout,omitempty" xml:"TcpIdleTimeout,omitempty"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *GetListenerAttributeResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerAttributeResponseBody) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *GetListenerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetListenerAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListenerAttributeResponseBody) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *GetListenerAttributeResponseBody) GetTags() []*GetListenerAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetListenerAttributeResponseBody) GetTcpIdleTimeout() *int32 {
	return s.TcpIdleTimeout
}

func (s *GetListenerAttributeResponseBody) SetListenerDescription(v string) *GetListenerAttributeResponseBody {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerId(v string) *GetListenerAttributeResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerStatus(v string) *GetListenerAttributeResponseBody {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLoadBalancerId(v string) *GetListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRegionId(v string) *GetListenerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetServerGroupId(v string) *GetListenerAttributeResponseBody {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetTags(v []*GetListenerAttributeResponseBodyTags) *GetListenerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetTcpIdleTimeout(v int32) *GetListenerAttributeResponseBody {
	s.TcpIdleTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetListenerAttributeResponseBodyTags struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs: `or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetListenerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetListenerAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetListenerAttributeResponseBodyTags) SetKey(v string) *GetListenerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetListenerAttributeResponseBodyTags) SetValue(v string) *GetListenerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetListenerAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
