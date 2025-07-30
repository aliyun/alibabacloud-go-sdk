// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateElasticRuleResponseBodyData) *CreateElasticRuleResponseBody
	GetData() *CreateElasticRuleResponseBodyData
	SetRequestId(v string) *CreateElasticRuleResponseBody
	GetRequestId() *string
}

type CreateElasticRuleResponseBody struct {
	// The data returned.
	Data *CreateElasticRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticRuleResponseBody) GetData() *CreateElasticRuleResponseBodyData {
	return s.Data
}

func (s *CreateElasticRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateElasticRuleResponseBody) SetData(v *CreateElasticRuleResponseBodyData) *CreateElasticRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateElasticRuleResponseBody) SetRequestId(v string) *CreateElasticRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateElasticRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateElasticRuleResponseBodyData struct {
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
	// selectdb-xxxb9f2w-be
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
	// 5465
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateElasticRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateElasticRuleResponseBodyData) GetClusterClass() *string {
	return s.ClusterClass
}

func (s *CreateElasticRuleResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateElasticRuleResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateElasticRuleResponseBodyData) GetElasticRuleStartTime() *string {
	return s.ElasticRuleStartTime
}

func (s *CreateElasticRuleResponseBodyData) GetExecutionPeriod() *string {
	return s.ExecutionPeriod
}

func (s *CreateElasticRuleResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateElasticRuleResponseBodyData) SetClusterClass(v string) *CreateElasticRuleResponseBodyData {
	s.ClusterClass = &v
	return s
}

func (s *CreateElasticRuleResponseBodyData) SetClusterId(v string) *CreateElasticRuleResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *CreateElasticRuleResponseBodyData) SetDbInstanceId(v string) *CreateElasticRuleResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *CreateElasticRuleResponseBodyData) SetElasticRuleStartTime(v string) *CreateElasticRuleResponseBodyData {
	s.ElasticRuleStartTime = &v
	return s
}

func (s *CreateElasticRuleResponseBodyData) SetExecutionPeriod(v string) *CreateElasticRuleResponseBodyData {
	s.ExecutionPeriod = &v
	return s
}

func (s *CreateElasticRuleResponseBodyData) SetRuleId(v int64) *CreateElasticRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateElasticRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
