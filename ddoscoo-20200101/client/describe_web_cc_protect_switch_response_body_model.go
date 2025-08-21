// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCcProtectSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProtectSwitchList(v []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) *DescribeWebCcProtectSwitchResponseBody
	GetProtectSwitchList() []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList
	SetRequestId(v string) *DescribeWebCcProtectSwitchResponseBody
	GetRequestId() *string
}

type DescribeWebCcProtectSwitchResponseBody struct {
	// The status of each mitigation policy for the website.
	ProtectSwitchList []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList `json:"ProtectSwitchList,omitempty" xml:"ProtectSwitchList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3ADD9EED-CA4B-488C-BC82-01B0B899363D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebCcProtectSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCcProtectSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchResponseBody) GetProtectSwitchList() []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	return s.ProtectSwitchList
}

func (s *DescribeWebCcProtectSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebCcProtectSwitchResponseBody) SetProtectSwitchList(v []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) *DescribeWebCcProtectSwitchResponseBody {
	s.ProtectSwitchList = v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBody) SetRequestId(v string) *DescribeWebCcProtectSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebCcProtectSwitchResponseBodyProtectSwitchList struct {
	// The mode of Intelligent Protection. Valid values:
	//
	// 	- **watch**: Warning
	//
	// 	- **defense**: Defense
	//
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// The status of Intelligent Protection. Valid values:
	//
	// 	- **0**: turned off
	//
	// 	- **1:*	- turned on
	//
	// example:
	//
	// 1
	AiRuleEnable *int32 `json:"AiRuleEnable,omitempty" xml:"AiRuleEnable,omitempty"`
	// The level of Intelligent Protection. Valid values:
	//
	// 	- **level30**: Loose
	//
	// 	- **level60**: Normal
	//
	// 	- **level90**: Strict
	//
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// The status of Blacklist/Whitelist (Domain Names). Valid values:
	//
	// 	- **0**: turned off
	//
	// 	- **1:*	- turned on
	//
	// example:
	//
	// 1
	BlackWhiteListEnable *int32 `json:"BlackWhiteListEnable,omitempty" xml:"BlackWhiteListEnable,omitempty"`
	// The status of the Custom Rules switch for Frequency Control. Valid values:
	//
	// 	- **0**: turned off
	//
	// 	- **1:*	- turned on
	//
	// example:
	//
	// 0
	CcCustomRuleEnable *int32 `json:"CcCustomRuleEnable,omitempty" xml:"CcCustomRuleEnable,omitempty"`
	// The status of Frequency Control. Valid values:
	//
	// 	- **0**: turned off
	//
	// 	- **1:*	- turned on
	//
	// example:
	//
	// 1
	CcEnable       *int32  `json:"CcEnable,omitempty" xml:"CcEnable,omitempty"`
	CcGlobalSwitch *string `json:"CcGlobalSwitch,omitempty" xml:"CcGlobalSwitch,omitempty"`
	// The mode of Frequency Control. Valid values:
	//
	// 	- **default**: Normal
	//
	// 	- **gf_under_attack**: Emergency
	//
	// 	- **gf_sos_verify**: Strict
	//
	// 	- **gf_sos_enhance**: Super Strict
	//
	// example:
	//
	// default
	CcTemplate *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The status of Accurate Access Control. Valid values:
	//
	// 	- **0**: turned off
	//
	// 	- **1:*	- turned on
	//
	// example:
	//
	// 0
	PreciseRuleEnable *int32 `json:"PreciseRuleEnable,omitempty" xml:"PreciseRuleEnable,omitempty"`
	// The status of Location Blacklist (Domain Names). Valid values:
	//
	// 	- **0**: turned off
	//
	// 	- **1:*	- turned on
	//
	// example:
	//
	// 0
	RegionBlockEnable *int32 `json:"RegionBlockEnable,omitempty" xml:"RegionBlockEnable,omitempty"`
}

func (s DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetAiMode() *string {
	return s.AiMode
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetAiRuleEnable() *int32 {
	return s.AiRuleEnable
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetAiTemplate() *string {
	return s.AiTemplate
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetBlackWhiteListEnable() *int32 {
	return s.BlackWhiteListEnable
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetCcCustomRuleEnable() *int32 {
	return s.CcCustomRuleEnable
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetCcEnable() *int32 {
	return s.CcEnable
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetCcGlobalSwitch() *string {
	return s.CcGlobalSwitch
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetCcTemplate() *string {
	return s.CcTemplate
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetPreciseRuleEnable() *int32 {
	return s.PreciseRuleEnable
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GetRegionBlockEnable() *int32 {
	return s.RegionBlockEnable
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetAiMode(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiMode = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetAiRuleEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiRuleEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetAiTemplate(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiTemplate = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetBlackWhiteListEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.BlackWhiteListEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcCustomRuleEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcCustomRuleEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcGlobalSwitch(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcGlobalSwitch = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcTemplate(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcTemplate = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetDomain(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.Domain = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetPreciseRuleEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.PreciseRuleEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetRegionBlockEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.RegionBlockEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) Validate() error {
	return dara.Validate(s)
}
