// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7RsPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v []*DescribeL7RsPolicyResponseBodyAttributes) *DescribeL7RsPolicyResponseBody
	GetAttributes() []*DescribeL7RsPolicyResponseBodyAttributes
	SetProxyMode(v string) *DescribeL7RsPolicyResponseBody
	GetProxyMode() *string
	SetRequestId(v string) *DescribeL7RsPolicyResponseBody
	GetRequestId() *string
	SetRsAttrRwTimeoutMax(v int64) *DescribeL7RsPolicyResponseBody
	GetRsAttrRwTimeoutMax() *int64
	SetUpstreamRetry(v int32) *DescribeL7RsPolicyResponseBody
	GetUpstreamRetry() *int32
}

type DescribeL7RsPolicyResponseBody struct {
	// The details about the parameters for back-to-origin settings.
	Attributes []*DescribeL7RsPolicyResponseBodyAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The scheduling algorithm for back-to-origin traffic. Valid values:
	//
	// 	- **ip_hash**: the IP hash algorithm. This algorithm is used to redirect the requests from the same IP address to the same origin server.
	//
	// 	- **rr**: the round-robin algorithm. This algorithm is used to redirect requests to origin servers in turn.
	//
	// 	- **least_time**: the least response time algorithm. This algorithm is used to minimize the latency when requests are forwarded from Anti-DDoS Pro or Anti-DDoS Premium instances to origin servers based on the intelligent DNS resolution feature.
	//
	// example:
	//
	// rr
	ProxyMode *string `json:"ProxyMode,omitempty" xml:"ProxyMode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9E7F6B2C-03F2-462F-9076-B782CF0DD502
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period for a read or write connection.
	//
	// example:
	//
	// 300
	RsAttrRwTimeoutMax *int64 `json:"RsAttrRwTimeoutMax,omitempty" xml:"RsAttrRwTimeoutMax,omitempty"`
	// The back-to-origin retry switch. Valid values:
	//
	// 	- **1**: on
	//
	// 	- **0**: off
	//
	// example:
	//
	// 1
	UpstreamRetry *int32 `json:"UpstreamRetry,omitempty" xml:"UpstreamRetry,omitempty"`
}

func (s DescribeL7RsPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7RsPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponseBody) GetAttributes() []*DescribeL7RsPolicyResponseBodyAttributes {
	return s.Attributes
}

func (s *DescribeL7RsPolicyResponseBody) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *DescribeL7RsPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeL7RsPolicyResponseBody) GetRsAttrRwTimeoutMax() *int64 {
	return s.RsAttrRwTimeoutMax
}

func (s *DescribeL7RsPolicyResponseBody) GetUpstreamRetry() *int32 {
	return s.UpstreamRetry
}

func (s *DescribeL7RsPolicyResponseBody) SetAttributes(v []*DescribeL7RsPolicyResponseBodyAttributes) *DescribeL7RsPolicyResponseBody {
	s.Attributes = v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) SetProxyMode(v string) *DescribeL7RsPolicyResponseBody {
	s.ProxyMode = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) SetRequestId(v string) *DescribeL7RsPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) SetRsAttrRwTimeoutMax(v int64) *DescribeL7RsPolicyResponseBody {
	s.RsAttrRwTimeoutMax = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) SetUpstreamRetry(v int32) *DescribeL7RsPolicyResponseBody {
	s.UpstreamRetry = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeL7RsPolicyResponseBodyAttributes struct {
	// The parameters for back-to-origin settings.
	Attribute *DescribeL7RsPolicyResponseBodyAttributesAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	// The address of the origin server.
	//
	// example:
	//
	// 1.***.***.1
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
	// The address type of the origin server. Valid values:
	//
	// 	- **0**: IP address
	//
	// 	- **1**: domain name
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s DescribeL7RsPolicyResponseBodyAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7RsPolicyResponseBodyAttributes) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) GetAttribute() *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	return s.Attribute
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) GetRealServer() *string {
	return s.RealServer
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) GetRsType() *int32 {
	return s.RsType
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) SetAttribute(v *DescribeL7RsPolicyResponseBodyAttributesAttribute) *DescribeL7RsPolicyResponseBodyAttributes {
	s.Attribute = v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) SetRealServer(v string) *DescribeL7RsPolicyResponseBodyAttributes {
	s.RealServer = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) SetRsType(v int32) *DescribeL7RsPolicyResponseBodyAttributes {
	s.RsType = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) Validate() error {
	if s.Attribute != nil {
		if err := s.Attribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeL7RsPolicyResponseBodyAttributesAttribute struct {
	// The timeout period for a new connection. Valid values: **1*	- to **10**. Unit: seconds. Default value: **5**.
	//
	// example:
	//
	// 5
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// The expiration time of a connection, in seconds. If the number of failures at the origin server exceeds the **MaxFails*	- value, the address of the origin server is set to down and the expiration time is **FailTimeout**. The final value is the maximum value of **ConnectTimeout*	- and **FailTimeout**. Valid values: **1*	- to **3600**. Unit: seconds. Default value: **10**.
	//
	// example:
	//
	// 10
	FailTimeout *int32 `json:"FailTimeout,omitempty" xml:"FailTimeout,omitempty"`
	// The maximum number of failures. This parameter is related to health check. Valid values: **1*	- to **10**. Unit: seconds. Default value: **3**.
	//
	// example:
	//
	// 3
	MaxFails *int32 `json:"MaxFails,omitempty" xml:"MaxFails,omitempty"`
	// The primary/secondary flag. Valid values:
	//
	// 	- **active**: primary
	//
	// 	- **backup**: secondary
	//
	// example:
	//
	// active
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The timeout period for a read connection. Valid values: **10*	- to **300**. Unit: seconds. Default value: **120**.
	//
	// example:
	//
	// 120
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The timeout period for a write connection. Valid values: **10*	- to **300**. Unit: seconds. Default value: **120**.
	//
	// example:
	//
	// 120
	SendTimeout *int32 `json:"SendTimeout,omitempty" xml:"SendTimeout,omitempty"`
	// The weight of the origin server. This parameter takes effect only if the value of **ProxyMode*	- is **rr*	- or **ip_hash**.****
	//
	// Valid values: **1*	- to **100**. Default value: **100**. A server with a higher weight receives more requests.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeL7RsPolicyResponseBodyAttributesAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7RsPolicyResponseBodyAttributesAttribute) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetFailTimeout() *int32 {
	return s.FailTimeout
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetMaxFails() *int32 {
	return s.MaxFails
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetMode() *string {
	return s.Mode
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetSendTimeout() *int32 {
	return s.SendTimeout
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetConnectTimeout(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetFailTimeout(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.FailTimeout = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetMaxFails(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.MaxFails = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetMode(v string) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.Mode = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetReadTimeout(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetSendTimeout(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.SendTimeout = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetWeight(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.Weight = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) Validate() error {
	return dara.Validate(s)
}
