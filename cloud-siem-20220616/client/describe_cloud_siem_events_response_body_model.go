// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCloudSiemEventsResponseBody
	GetCode() *int32
	SetData(v *DescribeCloudSiemEventsResponseBodyData) *DescribeCloudSiemEventsResponseBody
	GetData() *DescribeCloudSiemEventsResponseBodyData
	SetMessage(v string) *DescribeCloudSiemEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudSiemEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudSiemEventsResponseBody
	GetSuccess() *bool
}

type DescribeCloudSiemEventsResponseBody struct {
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
	Data *DescribeCloudSiemEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeCloudSiemEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCloudSiemEventsResponseBody) GetData() *DescribeCloudSiemEventsResponseBodyData {
	return s.Data
}

func (s *DescribeCloudSiemEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudSiemEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudSiemEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudSiemEventsResponseBody) SetCode(v int32) *DescribeCloudSiemEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetData(v *DescribeCloudSiemEventsResponseBodyData) *DescribeCloudSiemEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetMessage(v string) *DescribeCloudSiemEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetRequestId(v string) *DescribeCloudSiemEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetSuccess(v bool) *DescribeCloudSiemEventsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudSiemEventsResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeCloudSiemEventsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeCloudSiemEventsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyData) GetPageInfo() *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeCloudSiemEventsResponseBodyData) GetResponseData() []*DescribeCloudSiemEventsResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeCloudSiemEventsResponseBodyData) SetPageInfo(v *DescribeCloudSiemEventsResponseBodyDataPageInfo) *DescribeCloudSiemEventsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyData) SetResponseData(v []*DescribeCloudSiemEventsResponseBodyDataResponseData) *DescribeCloudSiemEventsResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudSiemEventsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudSiemEventsResponseBodyDataResponseData struct {
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
	// The tags of the ATT\\&CK techniques.
	//
	// example:
	//
	// ["T1595.002 Vulnerability Scanning"]
	AttCkLabels []*string                                                         `json:"AttCkLabels,omitempty" xml:"AttCkLabels,omitempty" type:"Repeated"`
	AttckStages []*DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages `json:"AttckStages,omitempty" xml:"AttckStages,omitempty" type:"Repeated"`
	// The sources of the alert.
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
	// The event description in English.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The extended event information in the JSON format.
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
	// The event name in English.
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
	// the refer account info.
	//
	// example:
	//
	// 127608589417****
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
	// 	- 0: unhandled.
	//
	// 	- 1: handling.
	//
	// 	- 5: handling failed.
	//
	// 	- 10: handled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The risk score of the event. Valid values: 0 to 100. A higher value indicates a higher risk level.
	//
	// example:
	//
	// 90.2
	ThreatScore *float32 `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetAlertNum() *int32 {
	return s.AlertNum
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetAssetNum() *int32 {
	return s.AssetNum
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetAttCkLabels() []*string {
	return s.AttCkLabels
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetAttckStages() []*DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	return s.AttckStages
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetDataSources() []*string {
	return s.DataSources
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetExtContent() *string {
	return s.ExtContent
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetIncidentName() *string {
	return s.IncidentName
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetIncidentNameEn() *string {
	return s.IncidentNameEn
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetIncidentType() *string {
	return s.IncidentType
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetReferAccount() *string {
	return s.ReferAccount
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) GetThreatScore() *float32 {
	return s.ThreatScore
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAlertNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAliuid(v int64) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAssetNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAttCkLabels(v []*string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AttCkLabels = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAttckStages(v []*DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AttckStages = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDataSources(v []*string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.DataSources = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDescription(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Description = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDescriptionEn(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetExtContent(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ExtContent = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentName(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentName = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentNameEn(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentNameEn = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentType(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentType = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetReferAccount(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ReferAccount = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetRemark(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Remark = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetRuleId(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.RuleId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetStatus(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetThreatLevel(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetThreatScore(v float32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ThreatScore = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages struct {
	AlertNum   *int32  `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	TacticId   *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	TacticName *string `json:"TacticName,omitempty" xml:"TacticName,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) GetAlertNum() *int32 {
	return s.AlertNum
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) GetTacticName() *string {
	return s.TacticName
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) SetAlertNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) SetTacticId(v string) *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	s.TacticId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) SetTacticName(v string) *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	s.TacticName = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) Validate() error {
	return dara.Validate(s)
}
