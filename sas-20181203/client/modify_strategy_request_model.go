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
	// The policy type. Valid values:
	//
	// - **custom**: custom policy.
	//
	// - **common**: standard policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// common
	CustomType *string `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The cycle of the baseline check. Valid values:
	//
	// - **1**: Every 1 day.
	//
	// - **3**: Every 3 days.
	//
	// - **7**: Every 7 days.
	//
	// - **30**: Every 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CycleDays *string `json:"CycleDays,omitempty" xml:"CycleDays,omitempty"`
	// The start time of the baseline check. Valid values:
	//
	// - **0**: The baseline check starts between 00:00 and 06:00.
	//
	// - **6**: The baseline check starts between 06:00 and 12:00.
	//
	// - **12**: The baseline check starts between 12:00 and 18:00.
	//
	// - **18**: The baseline check starts between 18:00 and 24:00.
	//
	// > This parameter is deprecated.
	//
	// The value indicates the start hour of the daily check period, in hours.
	//
	// example:
	//
	// 18
	CycleStartTime *string `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	// The end time of the policy execution. Format: hh:mm:ss.
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
	// The name of the baseline check policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// testStrategy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom configuration of baseline check items. The value is in JSON format and contains the following parameters:
	//
	// - **typeName**: The baseline name.
	//
	// - **checkDetails**: The check details. The value is in JSON format.
	//
	//     - **checkId**: The ID of the check item.
	//
	//     - **rules**: The policy configuration. The value is in JSON format.
	//
	//         - **ruleId**: The ID of the policy configuration.
	//
	//         - **paramList**: The collection of policy parameter settings. The value is in JSON format.
	//
	//             - **paramName**: The parameter name.
	//
	//             - **value**: The parameter settings value.
	//
	// example:
	//
	// [{"typeName":"hc_centos_6_custom","checkDetails":[{"checkId":4,"rules":[{"ruleId":"pass_min_days_login_defs.must.cus","paramList":[{"paramName":"range_val","value":"7"}]}]}]}]
	RiskCustomParams *string `json:"RiskCustomParams,omitempty" xml:"RiskCustomParams,omitempty"`
	// The subtype of the check item. You can call the [DescribeRiskType](~~DescribeRiskType~~) operation to obtain the subtype.
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
	// The start time of the policy execution. Format: hh:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:01:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The scan method of the policy. Valid values:
	//
	// - **groupId**: group-based scan.
	//
	// - **uuid**: asset-based scan.
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
