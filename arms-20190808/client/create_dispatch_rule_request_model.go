// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRule(v string) *CreateDispatchRuleRequest
	GetDispatchRule() *string
	SetRegionId(v string) *CreateDispatchRuleRequest
	GetRegionId() *string
}

type CreateDispatchRuleRequest struct {
	// The dispatch rule configuration. The value is a JSON string. For more information about this parameter, see the following **additional information about the DispatchRule parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// {   "system": false,   "ruleid": 10282,   "name": "Prometheus Alert",   "labelMatchExpressionGrid": {     "labelMatchExpressionGroups": [       {         "labelMatchExpressions": [           {             "key": "_aliyun_arms_involvedObject_kind",             "value": "app",             "operator": "eq"           }         ]       }     ]   },   "dispatchType": "CREATE_ALERT/DISCARD_ALERT",   "isRecover": true,   "groupRules": [     {       "groupId": 1,       "groupingFields": [         "alertname"       ],       "groupWait": 10,       "groupInterval": 15,       "repeatInterval": 20     }   ],   "notifyRules": [     {       "notifyObjects": [         {           "notifyType": "ARMS_CONTACT",           "name": "JohnDoe",           "notifyObjectId": 1         },         {           "notifyType": "ARMS_CONTACT_GROUP",           "name": "JohnDoe_group",           "notifyObjectId": 2         }       ],       "notifyChannels":["dingTalk","wechat","webhook","email"]     },   ], }
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

func (s CreateDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleRequest) GetDispatchRule() *string {
	return s.DispatchRule
}

func (s *CreateDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDispatchRuleRequest) SetDispatchRule(v string) *CreateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *CreateDispatchRuleRequest) SetRegionId(v string) *CreateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}
