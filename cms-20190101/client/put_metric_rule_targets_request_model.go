// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMetricRuleTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *PutMetricRuleTargetsRequest
	GetRegionId() *string
	SetRuleId(v string) *PutMetricRuleTargetsRequest
	GetRuleId() *string
	SetTargets(v []*PutMetricRuleTargetsRequestTargets) *PutMetricRuleTargetsRequest
	GetTargets() []*PutMetricRuleTargetsRequestTargets
}

type PutMetricRuleTargetsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ae06917_75a8c43178ab66****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// None.
	//
	// This parameter is required.
	Targets []*PutMetricRuleTargetsRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s PutMetricRuleTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutMetricRuleTargetsRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutMetricRuleTargetsRequest) GetTargets() []*PutMetricRuleTargetsRequestTargets {
	return s.Targets
}

func (s *PutMetricRuleTargetsRequest) SetRegionId(v string) *PutMetricRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *PutMetricRuleTargetsRequest) SetRuleId(v string) *PutMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

func (s *PutMetricRuleTargetsRequest) SetTargets(v []*PutMetricRuleTargetsRequestTargets) *PutMetricRuleTargetsRequest {
	s.Targets = v
	return s
}

func (s *PutMetricRuleTargetsRequest) Validate() error {
	return dara.Validate(s)
}

type PutMetricRuleTargetsRequestTargets struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource. Simple Message Queue (formerly MNS) (SMQ), Auto Scaling, Simple Log Service, and Function Compute are supported.
	//
	// The following part describes the ARN of SMQ and the parameters in the ARN:
	//
	// `acs:mns:{regionId}:{userId}:/{Resource type}/{Resource name}/message`.
	//
	// 	- {regionId}: the region ID of the SMQ queue or topic.
	//
	// 	- {userId}: the ID of the Alibaba Cloud account that owns the resource.
	//
	// 	- {Resource type}: the type of the resource for which alerts are triggered. Valid values:
	//
	//     	- **queues**
	//
	//     	- **topics**
	//
	// 	- {Resource name}: the resource name.
	//
	//     	- If the resource type is **queues**, the resource name is the queue name.
	//
	//     	- If the resource type is **topics**, the resource name is the topic name.
	//
	// ARN of Auto Scaling:
	//
	// acs:ess:{regionId}:{userId}:scalingGroupId/{Scaling group ID}:scalingRuleId/{Scaling rule ID}
	//
	// ARN of Simple Log Service:
	//
	// acs:log:{regionId}:{userId}:project/{Project name}/logstore/{Logstore name}
	//
	// ARN of Function Compute:
	//
	// acs:fc:{regionId}:{userId}:services/{Service name}/functions/{Function name}
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the resource for which alerts are triggered.
	//
	// For more information about how to obtain the ID of the resource for which alerts are triggered, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters of the alert callback. The parameters are in the JSON format.
	//
	// example:
	//
	// {"customField1":"value1","customField2":"$.name"}
	JsonParams *string `json:"JsonParams,omitempty" xml:"JsonParams,omitempty"`
	// The alert level. Valid values:
	//
	// 	- INFO
	//
	// 	- WARN
	//
	// 	- CRITICAL
	//
	// example:
	//
	// ["INFO", "WARN", "CRITICAL"]
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutMetricRuleTargetsRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsRequestTargets) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsRequestTargets) GetArn() *string {
	return s.Arn
}

func (s *PutMetricRuleTargetsRequestTargets) GetId() *string {
	return s.Id
}

func (s *PutMetricRuleTargetsRequestTargets) GetJsonParams() *string {
	return s.JsonParams
}

func (s *PutMetricRuleTargetsRequestTargets) GetLevel() *string {
	return s.Level
}

func (s *PutMetricRuleTargetsRequestTargets) SetArn(v string) *PutMetricRuleTargetsRequestTargets {
	s.Arn = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetId(v string) *PutMetricRuleTargetsRequestTargets {
	s.Id = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetJsonParams(v string) *PutMetricRuleTargetsRequestTargets {
	s.JsonParams = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetLevel(v string) *PutMetricRuleTargetsRequestTargets {
	s.Level = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) Validate() error {
	return dara.Validate(s)
}
