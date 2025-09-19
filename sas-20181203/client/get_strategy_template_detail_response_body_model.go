// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStrategyTemplateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStrategyTemplateDetailResponseBody
	GetCode() *string
	SetData(v *GetStrategyTemplateDetailResponseBodyData) *GetStrategyTemplateDetailResponseBody
	GetData() *GetStrategyTemplateDetailResponseBodyData
	SetMessage(v string) *GetStrategyTemplateDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStrategyTemplateDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStrategyTemplateDetailResponseBody
	GetSuccess() *bool
}

type GetStrategyTemplateDetailResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the template.
	Data *GetStrategyTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E36BEEA-0B27-58CC-8319-50279203B048
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStrategyTemplateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStrategyTemplateDetailResponseBody) GetData() *GetStrategyTemplateDetailResponseBodyData {
	return s.Data
}

func (s *GetStrategyTemplateDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStrategyTemplateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStrategyTemplateDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStrategyTemplateDetailResponseBody) SetCode(v string) *GetStrategyTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBody) SetData(v *GetStrategyTemplateDetailResponseBodyData) *GetStrategyTemplateDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBody) SetMessage(v string) *GetStrategyTemplateDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBody) SetRequestId(v string) *GetStrategyTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBody) SetSuccess(v bool) *GetStrategyTemplateDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyData struct {
	// The configuration of the rule.
	AlarmDetail *GetStrategyTemplateDetailResponseBodyDataAlarmDetail `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// hhht-cluster-02
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// Custom defense configuration
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The names of images.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The tags that are added to the containers.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// Indicates whether the rule supports malicious Internet images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	MaliciousImage *bool `json:"MaliciousImage,omitempty" xml:"MaliciousImage,omitempty"`
	// The namespaces.
	Namespace []*string `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Repeated"`
	// The action on requests. Valid values:
	//
	// 	- **1**: trigger alerts
	//
	// 	- **2**: block
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1005
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Blank template
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 1204
	StrategyTemplateId *int64 `json:"StrategyTemplateId,omitempty" xml:"StrategyTemplateId,omitempty"`
	// Indicates whether the rule supports unscanned images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	UnScanedImage *bool `json:"UnScanedImage,omitempty" xml:"UnScanedImage,omitempty"`
	// The whitelists of tags that are added to images.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s GetStrategyTemplateDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetAlarmDetail() *GetStrategyTemplateDetailResponseBodyDataAlarmDetail {
	return s.AlarmDetail
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetImageName() []*string {
	return s.ImageName
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetLabel() []*string {
	return s.Label
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetNamespace() []*string {
	return s.Namespace
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetStrategyName() *string {
	return s.StrategyName
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetStrategyTemplateId() *int64 {
	return s.StrategyTemplateId
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *GetStrategyTemplateDetailResponseBodyData) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetAlarmDetail(v *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) *GetStrategyTemplateDetailResponseBodyData {
	s.AlarmDetail = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetClusterId(v string) *GetStrategyTemplateDetailResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetClusterName(v string) *GetStrategyTemplateDetailResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetDescription(v string) *GetStrategyTemplateDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetImageName(v []*string) *GetStrategyTemplateDetailResponseBodyData {
	s.ImageName = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetLabel(v []*string) *GetStrategyTemplateDetailResponseBodyData {
	s.Label = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetMaliciousImage(v bool) *GetStrategyTemplateDetailResponseBodyData {
	s.MaliciousImage = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetNamespace(v []*string) *GetStrategyTemplateDetailResponseBodyData {
	s.Namespace = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetRuleAction(v int32) *GetStrategyTemplateDetailResponseBodyData {
	s.RuleAction = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetStrategyId(v int64) *GetStrategyTemplateDetailResponseBodyData {
	s.StrategyId = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetStrategyName(v string) *GetStrategyTemplateDetailResponseBodyData {
	s.StrategyName = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetStrategyTemplateId(v int64) *GetStrategyTemplateDetailResponseBodyData {
	s.StrategyTemplateId = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetUnScanedImage(v bool) *GetStrategyTemplateDetailResponseBodyData {
	s.UnScanedImage = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) SetWhiteList(v []*string) *GetStrategyTemplateDetailResponseBodyData {
	s.WhiteList = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetail struct {
	// The configuration of the baseline.
	Baseline *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Struct"`
	// The configuration of the alert rule for the malicious sample.
	MaliciousFile *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile `json:"MaliciousFile,omitempty" xml:"MaliciousFile,omitempty" type:"Struct"`
	// The configuration of the vulnerability detection rule.
	Vul *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul `json:"Vul,omitempty" xml:"Vul,omitempty" type:"Struct"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetail) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) GetBaseline() *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline {
	return s.Baseline
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) GetMaliciousFile() *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile {
	return s.MaliciousFile
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) GetVul() *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul {
	return s.Vul
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) SetBaseline(v *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) *GetStrategyTemplateDetailResponseBodyDataAlarmDetail {
	s.Baseline = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) SetMaliciousFile(v *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) *GetStrategyTemplateDetailResponseBodyDataAlarmDetail {
	s.MaliciousFile = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) SetVul(v *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) *GetStrategyTemplateDetailResponseBodyDataAlarmDetail {
	s.Vul = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetail) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline struct {
	// The baseline items.
	Item []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The severities of the baselines. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) GetItem() []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem {
	return s.Item
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) SetItem(v []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline {
	s.Item = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) SetRiskLevel(v []*string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline {
	s.RiskLevel = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaseline) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem struct {
	// The unique identifier of the baseline check item.
	//
	// example:
	//
	// ak_leak
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the baseline check item.
	//
	// example:
	//
	// Access Key plaintext storage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) GetId() *string {
	return s.Id
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) GetName() *string {
	return s.Name
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) SetId(v string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem {
	s.Id = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) SetName(v string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem {
	s.Name = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailBaselineItem) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile struct {
	// The items on which malicious samples are detected.
	Item []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The severities of the malicious samples.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) GetItem() []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem {
	return s.Item
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) SetItem(v []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile {
	s.Item = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) SetRiskLevel(v []*string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile {
	s.RiskLevel = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFile) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem struct {
	// The unique identifier of the malicious sample.
	//
	// example:
	//
	// test
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the malicious sample.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) GetId() *string {
	return s.Id
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) GetName() *string {
	return s.Name
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) SetId(v string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem {
	s.Id = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) SetName(v string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem {
	s.Name = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailMaliciousFileItem) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul struct {
	// The items on which vulnerabilities are detected.
	Item []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The severities of the vulnerabilities.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) GetItem() []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem {
	return s.Item
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) SetItem(v []*GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul {
	s.Item = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) SetRiskLevel(v []*string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul {
	s.RiskLevel = v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVul) Validate() error {
	return dara.Validate(s)
}

type GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem struct {
	// The ID of the vulnerability.
	//
	// example:
	//
	// AVD-2023-1680169
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The alias of the vulnerability.
	//
	// example:
	//
	// ezOffice evoInterfaceServlet Info Leak
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) GetId() *string {
	return s.Id
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) GetName() *string {
	return s.Name
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) SetId(v string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem {
	s.Id = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) SetName(v string) *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem {
	s.Name = &v
	return s
}

func (s *GetStrategyTemplateDetailResponseBodyDataAlarmDetailVulItem) Validate() error {
	return dara.Validate(s)
}
