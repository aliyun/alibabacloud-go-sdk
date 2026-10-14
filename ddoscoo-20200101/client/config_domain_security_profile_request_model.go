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
	// Not in use.
	//
	// example:
	//
	// default
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The Anti-DDoS Global Mitigation Policy for Protection for Website Services. The structure is as follows:
	//
	// - **global_rule_mode**: String type. Optional. The protection level of the DDoS global mitigation feature. Valid values:
	//
	//     - **weak**: Loose.
	//
	//     - **default**: Normal.
	//
	//     - **hard**: Strict.
	//
	// - **global_rule_enable**: Integer type. Optional. The switch of the DDoS global mitigation feature. Valid values:
	//
	//     - **0**: Shutdown.
	//
	//     - **1**: Enabled.
	//
	// The Blacklist/Whitelist feature of the custom mitigation policies for Protection for Website Services. The structure is as follows:
	//
	// - **bwlist_enable**: Integer type. Optional. The switch of the custom mitigation policy blacklist/whitelist. Valid values:
	//
	//     - **0**: Shutdown.
	//
	//     - **1**: Enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"global_rule_mode":"hard"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name of the website service.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
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
