// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*DescribeCdnUserConfigsResponseBodyConfigs) *DescribeCdnUserConfigsResponseBody
	GetConfigs() []*DescribeCdnUserConfigsResponseBodyConfigs
	SetRequestId(v string) *DescribeCdnUserConfigsResponseBody
	GetRequestId() *string
}

type DescribeCdnUserConfigsResponseBody struct {
	// The user configurations.
	Configs []*DescribeCdnUserConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 06D29681-B7CD-4034-A8CC-28AFFA213539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnUserConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsResponseBody) GetConfigs() []*DescribeCdnUserConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *DescribeCdnUserConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserConfigsResponseBody) SetConfigs(v []*DescribeCdnUserConfigsResponseBodyConfigs) *DescribeCdnUserConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *DescribeCdnUserConfigsResponseBody) SetRequestId(v string) *DescribeCdnUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnUserConfigsResponseBodyConfigs struct {
	// The name of the configuration.
	//
	// The configuration is specified by enterprise users and public service sectors.
	//
	// example:
	//
	// allow_function
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The value of the configuration. Valid values:
	//
	// 	- **cc_rule**: HTTP flood protection rules
	//
	// 	- **ddos_dispatch**: integration with Anti-DDoS
	//
	// 	- **edge_safe**: application security settings on POPs
	//
	// 	- **blocked_regions**: blocked regions
	//
	// 	- **http_acl_policy**: access control list (ACL) rules
	//
	// 	- **bot_manager**: bot traffic management
	//
	// 	- **ip_reputation**: IP reputation library
	//
	// example:
	//
	// {\\"dcdn_allow\\":[\\"cc_rule\\",\\"ddos_dispatch\\"]}
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// domain_business_control
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeCdnUserConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) SetArgName(v string) *DescribeCdnUserConfigsResponseBodyConfigs {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) SetArgValue(v string) *DescribeCdnUserConfigsResponseBodyConfigs {
	s.ArgValue = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) SetFunctionName(v string) *DescribeCdnUserConfigsResponseBodyConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
