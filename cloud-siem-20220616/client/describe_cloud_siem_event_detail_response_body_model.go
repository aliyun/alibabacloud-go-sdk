// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCloudSiemEventDetailResponseBody
	GetCode() *int32
	SetData(v *DescribeCloudSiemEventDetailResponseBodyData) *DescribeCloudSiemEventDetailResponseBody
	GetData() *DescribeCloudSiemEventDetailResponseBodyData
	SetMessage(v string) *DescribeCloudSiemEventDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudSiemEventDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudSiemEventDetailResponseBody
	GetSuccess() *bool
}

type DescribeCloudSiemEventDetailResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCloudSiemEventDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCloudSiemEventDetailResponseBody) GetData() *DescribeCloudSiemEventDetailResponseBodyData {
	return s.Data
}

func (s *DescribeCloudSiemEventDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudSiemEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudSiemEventDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetCode(v int32) *DescribeCloudSiemEventDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetData(v *DescribeCloudSiemEventDetailResponseBodyData) *DescribeCloudSiemEventDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetMessage(v string) *DescribeCloudSiemEventDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetRequestId(v string) *DescribeCloudSiemEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetSuccess(v bool) *DescribeCloudSiemEventDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudSiemEventDetailResponseBodyData struct {
	// The number of alerts that are associated with the event.
	//
	// example:
	//
	// 4
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// The ID of the Alibaba Cloud account to which the event belongs.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The number of assets that are associated with the event.
	//
	// example:
	//
	// 4
	AssetNum *int32 `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	// The tags of the ATT\\&CK attacks.
	//
	// example:
	//
	// ["T1595.002 Vulnerability Scanning"]
	AttCkLabels []*string                                                  `json:"AttCkLabels,omitempty" xml:"AttCkLabels,omitempty" type:"Repeated"`
	AttckStages []*DescribeCloudSiemEventDetailResponseBodyDataAttckStages `json:"AttckStages,omitempty" xml:"AttckStages,omitempty" type:"Repeated"`
	// The source of the alert.
	//
	// example:
	//
	// [sas,waf]
	DataSources []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The description of the event.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the event in English.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The extended information of the event in the JSON format.
	//
	// example:
	//
	// {"event_transfer_type":"customize_rule"}
	ExtContent *string `json:"ExtContent,omitempty" xml:"ExtContent,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the event was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The name of the event in English.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentNameEn *string `json:"IncidentNameEn,omitempty" xml:"IncidentNameEn,omitempty"`
	IncidentType   *string `json:"IncidentType,omitempty" xml:"IncidentType,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Users associated with the event.
	ReferAccount *string `json:"ReferAccount,omitempty" xml:"ReferAccount,omitempty"`
	// The remarks of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: not handled
	//
	// 	- 1: handing
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The risk score of the event. The score ranges from 0 to 100. A higher score indicates a higher risk level.
	//
	// example:
	//
	// 90.2
	ThreatScore *float32 `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetAlertNum() *int32 {
	return s.AlertNum
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetAssetNum() *int32 {
	return s.AssetNum
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetAttCkLabels() []*string {
	return s.AttCkLabels
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetAttckStages() []*DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	return s.AttckStages
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetDataSources() []*string {
	return s.DataSources
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetExtContent() *string {
	return s.ExtContent
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetIncidentName() *string {
	return s.IncidentName
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetIncidentNameEn() *string {
	return s.IncidentNameEn
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetIncidentType() *string {
	return s.IncidentType
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetReferAccount() *string {
	return s.ReferAccount
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) GetThreatScore() *float32 {
	return s.ThreatScore
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAlertNum(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAliuid(v int64) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAssetNum(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAttCkLabels(v []*string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AttCkLabels = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAttckStages(v []*DescribeCloudSiemEventDetailResponseBodyDataAttckStages) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AttckStages = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDataSources(v []*string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.DataSources = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDescription(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDescriptionEn(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetExtContent(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ExtContent = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetGmtCreate(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetGmtModified(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentName(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentName = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentNameEn(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentNameEn = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentType(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentType = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentUuid(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetReferAccount(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ReferAccount = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetRemark(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetRuleId(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetStatus(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetThreatLevel(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetThreatScore(v float32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ThreatScore = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) Validate() error {
	if s.AttckStages != nil {
		for _, item := range s.AttckStages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudSiemEventDetailResponseBodyDataAttckStages struct {
	AlertNum   *int32  `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	TacticId   *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	TacticName *string `json:"TacticName,omitempty" xml:"TacticName,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBodyDataAttckStages) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBodyDataAttckStages) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) GetAlertNum() *int32 {
	return s.AlertNum
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) GetTacticName() *string {
	return s.TacticName
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) SetAlertNum(v int32) *DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) SetTacticId(v string) *DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	s.TacticId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) SetTacticName(v string) *DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	s.TacticName = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) Validate() error {
	return dara.Validate(s)
}
