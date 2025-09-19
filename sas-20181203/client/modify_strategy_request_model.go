// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomType(v string) *ModifyStrategyRequest
	GetCustomType() *string
	SetCycleDays(v string) *ModifyStrategyRequest
	GetCycleDays() *string
	SetCycleStartTime(v string) *ModifyStrategyRequest
	GetCycleStartTime() *string
	SetEndTime(v string) *ModifyStrategyRequest
	GetEndTime() *string
	SetId(v string) *ModifyStrategyRequest
	GetId() *string
	SetName(v string) *ModifyStrategyRequest
	GetName() *string
	SetRiskCustomParams(v string) *ModifyStrategyRequest
	GetRiskCustomParams() *string
	SetRiskSubTypeName(v string) *ModifyStrategyRequest
	GetRiskSubTypeName() *string
	SetSourceIp(v string) *ModifyStrategyRequest
	GetSourceIp() *string
	SetStartTime(v string) *ModifyStrategyRequest
	GetStartTime() *string
	SetTargetType(v string) *ModifyStrategyRequest
	GetTargetType() *string
}

type ModifyStrategyRequest struct {
	// The type of the baseline check policy. Valid values:
	//
	// 	- **custom**: a custom baseline check policy
	//
	// 	- **common**: a standard baseline check policy
	//
	// This parameter is required.
	//
	// example:
	//
	// common
	CustomType *string `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The new interval of the baseline check. Valid values:
	//
	// 	- **1**: every 2 days
	//
	// 	- **3**: every 4 days
	//
	// 	- **7**: every 8 days
	//
	// 	- **30**: every 31 days
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CycleDays *string `json:"CycleDays,omitempty" xml:"CycleDays,omitempty"`
	// The new time range during which the baseline check starts. Valid values:
	//
	// 	- **0**: The baseline check starts within the time range from 00:00 to 06:00.
	//
	// 	- **6**: The baseline check starts within the time range from 06:00 to 12:00.
	//
	// 	- **12**: The baseline check starts within the time range from 12:00 to 18:00.
	//
	// 	- **18**: The baseline check starts within the time range from 18:00 to 24:00.
	//
	// >  This parameter is deprecated.
	//
	// example:
	//
	// 18
	CycleStartTime *string `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	// The time when the baseline check based on the baseline check policy ends. Specify the time in the hh:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 245
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name of the baseline check policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// testStrategy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom configurations of the baseline. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **typeName**: the name of the baseline.
	//
	// 	- **checkDetails**: the details of the baseline. The value is in the JSON format.
	//
	//     	- **checkId**: the ID of the check item.
	//
	//     	- **rules**: the rule configurations. The value is in the JSON format.
	//
	//         	- **ruleId**: the ID of the rule.
	//
	//         	- **paramList**: the list of parameters in the rule. The value is in the JSON format.
	//
	//             	- **paramName**: the name of the parameter.
	//
	//             	- **value**: the value of the parameter.
	//
	// example:
	//
	// [{"typeName":"hc_centos_6_custom","checkDetails":[{"checkId":4,"rules":[{"ruleId":"pass_min_days_login_defs.must.cus","paramList":[{"paramName":"range_val","value":"7"}]}]}]}]
	RiskCustomParams *string `json:"RiskCustomParams,omitempty" xml:"RiskCustomParams,omitempty"`
	// The subtype of the baselines. You can call the [DescribeRiskType](~~DescribeRiskType~~) operation to query the subtypes of baselines.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc_exploit_redis
	RiskSubTypeName *string `json:"RiskSubTypeName,omitempty" xml:"RiskSubTypeName,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The time when the baseline check based on the baseline check policy starts. Specify the time in the hh:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:01:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The method that is used to apply the baseline check policy. Valid values:
	//
	// 	- **groupId**: asset groups
	//
	// 	- **uuid**: assets
	//
	// This parameter is required.
	//
	// example:
	//
	// groupId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ModifyStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyStrategyRequest) GetCustomType() *string {
	return s.CustomType
}

func (s *ModifyStrategyRequest) GetCycleDays() *string {
	return s.CycleDays
}

func (s *ModifyStrategyRequest) GetCycleStartTime() *string {
	return s.CycleStartTime
}

func (s *ModifyStrategyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyStrategyRequest) GetId() *string {
	return s.Id
}

func (s *ModifyStrategyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyStrategyRequest) GetRiskCustomParams() *string {
	return s.RiskCustomParams
}

func (s *ModifyStrategyRequest) GetRiskSubTypeName() *string {
	return s.RiskSubTypeName
}

func (s *ModifyStrategyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyStrategyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyStrategyRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ModifyStrategyRequest) SetCustomType(v string) *ModifyStrategyRequest {
	s.CustomType = &v
	return s
}

func (s *ModifyStrategyRequest) SetCycleDays(v string) *ModifyStrategyRequest {
	s.CycleDays = &v
	return s
}

func (s *ModifyStrategyRequest) SetCycleStartTime(v string) *ModifyStrategyRequest {
	s.CycleStartTime = &v
	return s
}

func (s *ModifyStrategyRequest) SetEndTime(v string) *ModifyStrategyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyStrategyRequest) SetId(v string) *ModifyStrategyRequest {
	s.Id = &v
	return s
}

func (s *ModifyStrategyRequest) SetName(v string) *ModifyStrategyRequest {
	s.Name = &v
	return s
}

func (s *ModifyStrategyRequest) SetRiskCustomParams(v string) *ModifyStrategyRequest {
	s.RiskCustomParams = &v
	return s
}

func (s *ModifyStrategyRequest) SetRiskSubTypeName(v string) *ModifyStrategyRequest {
	s.RiskSubTypeName = &v
	return s
}

func (s *ModifyStrategyRequest) SetSourceIp(v string) *ModifyStrategyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyStrategyRequest) SetStartTime(v string) *ModifyStrategyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyStrategyRequest) SetTargetType(v string) *ModifyStrategyRequest {
	s.TargetType = &v
	return s
}

func (s *ModifyStrategyRequest) Validate() error {
	return dara.Validate(s)
}
