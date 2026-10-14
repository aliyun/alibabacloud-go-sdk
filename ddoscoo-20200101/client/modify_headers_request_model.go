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
	SetEmbeddedHeaders(v string) *ModifyHeadersRequest
	GetEmbeddedHeaders() *string
	SetResourceGroupId(v string) *ModifyHeadersRequest
	GetResourceGroupId() *string
}

type ModifyHeadersRequest struct {
	// The custom header, specified as a key-value pair. The key is the header name and the value is the header value. You can specify up to 5 key-value pairs. The total length cannot exceed 200 characters.
	//
	// Important:
	//
	// - - X-Forwarded-ClientSrcPort: Used by default to retrieve the client port that accesses the Layer 7 engine of Anti-DDoS Pro.
	//
	//   - X-Forwarded-ProxyPort: Used by default to retrieve the listening port of the Layer 7 engine of Anti-DDoS Pro.
	//
	//   - X-Forwarded-For: Used by default to retrieve the client IP address that accesses the Layer 7 engine of Anti-DDoS Pro.
	//
	// - Do not use standard HTTP header fields (such as `host`, `user-agent`, `connection`, and `upgrade`) or common custom HTTP header fields (such as `x-real-ip`, `x-true-ip`, `x-client-ip`, `web-server-type`, `wl-proxy-client-ip`, `eagleeye-rpcid`, `eagleeye-traceid`, `x-forwarded-cluster`, and `x-forwarded-proto`) as they will overwrite the original header values in requests.
	//
	// > To use the `X-Forwarded-ClientSrcPort` key to retrieve the client port that accesses the Layer 7 engine of Anti-DDoS Pro, set the value to `""`.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"X-Forwarded-ClientSrcPort":""}
	CustomHeaders *string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// The domain of the website.
	//
	// > A website service forwarding rule must be configured for the domain. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all your domains.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The forwarding configuration for request headers. This parameter is a JSON string of key-value pairs where the key is the header field name and the value is a boolean. Set the value to `true` to enable forwarding or `false` to disable it.
	//
	// Field description:
	//
	// `X-Client-IP`: Forwards the client\\"s real IP address.
	//
	// `X-True-IP`: Forwards the client\\"s connection IP address.
	//
	// `Web-Server-Type`: Informs the back-end server about the type of front-end web server or proxy that is handling the request.
	//
	// `WL-Proxy-Client-IP`: A request header specific to Oracle WebLogic Server that provides the same functionality as `X-Client-IP`.
	//
	// `X-Forwarded-Proto`: Indicates the protocol used for the connection between the client and the first proxy.
	//
	// Important:
	//
	// - **Supported Headers**: The configuration supports only the five preset header fields listed above.
	//
	// - **Default settings**: All the listed header fields are enabled by default. To make changes, you must explicitly set the corresponding value to `false`.
	//
	// example:
	//
	// {"X-Client-IP":true,"X-True-IP":true,"Web-Server-Type":true,"WL-Proxy-Client-IP":true,"X-Forwarded-Proto":true}
	EmbeddedHeaders *string `json:"EmbeddedHeaders,omitempty" xml:"EmbeddedHeaders,omitempty"`
	// The ID of the resource group.
	//
	// > - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation or use the console to obtain a list of resource group IDs. For more information, see [View basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// - To find the current resource group of an instance before changing it, call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation.
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

func (s *ModifyHeadersRequest) GetEmbeddedHeaders() *string {
	return s.EmbeddedHeaders
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

func (s *ModifyHeadersRequest) SetEmbeddedHeaders(v string) *ModifyHeadersRequest {
	s.EmbeddedHeaders = &v
	return s
}

func (s *ModifyHeadersRequest) SetResourceGroupId(v string) *ModifyHeadersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyHeadersRequest) Validate() error {
	return dara.Validate(s)
}
