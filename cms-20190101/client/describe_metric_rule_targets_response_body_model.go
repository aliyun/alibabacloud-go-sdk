// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricRuleTargetsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMetricRuleTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricRuleTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricRuleTargetsResponseBody
	GetSuccess() *bool
	SetTargets(v *DescribeMetricRuleTargetsResponseBodyTargets) *DescribeMetricRuleTargetsResponseBody
	GetTargets() *DescribeMetricRuleTargetsResponseBodyTargets
}

type DescribeMetricRuleTargetsResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 786E92D2-AC66-4250-B76F-F1E2FCDDBA1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the resource for which alerts are triggered.
	Targets *DescribeMetricRuleTargetsResponseBodyTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricRuleTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleTargetsResponseBody) GetTargets() *DescribeMetricRuleTargetsResponseBodyTargets {
	return s.Targets
}

func (s *DescribeMetricRuleTargetsResponseBody) SetCode(v string) *DescribeMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetMessage(v string) *DescribeMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetRequestId(v string) *DescribeMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetSuccess(v bool) *DescribeMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetTargets(v *DescribeMetricRuleTargetsResponseBodyTargets) *DescribeMetricRuleTargetsResponseBody {
	s.Targets = v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleTargetsResponseBodyTargets struct {
	Target []*DescribeMetricRuleTargetsResponseBodyTargetsTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTargetsResponseBodyTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponseBodyTargets) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponseBodyTargets) GetTarget() []*DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	return s.Target
}

func (s *DescribeMetricRuleTargetsResponseBodyTargets) SetTarget(v []*DescribeMetricRuleTargetsResponseBodyTargetsTarget) *DescribeMetricRuleTargetsResponseBodyTargets {
	s.Target = v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargets) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleTargetsResponseBodyTargetsTarget struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource. Format: `acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. Fields:
	//
	// 	- {Service name abbreviation}: the abbreviation of the service name. Valid value: mns.
	//
	// 	- {userId}: the ID of the Alibaba Cloud account.
	//
	// 	- {regionId}: the region ID of the message queue or topic.
	//
	// 	- {Resource type}`: the type of the resource for which alerts are triggered. Valid values: - **queues*	- - **topics*	- {Resource name}: the name of the resource. - If the resource type is set to **queues**, the resource name is the name of the message queue. - If the resource type is set to **topics**, the resource name is the name of the topic.`
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the resource for which alerts are triggered.
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
	// The level of the alert. Valid values:
	//
	// 	- INFO: information
	//
	// 	- WARN: warning
	//
	// 	- CRITICAL: critical
	//
	// example:
	//
	// ["INFO", "WARN", "CRITICAL"]
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeMetricRuleTargetsResponseBodyTargetsTarget) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponseBodyTargetsTarget) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) GetArn() *string {
	return s.Arn
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) GetId() *string {
	return s.Id
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) GetJsonParams() *string {
	return s.JsonParams
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) GetLevel() *string {
	return s.Level
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetArn(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.Arn = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetId(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.Id = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetJsonParams(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.JsonParams = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetLevel(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.Level = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) Validate() error {
	return dara.Validate(s)
}
