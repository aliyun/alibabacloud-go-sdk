// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeElasticRulesResponseBodyData) *DescribeElasticRulesResponseBody
	GetData() *DescribeElasticRulesResponseBodyData
	SetRequestId(v string) *DescribeElasticRulesResponseBody
	GetRequestId() *string
}

type DescribeElasticRulesResponseBody struct {
	// The data returned.
	Data *DescribeElasticRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticRulesResponseBody) GetData() *DescribeElasticRulesResponseBodyData {
	return s.Data
}

func (s *DescribeElasticRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticRulesResponseBody) SetData(v *DescribeElasticRulesResponseBodyData) *DescribeElasticRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeElasticRulesResponseBody) SetRequestId(v string) *DescribeElasticRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticRulesResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-nwy3jv1oa02-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The details of the rules.
	Rules []*DescribeElasticRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeElasticRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeElasticRulesResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeElasticRulesResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeElasticRulesResponseBodyData) GetRules() []*DescribeElasticRulesResponseBodyDataRules {
	return s.Rules
}

func (s *DescribeElasticRulesResponseBodyData) SetClusterId(v string) *DescribeElasticRulesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DescribeElasticRulesResponseBodyData) SetDbInstanceId(v string) *DescribeElasticRulesResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeElasticRulesResponseBodyData) SetRules(v []*DescribeElasticRulesResponseBodyDataRules) *DescribeElasticRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeElasticRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticRulesResponseBodyDataRules struct {
	// The rule for computing resources of the required cluster.
	//
	// example:
	//
	// selectdb.2xlarge
	ClusterClass *string `json:"ClusterClass,omitempty" xml:"ClusterClass,omitempty"`
	// The time when you want to execute the scheduled scaling rule.
	//
	// example:
	//
	// 00:00
	ElasticRuleStartTime *string `json:"ElasticRuleStartTime,omitempty" xml:"ElasticRuleStartTime,omitempty"`
	// The execution cycle.
	//
	// Valid value:
	//
	// 	- Day
	//
	// example:
	//
	// Day
	ExecutionPeriod *string `json:"ExecutionPeriod,omitempty" xml:"ExecutionPeriod,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 5467
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeElasticRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeElasticRulesResponseBodyDataRules) GetClusterClass() *string {
	return s.ClusterClass
}

func (s *DescribeElasticRulesResponseBodyDataRules) GetElasticRuleStartTime() *string {
	return s.ElasticRuleStartTime
}

func (s *DescribeElasticRulesResponseBodyDataRules) GetExecutionPeriod() *string {
	return s.ExecutionPeriod
}

func (s *DescribeElasticRulesResponseBodyDataRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeElasticRulesResponseBodyDataRules) SetClusterClass(v string) *DescribeElasticRulesResponseBodyDataRules {
	s.ClusterClass = &v
	return s
}

func (s *DescribeElasticRulesResponseBodyDataRules) SetElasticRuleStartTime(v string) *DescribeElasticRulesResponseBodyDataRules {
	s.ElasticRuleStartTime = &v
	return s
}

func (s *DescribeElasticRulesResponseBodyDataRules) SetExecutionPeriod(v string) *DescribeElasticRulesResponseBodyDataRules {
	s.ExecutionPeriod = &v
	return s
}

func (s *DescribeElasticRulesResponseBodyDataRules) SetRuleId(v int64) *DescribeElasticRulesResponseBodyDataRules {
	s.RuleId = &v
	return s
}

func (s *DescribeElasticRulesResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}
