// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyElasticRuleResponseBodyData) *ModifyElasticRuleResponseBody
	GetData() *ModifyElasticRuleResponseBodyData
	SetRequestId(v string) *ModifyElasticRuleResponseBody
	GetRequestId() *string
}

type ModifyElasticRuleResponseBody struct {
	// The data returned.
	Data *ModifyElasticRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5ED62C81-9948-5612-81E1-EA3853752306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticRuleResponseBody) GetData() *ModifyElasticRuleResponseBodyData {
	return s.Data
}

func (s *ModifyElasticRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticRuleResponseBody) SetData(v *ModifyElasticRuleResponseBodyData) *ModifyElasticRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyElasticRuleResponseBody) SetRequestId(v string) *ModifyElasticRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyElasticRuleResponseBodyData struct {
	// The rule for computing resources of the required cluster.
	//
	// example:
	//
	// selectdb.2xlarge
	ClusterClass *string `json:"ClusterClass,omitempty" xml:"ClusterClass,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-zpr3if5wq03-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The time when the scheduled scaling rule is executed.
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
	// 29252
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyElasticRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyElasticRuleResponseBodyData) GetClusterClass() *string {
	return s.ClusterClass
}

func (s *ModifyElasticRuleResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyElasticRuleResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyElasticRuleResponseBodyData) GetElasticRuleStartTime() *string {
	return s.ElasticRuleStartTime
}

func (s *ModifyElasticRuleResponseBodyData) GetExecutionPeriod() *string {
	return s.ExecutionPeriod
}

func (s *ModifyElasticRuleResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyElasticRuleResponseBodyData) SetClusterClass(v string) *ModifyElasticRuleResponseBodyData {
	s.ClusterClass = &v
	return s
}

func (s *ModifyElasticRuleResponseBodyData) SetClusterId(v string) *ModifyElasticRuleResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ModifyElasticRuleResponseBodyData) SetDbInstanceId(v string) *ModifyElasticRuleResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyElasticRuleResponseBodyData) SetElasticRuleStartTime(v string) *ModifyElasticRuleResponseBodyData {
	s.ElasticRuleStartTime = &v
	return s
}

func (s *ModifyElasticRuleResponseBodyData) SetExecutionPeriod(v string) *ModifyElasticRuleResponseBodyData {
	s.ExecutionPeriod = &v
	return s
}

func (s *ModifyElasticRuleResponseBodyData) SetRuleId(v int64) *ModifyElasticRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifyElasticRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
