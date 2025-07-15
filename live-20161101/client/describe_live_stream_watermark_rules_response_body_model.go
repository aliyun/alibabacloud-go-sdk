// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarkRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveStreamWatermarkRulesResponseBody
	GetRequestId() *string
	SetRuleInfoList(v *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) *DescribeLiveStreamWatermarkRulesResponseBody
	GetRuleInfoList() *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList
	SetTotal(v int32) *DescribeLiveStreamWatermarkRulesResponseBody
	GetTotal() *int32
}

type DescribeLiveStreamWatermarkRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The watermark rules.
	RuleInfoList *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList `json:"RuleInfoList,omitempty" xml:"RuleInfoList,omitempty" type:"Struct"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveStreamWatermarkRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) GetRuleInfoList() *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList {
	return s.RuleInfoList
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) SetRequestId(v string) *DescribeLiveStreamWatermarkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) SetRuleInfoList(v *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) *DescribeLiveStreamWatermarkRulesResponseBody {
	s.RuleInfoList = v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) SetTotal(v int32) *DescribeLiveStreamWatermarkRulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList struct {
	RuleInfo []*DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) GetRuleInfo() []*DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	return s.RuleInfo
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) SetRuleInfo(v []*DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList {
	s.RuleInfo = v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The description of the custom rule.
	//
	// example:
	//
	// my rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the custom rule.
	//
	// example:
	//
	// WatermarkRule****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the watermark rule.
	//
	// >  You can obtain the rule ID by checking the value of the RuleId parameter that is returned by the [AddLiveStreamWatermarkRule](https://help.aliyun.com/document_detail/2848100.html) operation.
	//
	// example:
	//
	// 445409ec-7eaa-461d -8f29-4bec2eb9****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the live stream. The following matching rules apply:
	//
	// 	- A stream name can be exactly matched. Example: liveStreamA.
	//
	// 	- Fuzzy match is also supported. The use of an asterisk (`*`) allows all approximate matches to be found.
	//
	// 	- You can place the asterisk before or after an approximate string.
	//
	// >
	//
	// 	- Fuzzy match: Only one asterisk (`*`) before or after an approximate string is allowed. The approximate string must be enclosed in `()`. Separate multiple strings with vertical bars (`|`).
	//
	// 	- For example, `*(t1|t2)` matches all streams whose name has the `t1` or `t2` suffix, and `(abc|123)*` matches all streams whose name has the `abc` or `123` prefix.
	//
	// example:
	//
	// liveStreamA
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The ID of the watermark template.
	//
	// >  You can obtain the template ID by checking the value of the TemplateId parameter that is returned by the [AddLiveStreamWatermark](https://help.aliyun.com/document_detail/2848096.html) operation.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9 ****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetApp() *string {
	return s.App
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetName() *string {
	return s.Name
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetApp(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.App = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetDescription(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Description = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetDomain(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Domain = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetName(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Name = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetRuleId(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.RuleId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetStream(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Stream = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetTemplateId(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.TemplateId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) Validate() error {
	return dara.Validate(s)
}
