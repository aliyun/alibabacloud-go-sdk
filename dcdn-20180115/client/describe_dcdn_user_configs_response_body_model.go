// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*DescribeDcdnUserConfigsResponseBodyConfigs) *DescribeDcdnUserConfigsResponseBody
	GetConfigs() []*DescribeDcdnUserConfigsResponseBodyConfigs
	SetRequestId(v string) *DescribeDcdnUserConfigsResponseBody
	GetRequestId() *string
}

type DescribeDcdnUserConfigsResponseBody struct {
	// The user configurations.
	Configs []*DescribeDcdnUserConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 06D29681-B7CD-4034-A8CC-28AFFA213539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserConfigsResponseBody) GetConfigs() []*DescribeDcdnUserConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *DescribeDcdnUserConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserConfigsResponseBody) SetConfigs(v []*DescribeDcdnUserConfigsResponseBodyConfigs) *DescribeDcdnUserConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *DescribeDcdnUserConfigsResponseBody) SetRequestId(v string) *DescribeDcdnUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserConfigsResponseBodyConfigs struct {
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
	// 	- cc_rule: HTTP flood protection
	//
	// 	- ddos_dispatch: DDoS mitigation
	//
	// 	- edge_safe: application security on points of presence (POPs)
	//
	// 	- blocked_regions: region blacklist
	//
	// 	- http_acl_policy: precise access control
	//
	// 	- bot_manager: bot traffic management
	//
	// 	- ip_reputation: IP reputation library
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

func (s DescribeDcdnUserConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) SetArgName(v string) *DescribeDcdnUserConfigsResponseBodyConfigs {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) SetArgValue(v string) *DescribeDcdnUserConfigsResponseBodyConfigs {
	s.ArgValue = &v
	return s
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) SetFunctionName(v string) *DescribeDcdnUserConfigsResponseBodyConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnUserConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
