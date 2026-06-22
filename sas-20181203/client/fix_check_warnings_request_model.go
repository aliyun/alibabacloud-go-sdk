// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixCheckWarningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckParams(v string) *FixCheckWarningsRequest
	GetCheckParams() *string
	SetLang(v string) *FixCheckWarningsRequest
	GetLang() *string
	SetRetentionDays(v int32) *FixCheckWarningsRequest
	GetRetentionDays() *int32
	SetRiskId(v int64) *FixCheckWarningsRequest
	GetRiskId() *int64
	SetSnapshotName(v string) *FixCheckWarningsRequest
	GetSnapshotName() *string
	SetSourceIp(v string) *FixCheckWarningsRequest
	GetSourceIp() *string
	SetUuids(v string) *FixCheckWarningsRequest
	GetUuids() *string
}

type FixCheckWarningsRequest struct {
	// The parameters of the baseline check items to fix. The following parameters are included:
	//
	// - **checkId**: The ID of the check item.
	//
	// - **rules**: The fix rules (in array format).
	//
	//     - **value**: Specifies whether the fix method is selected. Valid values: **0*	- (not selected), **1*	- (selected).
	//
	//     - **ruleId**: The ID of the fix method.
	//
	//     - **paramList**: The list of fix methods (in array format).
	//
	//                 • **paramName**: The name of the fix method.
	//
	//                 • **value**: The value of the fix method.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"checkId":8,"rules":[{"ruleId":"rule.ssh_Idle.interval","value":1,"paramList":[{"paramName":"range_val","value":"600"},{"paramName":"range_val","value":"600"}]},{"ruleId":"rule.ssh_Idle.count","value":1,"paramList":[{"paramName":"range_val","value":"3"}]}]}]
	CheckParams *string `json:"CheckParams,omitempty" xml:"CheckParams,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The retention period of the snapshot to create when performing the baseline fix operation. Valid values: 1 to 365. Unit: days.
	//
	// example:
	//
	// 1
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the risk item.
	//
	// > To query the check item information for a specified risk item and a specified server, you must provide the risk item ID. You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to obtain the ID.
	//
	// example:
	//
	// 10354
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the snapshot to create when performing the baseline fix operation.
	//
	// example:
	//
	// sas_fix_2024-12-04
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 165.225.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The unique ID of the asset instance to fix. You can call the [DescribeWarningMachines](~~DescribeWarningMachines~~) operation to obtain the ID.
	//
	// example:
	//
	// 75a417dda5f25edb5bed8f208a9a****,c7e10fd794262a1510d5648f9e5d****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s FixCheckWarningsRequest) String() string {
	return dara.Prettify(s)
}

func (s FixCheckWarningsRequest) GoString() string {
	return s.String()
}

func (s *FixCheckWarningsRequest) GetCheckParams() *string {
	return s.CheckParams
}

func (s *FixCheckWarningsRequest) GetLang() *string {
	return s.Lang
}

func (s *FixCheckWarningsRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *FixCheckWarningsRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *FixCheckWarningsRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *FixCheckWarningsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *FixCheckWarningsRequest) GetUuids() *string {
	return s.Uuids
}

func (s *FixCheckWarningsRequest) SetCheckParams(v string) *FixCheckWarningsRequest {
	s.CheckParams = &v
	return s
}

func (s *FixCheckWarningsRequest) SetLang(v string) *FixCheckWarningsRequest {
	s.Lang = &v
	return s
}

func (s *FixCheckWarningsRequest) SetRetentionDays(v int32) *FixCheckWarningsRequest {
	s.RetentionDays = &v
	return s
}

func (s *FixCheckWarningsRequest) SetRiskId(v int64) *FixCheckWarningsRequest {
	s.RiskId = &v
	return s
}

func (s *FixCheckWarningsRequest) SetSnapshotName(v string) *FixCheckWarningsRequest {
	s.SnapshotName = &v
	return s
}

func (s *FixCheckWarningsRequest) SetSourceIp(v string) *FixCheckWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *FixCheckWarningsRequest) SetUuids(v string) *FixCheckWarningsRequest {
	s.Uuids = &v
	return s
}

func (s *FixCheckWarningsRequest) Validate() error {
	return dara.Validate(s)
}
