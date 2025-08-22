// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDdosSpecInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthLimit(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetBandwidthLimit() *string
	SetEdition(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetEdition() *string
	SetEnable(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetEnable() *string
	SetIsSpecialPort(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetIsSpecialPort() *string
	SetProtectedArea(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetProtectedArea() *string
	SetQpsLimit(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetQpsLimit() *string
	SetRequestId(v string) *DescribeDcdnDdosSpecInfoResponseBody
	GetRequestId() *string
	SetSpecInfos(v []*DescribeDcdnDdosSpecInfoResponseBodySpecInfos) *DescribeDcdnDdosSpecInfoResponseBody
	GetSpecInfos() []*DescribeDcdnDdosSpecInfoResponseBodySpecInfos
}

type DescribeDcdnDdosSpecInfoResponseBody struct {
	// The bandwidth limit of a single instance.
	//
	// example:
	//
	// 40Gbps
	BandwidthLimit *string `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The version. Valid values:
	//
	// 	- **poc**: POC Edition
	//
	// 	- **basic**: Basic Edition
	//
	// 	- **insurance**: Insurance Edition
	//
	// 	- **unlimited**: Unlimited Edition
	//
	// 	- **port_enhancement**: Special Port Enhanced Edition
	//
	// example:
	//
	// insurance
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Specifies whether to enable DDoS mitigation. Valid values:
	//
	// 	- **on:**
	//
	// 	- **off**.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Specifies whether custom ports are supported. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	IsSpecialPort *string `json:"IsSpecialPort,omitempty" xml:"IsSpecialPort,omitempty"`
	// Protected region. Valid values:
	//
	// 	- **global**: global
	//
	// 	- **chinese_mainland**: Chinese mainland
	//
	// 	- **global_excluding_the_chinese_mainland**: outside the Chinese mainland
	//
	// example:
	//
	// global
	ProtectedArea *string `json:"ProtectedArea,omitempty" xml:"ProtectedArea,omitempty"`
	// The QPS limit.
	//
	// example:
	//
	// 100
	QpsLimit *string `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DEA8E7BE-33C6-56F5-AC56-74D50547CF34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code and configurations of the security rules.
	SpecInfos []*DescribeDcdnDdosSpecInfoResponseBodySpecInfos `json:"SpecInfos,omitempty" xml:"SpecInfos,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDdosSpecInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosSpecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetBandwidthLimit() *string {
	return s.BandwidthLimit
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetEdition() *string {
	return s.Edition
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetIsSpecialPort() *string {
	return s.IsSpecialPort
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetProtectedArea() *string {
	return s.ProtectedArea
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetQpsLimit() *string {
	return s.QpsLimit
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) GetSpecInfos() []*DescribeDcdnDdosSpecInfoResponseBodySpecInfos {
	return s.SpecInfos
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetBandwidthLimit(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetEdition(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetEnable(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.Enable = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetIsSpecialPort(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.IsSpecialPort = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetProtectedArea(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.ProtectedArea = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetQpsLimit(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.QpsLimit = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetRequestId(v string) *DescribeDcdnDdosSpecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) SetSpecInfos(v []*DescribeDcdnDdosSpecInfoResponseBodySpecInfos) *DescribeDcdnDdosSpecInfoResponseBody {
	s.SpecInfos = v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDdosSpecInfoResponseBodySpecInfos struct {
	// The configurations of the version rule.
	Configs []*DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The version rule. Valid values:
	//
	// 	- **version_defense_num**: the rule for the number of version mitigation sessions
	//
	// 	- **domain_num**: the rule for the limit on the number of domain names
	//
	// 	- **defence_package_num**: the rule for extra mitigation session plans
	//
	// example:
	//
	// version_defense_num
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s DescribeDcdnDdosSpecInfoResponseBodySpecInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosSpecInfoResponseBodySpecInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfos) GetConfigs() []*DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs {
	return s.Configs
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfos) GetRule() *string {
	return s.Rule
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfos) SetConfigs(v []*DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) *DescribeDcdnDdosSpecInfoResponseBodySpecInfos {
	s.Configs = v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfos) SetRule(v string) *DescribeDcdnDdosSpecInfoResponseBodySpecInfos {
	s.Rule = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs struct {
	// The configuration code of the version rule. Valid values:
	//
	// 	- **total_defense_num**: the total number of mitigation sessions of the version.
	//
	// 	- **consume_defense_num**: the number of used mitigation sessions of the version.
	//
	// 	- **max_domain_num**: the limit on the number of added domain names.
	//
	// 	- **emain_domain_num**: the number of added domain names.
	//
	// 	- **defence_package_num**: the total number of purchased additional mitigation sessions.
	//
	// 	- **consume_defence_package_num**: the number of used additional mitigation sessions.
	//
	// example:
	//
	// total_defense_num
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration expression of the version rule.
	//
	// example:
	//
	// equal
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The value of the configuration expression of the version rule.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) GetConfig() *string {
	return s.Config
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) GetExpr() *string {
	return s.Expr
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) SetConfig(v string) *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs {
	s.Config = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) SetExpr(v string) *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs {
	s.Expr = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) SetValue(v string) *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponseBodySpecInfosConfigs) Validate() error {
	return dara.Validate(s)
}
