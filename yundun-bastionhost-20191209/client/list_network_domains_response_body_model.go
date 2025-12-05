// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkDomains(v []*ListNetworkDomainsResponseBodyNetworkDomains) *ListNetworkDomainsResponseBody
	GetNetworkDomains() []*ListNetworkDomainsResponseBodyNetworkDomains
	SetRequestId(v string) *ListNetworkDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListNetworkDomainsResponseBody
	GetTotalCount() *int64
}

type ListNetworkDomainsResponseBody struct {
	// The network domains that are returned.
	NetworkDomains []*ListNetworkDomainsResponseBodyNetworkDomains `json:"NetworkDomains,omitempty" xml:"NetworkDomains,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of network domains that are returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkDomainsResponseBody) GetNetworkDomains() []*ListNetworkDomainsResponseBodyNetworkDomains {
	return s.NetworkDomains
}

func (s *ListNetworkDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListNetworkDomainsResponseBody) SetNetworkDomains(v []*ListNetworkDomainsResponseBodyNetworkDomains) *ListNetworkDomainsResponseBody {
	s.NetworkDomains = v
	return s
}

func (s *ListNetworkDomainsResponseBody) SetRequestId(v string) *ListNetworkDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkDomainsResponseBody) SetTotalCount(v int64) *ListNetworkDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkDomainsResponseBody) Validate() error {
	if s.NetworkDomains != nil {
		for _, item := range s.NetworkDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkDomainsResponseBodyNetworkDomains struct {
	// The remarks of the network domain.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether the network domain is built-in.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Default *bool `json:"Default,omitempty" xml:"Default,omitempty"`
	// The network domain ID.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The name of the network domain.
	//
	// example:
	//
	// test
	NetworkDomainName *string `json:"NetworkDomainName,omitempty" xml:"NetworkDomainName,omitempty"`
	// The connection mode of the network domain. Valid values:
	//
	// 	- **Direct**
	//
	// 	- **Proxy**
	//
	// example:
	//
	// Proxy
	NetworkDomainType *string                                                     `json:"NetworkDomainType,omitempty" xml:"NetworkDomainType,omitempty"`
	ProxiesState      []*ListNetworkDomainsResponseBodyNetworkDomainsProxiesState `json:"ProxiesState,omitempty" xml:"ProxiesState,omitempty" type:"Repeated"`
}

func (s ListNetworkDomainsResponseBodyNetworkDomains) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkDomainsResponseBodyNetworkDomains) GoString() string {
	return s.String()
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) GetComment() *string {
	return s.Comment
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) GetDefault() *bool {
	return s.Default
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) GetNetworkDomainName() *string {
	return s.NetworkDomainName
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) GetNetworkDomainType() *string {
	return s.NetworkDomainType
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) GetProxiesState() []*ListNetworkDomainsResponseBodyNetworkDomainsProxiesState {
	return s.ProxiesState
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) SetComment(v string) *ListNetworkDomainsResponseBodyNetworkDomains {
	s.Comment = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) SetDefault(v bool) *ListNetworkDomainsResponseBodyNetworkDomains {
	s.Default = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) SetNetworkDomainId(v string) *ListNetworkDomainsResponseBodyNetworkDomains {
	s.NetworkDomainId = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) SetNetworkDomainName(v string) *ListNetworkDomainsResponseBodyNetworkDomains {
	s.NetworkDomainName = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) SetNetworkDomainType(v string) *ListNetworkDomainsResponseBodyNetworkDomains {
	s.NetworkDomainType = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) SetProxiesState(v []*ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) *ListNetworkDomainsResponseBodyNetworkDomains {
	s.ProxiesState = v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomains) Validate() error {
	if s.ProxiesState != nil {
		for _, item := range s.ProxiesState {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkDomainsResponseBodyNetworkDomainsProxiesState struct {
	// The node type of the proxy server. Valid values:
	//
	// 	- **Master**: primary proxy server.
	//
	// 	- **Slave**: secondary proxy server.
	//
	// example:
	//
	// Master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The status of the proxy server.
	//
	// 	- **Available**
	//
	// 	- **Unavailable**
	//
	// example:
	//
	// Available
	ProxyState *string `json:"ProxyState,omitempty" xml:"ProxyState,omitempty"`
}

func (s ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) GoString() string {
	return s.String()
}

func (s *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) GetNodeType() *string {
	return s.NodeType
}

func (s *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) GetProxyState() *string {
	return s.ProxyState
}

func (s *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) SetNodeType(v string) *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState {
	s.NodeType = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) SetProxyState(v string) *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState {
	s.ProxyState = &v
	return s
}

func (s *ListNetworkDomainsResponseBodyNetworkDomainsProxiesState) Validate() error {
	return dara.Validate(s)
}
