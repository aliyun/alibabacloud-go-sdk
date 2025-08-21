// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAIProtectModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyWebAIProtectModeRequest
	GetConfig() *string
	SetDomain(v string) *ModifyWebAIProtectModeRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyWebAIProtectModeRequest
	GetResourceGroupId() *string
}

type ModifyWebAIProtectModeRequest struct {
	// The details of the Intelligent Protection policy. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **AiTemplate**: the level of the Intelligent Protection policy. This field is required and must be of the STRING type. Valid values:
	//
	//     	- **level30**: the Low level
	//
	//     	- **level60**: the Normal level
	//
	//     	- **level90**: the Strict level
	//
	// 	- **AiMode**: the mode of the Intelligent Protection policy. This field is required and must be of the string type. Valid values:
	//
	//     	- **watch**: the Warning mode
	//
	//     	- **defense**: the Defense mode
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AiTemplate":"level60","AiMode":"defense"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/474212.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyWebAIProtectModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAIProtectModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectModeRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyWebAIProtectModeRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebAIProtectModeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebAIProtectModeRequest) SetConfig(v string) *ModifyWebAIProtectModeRequest {
	s.Config = &v
	return s
}

func (s *ModifyWebAIProtectModeRequest) SetDomain(v string) *ModifyWebAIProtectModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAIProtectModeRequest) SetResourceGroupId(v string) *ModifyWebAIProtectModeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAIProtectModeRequest) Validate() error {
	return dara.Validate(s)
}
