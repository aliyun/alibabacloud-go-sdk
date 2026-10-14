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
	// The domain name of the website whose forwarding rules you want to manage.
	//
	// > The domain name must have website forwarding rules configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names that have website forwarding rules configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The back-to-origin policy. The value is a character string that represents a JSON structure. The JSON structure contains the following fields:
	//
	//  - **ProxyMode**: String | Required | The load balancing algorithm for back-to-origin. Valid values:
	//
	//     - **ip_hash**: IP Hash algorithm. Requests from the same source IP address are directed to the same origin server based on hash mapping.
	//
	//     - **rr**: Polling algorithm. All requests are allocated to different origin servers in turn. When you use this algorithm, you can set weights for servers based on their performance.
	//
	//      - **least_time**: Least Time algorithm. This algorithm uses intelligent DNS resolution to ensure the lowest latency for service traffic across the entire path from the protection node to forward back to the origin server.
	//
	// - **Attributes**: JSONArray | Optional | The back-to-origin parameter information. Each element in the array contains the following fields:
	//
	//     - **RealServer**: String | Optional | The address of the origin server.
	//
	//     - **Attribute**: JSONObject | Optional | The back-to-origin parameters. This field contains the following fields:
	//
	//         - **Weight**: Integer | Optional | The weight of the server. This property takes effect only when the polling algorithm is used (**ProxyMode*	- is set to **rr**). Valid values: **1*	- to **100**. Default value: **100**. A server with a higher weight receives more requests.
	//
	//         - **ConnectTimeout**: Integer | Optional | The timeout period for new connections. Valid values: **1*	- to **10**. Unit: seconds. Default value: **5**.
	//
	//         - **FailTimeout**: Integer | Optional | The failed timeout period. Valid values: **1*	- to **3600**. Unit: seconds. Default value: **10**.
	//
	//         - **MaxFails**: Integer | Optional | The maximum number of failures for health checks. Valid values: **1*	- to **10**. Default value: **3**.
	//
	//         - **Mode**: String | Optional | The primary/secondary property flag. Valid values: **active*	- indicates primary. **backup*	- indicates secondary.
	//
	//         - **ReadTimeout**: Integer | Optional | The read timeout period. Valid values: **10*	- to **300**. Unit: seconds. Default value: **120**.
	//
	//         - **SendTimeout**: Integer | Optional | The write timeout period. Valid values: **10*	- to **300**. Unit: seconds. Default value: **120**.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ProxyMode":"rr","Attributes":[{"RealServer":"1.***.***.1","Attribute":{"Weight":100,"ConnectTimeout":5,"FailTimeout":10,"MaxFails":3,"Mode":"active","ReadTimeout":120,"SendTimeout":120}},{"RealServer":"2.***.***.2","Attribute":{"Weight":100,"ConnectTimeout":5,"FailTimeout":10,"MaxFails":3,"Mode":"active","ReadTimeout":120,"SendTimeout":120}}]}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to enable back-to-origin retry. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
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
