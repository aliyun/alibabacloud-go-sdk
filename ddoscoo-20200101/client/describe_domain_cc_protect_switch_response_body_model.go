// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCcProtectSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProtectSwitchList(v []*DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) *DescribeDomainCcProtectSwitchResponseBody
	GetProtectSwitchList() []*DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList
	SetRequestId(v string) *DescribeDomainCcProtectSwitchResponseBody
	GetRequestId() *string
}

type DescribeDomainCcProtectSwitchResponseBody struct {
	// The status of each protection feature for the website.
	ProtectSwitchList []*DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList `json:"ProtectSwitchList,omitempty" xml:"ProtectSwitchList,omitempty" type:"Repeated"`
	// The ID of the request. This ID is unique to the request. Use the request ID to troubleshoot issues.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainCcProtectSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcProtectSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcProtectSwitchResponseBody) GetProtectSwitchList() []*DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	return s.ProtectSwitchList
}

func (s *DescribeDomainCcProtectSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCcProtectSwitchResponseBody) SetProtectSwitchList(v []*DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) *DescribeDomainCcProtectSwitchResponseBody {
	s.ProtectSwitchList = v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBody) SetRequestId(v string) *DescribeDomainCcProtectSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBody) Validate() error {
	if s.ProtectSwitchList != nil {
		for _, item := range s.ProtectSwitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList struct {
	// The mode of AI-powered protection. Valid values:
	//
	// - **watch**: Warning mode
	//
	// - **defense**: Defense mode
	//
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// The status of the AI-powered protection feature. Valid values:
	//
	// - **0**: Disabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 1
	AiRuleEnable *int32 `json:"AiRuleEnable,omitempty" xml:"AiRuleEnable,omitempty"`
	// The level of AI-powered protection. Valid values:
	//
	// - **level30**: Loose
	//
	// - **level60**: Normal
	//
	// - **level90**: Strict
	//
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// The status of the blacklists and whitelists feature. Valid values:
	//
	// - **0**: Disabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 1
	BlackWhiteListEnable *int32 `json:"BlackWhiteListEnable,omitempty" xml:"BlackWhiteListEnable,omitempty"`
	// The status of custom CC rules. Valid values:
	//
	// - **0**: Disabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 0
	CcCustomRuleEnable *int32 `json:"CcCustomRuleEnable,omitempty" xml:"CcCustomRuleEnable,omitempty"`
	// The status of the rate limiting (CC protection) feature. Valid values:
	//
	// - **0**: Disabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 1
	CcEnable *int32 `json:"CcEnable,omitempty" xml:"CcEnable,omitempty"`
	// The status of the global switch for custom and precise CC protection. Valid values:
	//
	// - **open**: Enabled
	//
	// - **close**: Disabled
	//
	// example:
	//
	// close
	CcGlobalSwitch *string `json:"CcGlobalSwitch,omitempty" xml:"CcGlobalSwitch,omitempty"`
	// The CC protection mode. Valid values:
	//
	// - **default**: Normal
	//
	// - **gf_under_attack**: Emergency
	//
	// - **gf_sos_verify**: Strict
	//
	// - **gf_sos_enhance**: Very Strict
	//
	// example:
	//
	// default
	CcTemplate *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The status of the precise access control feature. Valid values:
	//
	// - **0**: Disabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 0
	PreciseRuleEnable *int32 `json:"PreciseRuleEnable,omitempty" xml:"PreciseRuleEnable,omitempty"`
	// The status of the Location Blacklist (Domain Names) feature. Valid values:
	//
	// - **0**: Disabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 0
	RegionBlockEnable *int32 `json:"RegionBlockEnable,omitempty" xml:"RegionBlockEnable,omitempty"`
}

func (s DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetAiMode() *string {
	return s.AiMode
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetAiRuleEnable() *int32 {
	return s.AiRuleEnable
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetAiTemplate() *string {
	return s.AiTemplate
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetBlackWhiteListEnable() *int32 {
	return s.BlackWhiteListEnable
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetCcCustomRuleEnable() *int32 {
	return s.CcCustomRuleEnable
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetCcEnable() *int32 {
	return s.CcEnable
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetCcGlobalSwitch() *string {
	return s.CcGlobalSwitch
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetCcTemplate() *string {
	return s.CcTemplate
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetPreciseRuleEnable() *int32 {
	return s.PreciseRuleEnable
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) GetRegionBlockEnable() *int32 {
	return s.RegionBlockEnable
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetAiMode(v string) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiMode = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetAiRuleEnable(v int32) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiRuleEnable = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetAiTemplate(v string) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiTemplate = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetBlackWhiteListEnable(v int32) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.BlackWhiteListEnable = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetCcCustomRuleEnable(v int32) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcCustomRuleEnable = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetCcEnable(v int32) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcEnable = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetCcGlobalSwitch(v string) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcGlobalSwitch = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetCcTemplate(v string) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetDomain(v string) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetPreciseRuleEnable(v int32) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.PreciseRuleEnable = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) SetRegionBlockEnable(v int32) *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList {
	s.RegionBlockEnable = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponseBodyProtectSwitchList) Validate() error {
	return dara.Validate(s)
}
