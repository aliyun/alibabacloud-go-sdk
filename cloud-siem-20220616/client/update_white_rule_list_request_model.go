// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *UpdateWhiteRuleListRequest
	GetExpression() *string
	SetIncidentUuid(v string) *UpdateWhiteRuleListRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *UpdateWhiteRuleListRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateWhiteRuleListRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *UpdateWhiteRuleListRequest
	GetRoleType() *int32
	SetWhiteRuleId(v int64) *UpdateWhiteRuleListRequest
	GetWhiteRuleId() *int64
}

type UpdateWhiteRuleListRequest struct {
	// The alert whitelist rule. This is a JSON object.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "alertName": "webshell",
	//
	//             "alertNameId": "webshell",
	//
	//             "alertType": "command",
	//
	//             "alertTypeId": "command",
	//
	//             "expression": {
	//
	//                   "status": 1,
	//
	//                   "conditions": [
	//
	//                         {
	//
	//                               "isNot": false,
	//
	//                               "left": {
	//
	//                                     "value": "file_path"
	//
	//                               },
	//
	//                               "operator": "gt",
	//
	//                               "right": {
	//
	//                                     "value": "cp"
	//
	//                               }
	//
	//                         }
	//
	//                   ]
	//
	//             }
	//
	//       }
	//
	// ]
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The global unique ID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region of the Data Management center for threat analysis. Select a region for the Data Management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. This parameter is used when an administrator switches to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The unique ID of the whitelist rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	WhiteRuleId *int64 `json:"WhiteRuleId,omitempty" xml:"WhiteRuleId,omitempty"`
}

func (s UpdateWhiteRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListRequest) GetExpression() *string {
	return s.Expression
}

func (s *UpdateWhiteRuleListRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *UpdateWhiteRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWhiteRuleListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateWhiteRuleListRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *UpdateWhiteRuleListRequest) GetWhiteRuleId() *int64 {
	return s.WhiteRuleId
}

func (s *UpdateWhiteRuleListRequest) SetExpression(v string) *UpdateWhiteRuleListRequest {
	s.Expression = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetIncidentUuid(v string) *UpdateWhiteRuleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRegionId(v string) *UpdateWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRoleFor(v int64) *UpdateWhiteRuleListRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRoleType(v int32) *UpdateWhiteRuleListRequest {
	s.RoleType = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetWhiteRuleId(v int64) *UpdateWhiteRuleListRequest {
	s.WhiteRuleId = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) Validate() error {
	return dara.Validate(s)
}
