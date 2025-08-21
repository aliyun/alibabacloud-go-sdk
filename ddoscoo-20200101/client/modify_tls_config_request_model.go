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
	// The details of the TLS policy. The value is a JSON string that contains the following fields:
	//
	// 	- **ssl_protocols**: the version of TLS. This field is required. Data type: string. Valid values:
	//
	//     	- **tls1.0**: TLS 1.0 and later
	//
	//     	- **tls1.1**: TLS 1.1 and later
	//
	//     	- **tls1.2**: TLS 1.2 and later
	//
	// 	- **ssl_ciphers**: the type of the cipher suite. This field is required. Data type: string. Valid values:
	//
	//     	- **all**: all cipher suites, which include strong and weak cipher suites
	//
	//     	- **improved**: enhanced cipher suites
	//
	//     	- **strong**: strong cipher suites
	//
	//     	- **default**: default cipher suites, which include only strong cipher suites
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ssl_protocols":"tls1.0","ssl_ciphers":"all"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
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
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
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
