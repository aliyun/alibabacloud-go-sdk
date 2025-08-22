// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafSpecInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *DescribeDcdnWafSpecInfoResponseBody
	GetEdition() *string
	SetRequestId(v string) *DescribeDcdnWafSpecInfoResponseBody
	GetRequestId() *string
	SetSpecInfos(v []*DescribeDcdnWafSpecInfoResponseBodySpecInfos) *DescribeDcdnWafSpecInfoResponseBody
	GetSpecInfos() []*DescribeDcdnWafSpecInfoResponseBodySpecInfos
}

type DescribeDcdnWafSpecInfoResponseBody struct {
	// The version of WAF.
	//
	// example:
	//
	// dcdnwaf_afterpay
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7281593a-f414-42c1-b7ba-2ce65e21cc00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported types of protection policies and the configuration information of protection rules.
	SpecInfos []*DescribeDcdnWafSpecInfoResponseBodySpecInfos `json:"SpecInfos,omitempty" xml:"SpecInfos,omitempty" type:"Repeated"`
}

func (s DescribeDcdnWafSpecInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafSpecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafSpecInfoResponseBody) GetEdition() *string {
	return s.Edition
}

func (s *DescribeDcdnWafSpecInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafSpecInfoResponseBody) GetSpecInfos() []*DescribeDcdnWafSpecInfoResponseBodySpecInfos {
	return s.SpecInfos
}

func (s *DescribeDcdnWafSpecInfoResponseBody) SetEdition(v string) *DescribeDcdnWafSpecInfoResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBody) SetRequestId(v string) *DescribeDcdnWafSpecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBody) SetSpecInfos(v []*DescribeDcdnWafSpecInfoResponseBodySpecInfos) *DescribeDcdnWafSpecInfoResponseBody {
	s.SpecInfos = v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafSpecInfoResponseBodySpecInfos struct {
	// The configuration information of the protection rule.
	Configs []*DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The type of the protection policy. Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom
	//
	// 	- whitelist: whitelist
	//
	// 	- ip_blacklist: IP address blacklist
	//
	// 	- region_block: region blacklist
	//
	// 	- bot: bot management
	//
	// example:
	//
	// custom_acl
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
}

func (s DescribeDcdnWafSpecInfoResponseBodySpecInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafSpecInfoResponseBodySpecInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfos) GetConfigs() []*DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs {
	return s.Configs
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfos) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfos) SetConfigs(v []*DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) *DescribeDcdnWafSpecInfoResponseBodySpecInfos {
	s.Configs = v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfos) SetDefenseScene(v string) *DescribeDcdnWafSpecInfoResponseBodySpecInfos {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs struct {
	// The configuration code of the protection rule.
	//
	// example:
	//
	// enable
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration expression of the protection rule.
	//
	// example:
	//
	// equal
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The value of the configuration expression of the protection rule.
	//
	// example:
	//
	// on
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) GetConfig() *string {
	return s.Config
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) GetExpr() *string {
	return s.Expr
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) SetConfig(v string) *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs {
	s.Config = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) SetExpr(v string) *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs {
	s.Expr = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) SetValue(v string) *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs {
	s.Value = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponseBodySpecInfosConfigs) Validate() error {
	return dara.Validate(s)
}
