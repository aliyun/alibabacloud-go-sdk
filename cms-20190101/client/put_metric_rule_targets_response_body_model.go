// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMetricRuleTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutMetricRuleTargetsResponseBody
	GetCode() *string
	SetFailData(v *PutMetricRuleTargetsResponseBodyFailData) *PutMetricRuleTargetsResponseBody
	GetFailData() *PutMetricRuleTargetsResponseBodyFailData
	SetMessage(v string) *PutMetricRuleTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutMetricRuleTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutMetricRuleTargetsResponseBody
	GetSuccess() *bool
}

type PutMetricRuleTargetsResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The failed data.
	FailData *PutMetricRuleTargetsResponseBodyFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A569B0D-9055-58AF-9E82-BAEAF95C0FD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutMetricRuleTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutMetricRuleTargetsResponseBody) GetFailData() *PutMetricRuleTargetsResponseBodyFailData {
	return s.FailData
}

func (s *PutMetricRuleTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutMetricRuleTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutMetricRuleTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutMetricRuleTargetsResponseBody) SetCode(v string) *PutMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetFailData(v *PutMetricRuleTargetsResponseBodyFailData) *PutMetricRuleTargetsResponseBody {
	s.FailData = v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetMessage(v string) *PutMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetRequestId(v string) *PutMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetSuccess(v bool) *PutMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) Validate() error {
	if s.FailData != nil {
		if err := s.FailData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutMetricRuleTargetsResponseBodyFailData struct {
	// The information about the resources for which alerts are triggered.
	Targets *PutMetricRuleTargetsResponseBodyFailDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Struct"`
}

func (s PutMetricRuleTargetsResponseBodyFailData) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailData) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailData) GetTargets() *PutMetricRuleTargetsResponseBodyFailDataTargets {
	return s.Targets
}

func (s *PutMetricRuleTargetsResponseBodyFailData) SetTargets(v *PutMetricRuleTargetsResponseBodyFailDataTargets) *PutMetricRuleTargetsResponseBodyFailData {
	s.Targets = v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailData) Validate() error {
	if s.Targets != nil {
		if err := s.Targets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutMetricRuleTargetsResponseBodyFailDataTargets struct {
	Target []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargets) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargets) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargets) GetTarget() []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	return s.Target
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargets) SetTarget(v []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) *PutMetricRuleTargetsResponseBodyFailDataTargets {
	s.Target = v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargets) Validate() error {
	if s.Target != nil {
		for _, item := range s.Target {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutMetricRuleTargetsResponseBodyFailDataTargetsTarget struct {
	// The ARN of the resource. Format: `acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message`. SMQ, Auto Scaling, Simple Log Service, and Function Compute are supported. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. The following part describes the ARN of SMQ and the parameters in the ARN:
	//
	// 	- {Service name abbreviation}: mns.
	//
	// 	- {userId}: the ID of the Alibaba Cloud account.
	//
	// 	- {regionId}: the region ID of the SMQ queue or topic.
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
	// example:
	//
	// acs:mns:cn-hangzhou:111:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the resource for which alerts are triggered.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) GetArn() *string {
	return s.Arn
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) GetId() *string {
	return s.Id
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) GetLevel() *string {
	return s.Level
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetArn(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Arn = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetId(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Id = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetLevel(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Level = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) Validate() error {
	return dara.Validate(s)
}
