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
	// The alert whitelist rule. The value is a JSON object.
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
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
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
