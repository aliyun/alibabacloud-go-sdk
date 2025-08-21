// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigDomainSecurityProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *ConfigDomainSecurityProfileRequest
	GetCluster() *string
	SetConfig(v string) *ConfigDomainSecurityProfileRequest
	GetConfig() *string
	SetDomain(v string) *ConfigDomainSecurityProfileRequest
	GetDomain() *string
}

type ConfigDomainSecurityProfileRequest struct {
	// This parameter is unavailable.
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The configurations for the global mitigation policy feature. The configurations include the following fields:
	//
	// 	- **global_rule_mode**: optional. The mode for the global mitigation policy feature. Data type: string. Valid values:
	//
	//     	- **weak**: loose.
	//
	//     	- **default**: normal.
	//
	//     	- **hard**: strict.
	//
	// 	- **global_rule_enable**: optional. Specifies whether to enable the global mitigation policy feature. Data type: string. Valid values:
	//
	//     	- **0**: disabled.
	//
	//     	- **1**: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"global_rule_mode\\":\\"hard\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// live.abcde.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s ConfigDomainSecurityProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigDomainSecurityProfileRequest) GoString() string {
	return s.String()
}

func (s *ConfigDomainSecurityProfileRequest) GetCluster() *string {
	return s.Cluster
}

func (s *ConfigDomainSecurityProfileRequest) GetConfig() *string {
	return s.Config
}

func (s *ConfigDomainSecurityProfileRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigDomainSecurityProfileRequest) SetCluster(v string) *ConfigDomainSecurityProfileRequest {
	s.Cluster = &v
	return s
}

func (s *ConfigDomainSecurityProfileRequest) SetConfig(v string) *ConfigDomainSecurityProfileRequest {
	s.Config = &v
	return s
}

func (s *ConfigDomainSecurityProfileRequest) SetDomain(v string) *ConfigDomainSecurityProfileRequest {
	s.Domain = &v
	return s
}

func (s *ConfigDomainSecurityProfileRequest) Validate() error {
	return dara.Validate(s)
}
