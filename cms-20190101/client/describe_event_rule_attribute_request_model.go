// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeEventRuleAttributeRequest
	GetRegionId() *string
	SetRuleName(v string) *DescribeEventRuleAttributeRequest
	GetRuleName() *string
	SetSilenceTime(v string) *DescribeEventRuleAttributeRequest
	GetSilenceTime() *string
}

type DescribeEventRuleAttributeRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the event-triggered alert rule.
	//
	// For information about how to obtain the name of an event-triggered alert rule, see [DescribeEventRuleList](https://help.aliyun.com/document_detail/114996.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period.
	//
	// Unit: seconds. Default value: 86400 (1 day).
	//
	// > When monitoring data continuously exceeds the alert rule threshold, only one alert notification is sent within each mute period.
	//
	// example:
	//
	// 86400
	SilenceTime *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
}

func (s DescribeEventRuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventRuleAttributeRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeEventRuleAttributeRequest) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *DescribeEventRuleAttributeRequest) SetRegionId(v string) *DescribeEventRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventRuleAttributeRequest) SetRuleName(v string) *DescribeEventRuleAttributeRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeEventRuleAttributeRequest) SetSilenceTime(v string) *DescribeEventRuleAttributeRequest {
	s.SilenceTime = &v
	return s
}

func (s *DescribeEventRuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
