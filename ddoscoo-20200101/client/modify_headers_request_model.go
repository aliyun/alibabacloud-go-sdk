// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHeadersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHeaders(v string) *ModifyHeadersRequest
	GetCustomHeaders() *string
	SetDomain(v string) *ModifyHeadersRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyHeadersRequest
	GetResourceGroupId() *string
}

type ModifyHeadersRequest struct {
	// The key-value pair of the custom header. The key specifies the header name, and the value specifies the header value. You can specify up to five key-value pairs. The key-value pairs can be up to 200 characters in length.
	//
	// Take note of the following items:
	//
	// 	- Do not use the following default HTTP headers:
	//
	//     	- X-Forwarded-ClientSrcPort: This header is used to obtain the source ports of clients that access Anti-DDoS Proxy (a Layer 7 proxy).
	//
	//     	- X-Forwarded-ProxyPort: This header is used to obtain the ports of listeners that access Anti-DDoS Proxy (a Layer 7 proxy).
	//
	//     	- X-Forwarded-For: This header is used to obtain the IP addresses of clients that access Anti-DDoS Proxy (a Layer 7 proxy).
	//
	// 	- Do not use standard HTTP headers or specific widely used custom HTTP headers. The standard HTTP headers include Host, User-Agent, Connection, and Upgrade, and the widely used custom HTTP headers include X-Real-IP, X-True-IP, X-Client-IP, Web-Server-Type, WL-Proxy-Client-IP, eEagleEye-RpcID, EagleEye-TraceID, X-Forwarded-Cluster, and X-Forwarded-Proto. If the preceding headers are used, the original content of the headers is overwritten.
	//
	// >  If you specify a key of X-Forwarded-ClientSrcPort, the system obtains the originating ports of clients that access Anti-DDoS Proxy (a Layer 7 proxy). In this case, the value is an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"X-Forwarded-ClientSrcPort\\":\\"\\"}
	CustomHeaders *string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// >
	//
	// 	- You can query resource group IDs in the Anti-DDoS Pro or Anti-DDoS Premium console or by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// 	- Before you modify the resource group to which an instance belongs, you can call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the current resource group of the instance.
	//
	// example:
	//
	// rg-acfmz6jbof5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyHeadersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHeadersRequest) GoString() string {
	return s.String()
}

func (s *ModifyHeadersRequest) GetCustomHeaders() *string {
	return s.CustomHeaders
}

func (s *ModifyHeadersRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyHeadersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyHeadersRequest) SetCustomHeaders(v string) *ModifyHeadersRequest {
	s.CustomHeaders = &v
	return s
}

func (s *ModifyHeadersRequest) SetDomain(v string) *ModifyHeadersRequest {
	s.Domain = &v
	return s
}

func (s *ModifyHeadersRequest) SetResourceGroupId(v string) *ModifyHeadersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyHeadersRequest) Validate() error {
	return dara.Validate(s)
}
