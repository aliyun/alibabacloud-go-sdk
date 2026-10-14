// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTlsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyTlsConfigRequest
	GetConfig() *string
	SetDomain(v string) *ModifyTlsConfigRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyTlsConfigRequest
	GetResourceGroupId() *string
}

type ModifyTlsConfigRequest struct {
	// The details of the TLS security policy, expressed as a JSON-formatted string with the following structure.
	//
	// - **ssl_protocols**: String type, required. The TLS version. Valid values:
	//
	//     - **tls1.0**: TLS 1.0 and later are supported.
	//
	//     - **tls1.1**: TLS 1.1 and later are supported.
	//
	//     - **tls1.2**: TLS 1.2 and later are supported.
	//
	// - **ssl_ciphers**: String type, required. The cipher suite type. Valid values:
	//
	//     - **all**: all cipher suites, including strong and weak cipher suites
	//
	//     - **improved**: enhanced cipher suites
	//
	//     - **strong**: strong cipher suites only
	//
	//     - **default**: default cipher suites, which include only strong cipher suites
	//
	// - **ssl_13_enabled**: Boolean type, optional. Specifies whether to enable TLS 1.3. Valid values:
	//
	//     - **true**: Enables TLS 1.3.
	//
	//     - **false**: Disables TLS 1.3.
	//
	// - **custom_ciphers**: Array type, optional. The list of custom cipher suites for TLS versions earlier than 1.3.
	//
	// - **tls_13_custom_ciphers**: Array type, optional. The list of custom cipher suites for TLS 1.3. If TLS 1.3 is enabled, select at least one cipher suite for TLS versions earlier than 1.3.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "ssl_protocols": "tls1.2",
	//
	//   "ssl_ciphers": "default",
	//
	//   "ssl_13_enabled": true,
	//
	//   "custom_ciphers": [
	//
	//     "ECDHE-ECDSA-AES128-GCM-SHA256"
	//
	//   ],
	//
	//   "tls_13_custom_ciphers": [
	//
	//     "TLS_AES_256_GCM_SHA384"
	//
	//   ]
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name of the website.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyTlsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTlsConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTlsConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyTlsConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyTlsConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyTlsConfigRequest) SetConfig(v string) *ModifyTlsConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyTlsConfigRequest) SetDomain(v string) *ModifyTlsConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyTlsConfigRequest) SetResourceGroupId(v string) *ModifyTlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyTlsConfigRequest) Validate() error {
	return dara.Validate(s)
}
