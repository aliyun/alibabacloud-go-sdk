// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMetricRuleTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *PutMetricRuleTargetsRequest
	GetDryRun() *bool
	SetRegionId(v string) *PutMetricRuleTargetsRequest
	GetRegionId() *string
	SetRuleId(v string) *PutMetricRuleTargetsRequest
	GetRuleId() *string
	SetTargets(v []*PutMetricRuleTargetsRequestTargets) *PutMetricRuleTargetsRequest
	GetTargets() []*PutMetricRuleTargetsRequestTargets
}

type PutMetricRuleTargetsRequest struct {
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// false
	DryRun   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// For more information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
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

func (s *PutMetricRuleTargetsRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *PutMetricRuleTargetsRequest) SetDryRun(v bool) *PutMetricRuleTargetsRequest {
	s.DryRun = &v
	return s
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
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutMetricRuleTargetsRequestTargets struct {
	// The ARN of the resource. Simple Message Queue (formerly MNS), Auto Scaling (ESS), Simple Log Service (SLS), and Function Compute (FC) are supported.
	//
	// The ARN parameters for Simple Message Queue (formerly MNS) are as follows:
	//
	// `acs:mns:{regionId}:{userId}:/{resource type}/{resource name}/message`.
	//
	// - {regionId}: the region where the Simple Message Queue (formerly MNS) queue or topic resides.
	//
	// - {userId}: the Alibaba Cloud account to which the resource belongs.
	//
	// - {resource type}: the type of the resource that accepts alerts. Valid values:
	//
	//   - **queues**: queue
	//
	//   - **topics**: topic
	//
	// - {resource name}: the name of the resource.
	//
	//   - If the resource type is **queues**, the resource name is the queue name.
	//
	//   - If the resource type is **topics**, the resource name is the topic name.
	//
	// The ARN for Auto Scaling:
	//
	// acs:ess:{regionId}:{userId}:scalingGroupId/{scaling group ID}:scalingRuleId/{scaling rule ID}
	//
	// The ARN for Simple Log Service:
	//
	// acs:log:{regionId}:{userId}:project/{project name}/logstore/{Logstore name}
	//
	// The ARN for Function Compute:
	//
	// acs:fc:{regionId}:{userId}:services/{service name}/functions/{function name}
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the alert trigger target.
	//
	// For more information about how to obtain the ID of an alert trigger target, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-formatted parameters for the alert callback.
	//
	// example:
	//
	// {"customField1":"value1","customField2":"$.name"}
	JsonParams *string `json:"JsonParams,omitempty" xml:"JsonParams,omitempty"`
	// The alert level. Valid values:
	//
	// - INFO: information
	//
	// - WARN: warning
	//
	// - CRITICAL: critical
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
