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
	// The parameters for the baseline risk item that you want to fix.
	//
	// 	- **checkId**: the ID of the check item that corresponds to the baseline risk item.
	//
	// 	- **rules**: an array that consists of the rules applied to fixes.
	//
	//     	- **value**: specifies whether a fix method is selected. Valid values: **0*	- and **1**. The value 0 indicates that no fix method is selected and the value 1 indicates that a fix method is selected.
	//
	//     	- **ruleId**: the ID of the fix method.
	//
	//     	- **paramList**: an array that consists of the details about the fix method.\\
	//
	//         • **paramName**: the name of the fix method.\\
	//
	//         • **value**: the value of the fix method.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"checkId":8,"rules":[{"ruleId":"rule.ssh_Idle.interval","value":1,"paramList":[{"paramName":"range_val","value":"600"},{"paramName":"range_val","value":"600"}]},{"ruleId":"rule.ssh_Idle.count","value":1,"paramList":[{"paramName":"range_val","value":"3"}]}]}]
	CheckParams *string `json:"CheckParams,omitempty" xml:"CheckParams,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The retention period of the snapshot that is created when you fix the baseline risk. Valid values: 1 to 365. Unit: days.
	//
	// example:
	//
	// 1
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the risk item.
	//
	// >  To query the information about the risk items and check items of a server, you must specify the IDs of the risk items. You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the IDs of risk items.
	//
	// example:
	//
	// 10354
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the snapshot that is created when you fix the baseline risk.
	//
	// example:
	//
	// sas_fix_2024-12-04
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 165.225.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the asset for which you want to fix the baseline risk item. You can call the [DescribeWarningMachines](~~DescribeWarningMachines~~) operation to query the UUIDs of assets.
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
