// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRule(v string) *UpdateDispatchRuleRequest
	GetDispatchRule() *string
	SetRegionId(v string) *UpdateDispatchRuleRequest
	GetRegionId() *string
}

type UpdateDispatchRuleRequest struct {
	// The dispatch rule configuration. The value is a JSON string. For more information about this parameter, see the following **additional information about the DispatchRule parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// {   "id": 123,     "system": false,   "ruleid": 10282,   "name": "Prometheus Alert",   "labelMatchExpressionGrid": {     "labelMatchExpressionGroups": [       {         "labelMatchExpressions": [           {             "key": "_aliyun_arms_involvedObject_kind",             "value": "app",             "operator": "eq"           }         ]       }     ]   },   "dispatchType": "CREATE_ALERT/DISCARD_ALERT",   "isRecover": true,   "groupRules": [     {       "groupId": 1,       "groupingFields": [         "alertname"       ],       "groupWait": 10,       "groupInterval": 15,       "repeatInterval": 20     }   ],   "notifyRules": [     {       "notifyObjects": [         {           "notifyType": "ARMS_CONTACT",           "name": "JohnDoe",           "notifyObjectId": 1         },         {           "notifyType": "ARMS_CONTACT_GROUP",           "name": "JohnDoe_group",           "notifyObjectId": 2         }       ],       "notifyChannels":["dingTalk","wechat","webhook","email"]     },   ], }
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleRequest) GetDispatchRule() *string {
	return s.DispatchRule
}

func (s *UpdateDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDispatchRuleRequest) SetDispatchRule(v string) *UpdateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *UpdateDispatchRuleRequest) SetRegionId(v string) *UpdateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}
