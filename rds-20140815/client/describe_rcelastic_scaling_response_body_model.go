// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCElasticScalingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcuType(v string) *DescribeRCElasticScalingResponseBody
	GetAcuType() *string
	SetInstanceId(v string) *DescribeRCElasticScalingResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeRCElasticScalingResponseBody
	GetRequestId() *string
	SetScalingEnabled(v bool) *DescribeRCElasticScalingResponseBody
	GetScalingEnabled() *bool
	SetScalingSupported(v bool) *DescribeRCElasticScalingResponseBody
	GetScalingSupported() *bool
	SetScheduledRule(v string) *DescribeRCElasticScalingResponseBody
	GetScheduledRule() *string
	SetScheduledRuleTemplates(v *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) *DescribeRCElasticScalingResponseBody
	GetScheduledRuleTemplates() *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates
	SetTargetInstanceType(v string) *DescribeRCElasticScalingResponseBody
	GetTargetInstanceType() *string
}

type DescribeRCElasticScalingResponseBody struct {
	// example:
	//
	// gn8is
	AcuType *string `json:"AcuType,omitempty" xml:"AcuType,omitempty"`
	// example:
	//
	// rc-a0e466b01tif2pk
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1F7B8B09-36F3-1165-BADB-13E376FE9BD7
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingEnabled   *bool   `json:"ScalingEnabled,omitempty" xml:"ScalingEnabled,omitempty"`
	ScalingSupported *bool   `json:"ScalingSupported,omitempty" xml:"ScalingSupported,omitempty"`
	// example:
	//
	// {"rule":[{"beginTime":"09:00","endTime":"17:00","acu":4}]}
	ScheduledRule          *string                                                     `json:"ScheduledRule,omitempty" xml:"ScheduledRule,omitempty"`
	ScheduledRuleTemplates *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates `json:"ScheduledRuleTemplates,omitempty" xml:"ScheduledRuleTemplates,omitempty" type:"Struct"`
	// example:
	//
	// mysql.x2.medium.9cm
	TargetInstanceType *string `json:"TargetInstanceType,omitempty" xml:"TargetInstanceType,omitempty"`
}

func (s DescribeRCElasticScalingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCElasticScalingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCElasticScalingResponseBody) GetAcuType() *string {
	return s.AcuType
}

func (s *DescribeRCElasticScalingResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCElasticScalingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCElasticScalingResponseBody) GetScalingEnabled() *bool {
	return s.ScalingEnabled
}

func (s *DescribeRCElasticScalingResponseBody) GetScalingSupported() *bool {
	return s.ScalingSupported
}

func (s *DescribeRCElasticScalingResponseBody) GetScheduledRule() *string {
	return s.ScheduledRule
}

func (s *DescribeRCElasticScalingResponseBody) GetScheduledRuleTemplates() *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates {
	return s.ScheduledRuleTemplates
}

func (s *DescribeRCElasticScalingResponseBody) GetTargetInstanceType() *string {
	return s.TargetInstanceType
}

func (s *DescribeRCElasticScalingResponseBody) SetAcuType(v string) *DescribeRCElasticScalingResponseBody {
	s.AcuType = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetInstanceId(v string) *DescribeRCElasticScalingResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetRequestId(v string) *DescribeRCElasticScalingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetScalingEnabled(v bool) *DescribeRCElasticScalingResponseBody {
	s.ScalingEnabled = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetScalingSupported(v bool) *DescribeRCElasticScalingResponseBody {
	s.ScalingSupported = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetScheduledRule(v string) *DescribeRCElasticScalingResponseBody {
	s.ScheduledRule = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetScheduledRuleTemplates(v *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) *DescribeRCElasticScalingResponseBody {
	s.ScheduledRuleTemplates = v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) SetTargetInstanceType(v string) *DescribeRCElasticScalingResponseBody {
	s.TargetInstanceType = &v
	return s
}

func (s *DescribeRCElasticScalingResponseBody) Validate() error {
	if s.ScheduledRuleTemplates != nil {
		if err := s.ScheduledRuleTemplates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCElasticScalingResponseBodyScheduledRuleTemplates struct {
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) GoString() string {
	return s.String()
}

func (s *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) GetItems() []*string {
	return s.Items
}

func (s *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) SetItems(v []*string) *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates {
	s.Items = v
	return s
}

func (s *DescribeRCElasticScalingResponseBodyScheduledRuleTemplates) Validate() error {
	return dara.Validate(s)
}
