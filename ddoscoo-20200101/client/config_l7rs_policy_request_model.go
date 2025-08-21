// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7RsPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigL7RsPolicyRequest
	GetDomain() *string
	SetPolicy(v string) *ConfigL7RsPolicyRequest
	GetPolicy() *string
	SetResourceGroupId(v string) *ConfigL7RsPolicyRequest
	GetResourceGroupId() *string
	SetUpstreamRetry(v int32) *ConfigL7RsPolicyRequest
	GetUpstreamRetry() *int32
}

type ConfigL7RsPolicyRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query the domain names for which forwarding rules are configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The back-to-origin policy. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **ProxyMode**: The load balancing algorithm for back-to-origin traffic. This field is required and must be a string. Valid values:
	//
	//     	- **ip_hash**: the IP hash algorithm. This algorithm is used to redirect requests from the same IP address to the same origin server.
	//
	//     	- **rr**: the round-robin algorithm. This algorithm is used to redirect requests to origin servers in turn. If you use this algorithm, you can specify a weight for each server based on server performance.
	//
	//     	- **least_time**: the least response time algorithm. This algorithm is used to minimize the latency when requests are forwarded from the instance to origin servers based on the intelligent DNS resolution feature.
	//
	// 	- **Attributes**: the parameters for back-to-origin processing. This field is optional and must be a JSON array. Each element in the array contains the following fields:
	//
	//     	- **RealServer**: the address of the origin server. This field is optional and must be a string.
	//
	//     	- **Attribute**: the parameter for back-to-origin processing. This field is optional and must be a JSON object. Valid values:
	//
	//         	- **Weight**: the weight of the server. This field is optional and must be an integer. This field takes effect only when **ProxyMode*	- is set to **rr**. Valid values: **1*	- to **100**. Default value: **100**. An origin server with a higher weight receives more requests.
	//
	//         	- **ConnectTimeout**: the timeout period for new connections. This field is optional and must be an integer. Valid values: **1*	- to **10**. Unit: seconds. Default value: **5**.
	//
	//         	- **FailTimeout**: the period after which a connection is considered to have failed. This field is optional and must be an integer. Valid values: **1*	- to **3600**. Unit: seconds. Default value: **10**.
	//
	//         	- **MaxFails**: the maximum number of failures allowed. This field is related to health checks. This field is optional and must be an integer. Valid values: **1*	- to **10**. Unit: seconds. Default value: **3**.
	//
	//         	- **Mode**: the primary/secondary attribute flag. This parameter is optional and must be a string. Valid values: **active*	- (primary) and **backup*	- (secondary).
	//
	//         	- **ReadTimeout**: the read timeout period. This field is optional and must be an integer. Valid values: **10*	- to **300**. Unit: seconds. Default value: **120**.
	//
	//         	- **SendTimeout**: the write timeout period. This field is optional and must be an integer. Valid values: **10*	- to **300**. Unit: seconds. Default value: **120**.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ProxyMode":"rr","Attributes":[{"RealServer":"1.***.***.1","Attribute":{"Weight":100}},{"RealServer":"2.***.***.2","Attribute":{"Weight":100}}]}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The retry switch. Valid values:
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

func (s ConfigL7RsPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7RsPolicyRequest) GoString() string {
	return s.String()
}

func (s *ConfigL7RsPolicyRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigL7RsPolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ConfigL7RsPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigL7RsPolicyRequest) GetUpstreamRetry() *int32 {
	return s.UpstreamRetry
}

func (s *ConfigL7RsPolicyRequest) SetDomain(v string) *ConfigL7RsPolicyRequest {
	s.Domain = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) SetPolicy(v string) *ConfigL7RsPolicyRequest {
	s.Policy = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) SetResourceGroupId(v string) *ConfigL7RsPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) SetUpstreamRetry(v int32) *ConfigL7RsPolicyRequest {
	s.UpstreamRetry = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) Validate() error {
	return dara.Validate(s)
}
