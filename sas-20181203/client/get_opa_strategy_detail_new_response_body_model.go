// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaStrategyDetailNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpaStrategyDetailNewResponseBody
	GetCode() *string
	SetData(v *GetOpaStrategyDetailNewResponseBodyData) *GetOpaStrategyDetailNewResponseBody
	GetData() *GetOpaStrategyDetailNewResponseBodyData
	SetHttpStatusCode(v int32) *GetOpaStrategyDetailNewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOpaStrategyDetailNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpaStrategyDetailNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpaStrategyDetailNewResponseBody
	GetSuccess() *bool
}

type GetOpaStrategyDetailNewResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetOpaStrategyDetailNewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// E819FD71-D240-5E54-AA7F-20FED2ECBEB6
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

func (s GetOpaStrategyDetailNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpaStrategyDetailNewResponseBody) GetData() *GetOpaStrategyDetailNewResponseBodyData {
	return s.Data
}

func (s *GetOpaStrategyDetailNewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOpaStrategyDetailNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpaStrategyDetailNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaStrategyDetailNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpaStrategyDetailNewResponseBody) SetCode(v string) *GetOpaStrategyDetailNewResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBody) SetData(v *GetOpaStrategyDetailNewResponseBodyData) *GetOpaStrategyDetailNewResponseBody {
	s.Data = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBody) SetHttpStatusCode(v int32) *GetOpaStrategyDetailNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBody) SetMessage(v string) *GetOpaStrategyDetailNewResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBody) SetRequestId(v string) *GetOpaStrategyDetailNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBody) SetSuccess(v bool) *GetOpaStrategyDetailNewResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyData struct {
	// The rule configuration.
	AlarmDetail *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number. Default value: **1**. Pages start from page 1.
	//
	// example:
	//
	// 4
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description.
	//
	// example:
	//
	// Custom defense configuration
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image names.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The image tags.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Indicates whether the rule supports malicious Internet images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MaliciousImage *bool `json:"MaliciousImage,omitempty" xml:"MaliciousImage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The action that is performed when the rule is hit. Valid values:
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
	// The application scope.
	Scopes []*GetOpaStrategyDetailNewResponseBodyDataScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 1005
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// test001
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The ID of the rule template.
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
	// true
	UnScanedImage *bool `json:"UnScanedImage,omitempty" xml:"UnScanedImage,omitempty"`
	// The image tags that are added to the whitelist.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetAlarmDetail() *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail {
	return s.AlarmDetail
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetImageName() []*string {
	return s.ImageName
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetLabel() []*string {
	return s.Label
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetLang() *string {
	return s.Lang
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetScopes() []*GetOpaStrategyDetailNewResponseBodyDataScopes {
	return s.Scopes
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetStrategyName() *string {
	return s.StrategyName
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetStrategyTemplateId() *int64 {
	return s.StrategyTemplateId
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *GetOpaStrategyDetailNewResponseBodyData) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetAlarmDetail(v *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) *GetOpaStrategyDetailNewResponseBodyData {
	s.AlarmDetail = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetClusterId(v string) *GetOpaStrategyDetailNewResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetClusterName(v string) *GetOpaStrategyDetailNewResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetCurrentPage(v int32) *GetOpaStrategyDetailNewResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetDescription(v string) *GetOpaStrategyDetailNewResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetImageName(v []*string) *GetOpaStrategyDetailNewResponseBodyData {
	s.ImageName = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetLabel(v []*string) *GetOpaStrategyDetailNewResponseBodyData {
	s.Label = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetLang(v string) *GetOpaStrategyDetailNewResponseBodyData {
	s.Lang = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetMaliciousImage(v bool) *GetOpaStrategyDetailNewResponseBodyData {
	s.MaliciousImage = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetPageSize(v int32) *GetOpaStrategyDetailNewResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetRuleAction(v int32) *GetOpaStrategyDetailNewResponseBodyData {
	s.RuleAction = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetScopes(v []*GetOpaStrategyDetailNewResponseBodyDataScopes) *GetOpaStrategyDetailNewResponseBodyData {
	s.Scopes = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetStrategyId(v int64) *GetOpaStrategyDetailNewResponseBodyData {
	s.StrategyId = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetStrategyName(v string) *GetOpaStrategyDetailNewResponseBodyData {
	s.StrategyName = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetStrategyTemplateId(v int64) *GetOpaStrategyDetailNewResponseBodyData {
	s.StrategyTemplateId = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetUnScanedImage(v bool) *GetOpaStrategyDetailNewResponseBodyData {
	s.UnScanedImage = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) SetWhiteList(v []*string) *GetOpaStrategyDetailNewResponseBodyData {
	s.WhiteList = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyData) Validate() error {
	if s.AlarmDetail != nil {
		if err := s.AlarmDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Scopes != nil {
		for _, item := range s.Scopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetail struct {
	// The baseline check configuration.
	Baseline *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Struct"`
	// The configuration of image build risk.
	BuildRisk *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk `json:"BuildRisk,omitempty" xml:"BuildRisk,omitempty" type:"Struct"`
	// The configuration of malicious samples.
	MaliciousFile *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile `json:"MaliciousFile,omitempty" xml:"MaliciousFile,omitempty" type:"Struct"`
	// The configuration of sensitive file.
	SensitiveFile *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile `json:"SensitiveFile,omitempty" xml:"SensitiveFile,omitempty" type:"Struct"`
	// The vulnerability configuration.
	Vul *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul `json:"Vul,omitempty" xml:"Vul,omitempty" type:"Struct"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) GetBaseline() *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline {
	return s.Baseline
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) GetBuildRisk() *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk {
	return s.BuildRisk
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) GetMaliciousFile() *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile {
	return s.MaliciousFile
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) GetSensitiveFile() *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile {
	return s.SensitiveFile
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) GetVul() *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul {
	return s.Vul
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) SetBaseline(v *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail {
	s.Baseline = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) SetBuildRisk(v *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail {
	s.BuildRisk = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) SetMaliciousFile(v *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail {
	s.MaliciousFile = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) SetSensitiveFile(v *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail {
	s.SensitiveFile = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) SetVul(v *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail {
	s.Vul = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetail) Validate() error {
	if s.Baseline != nil {
		if err := s.Baseline.Validate(); err != nil {
			return err
		}
	}
	if s.BuildRisk != nil {
		if err := s.BuildRisk.Validate(); err != nil {
			return err
		}
	}
	if s.MaliciousFile != nil {
		if err := s.MaliciousFile.Validate(); err != nil {
			return err
		}
	}
	if s.SensitiveFile != nil {
		if err := s.SensitiveFile.Validate(); err != nil {
			return err
		}
	}
	if s.Vul != nil {
		if err := s.Vul.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline struct {
	// The information about the baseline check item.
	Item []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) GetItem() []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem {
	return s.Item
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) SetItem(v []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline {
	s.Item = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) SetRiskLevel(v []*string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline {
	s.RiskLevel = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaseline) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem struct {
	// The ID of the baseline check item.
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

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) GetId() *string {
	return s.Id
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) GetName() *string {
	return s.Name
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) SetId(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem {
	s.Id = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) SetName(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem {
	s.Name = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBaselineItem) Validate() error {
	return dara.Validate(s)
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk struct {
	// The configuration of image build risk.
	Item []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) GetItem() []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem {
	return s.Item
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) SetItem(v []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk {
	s.Item = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) SetRiskLevel(v []*string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk {
	s.RiskLevel = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRisk) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem struct {
	// The ID of the image build risk.
	//
	// >  You can call the [ListImageBuildRiskItem](~~ListImageBuildRiskItem~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// key
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the image build risk.
	//
	// >  You can call the [ListImageBuildRiskItem](~~ListImageBuildRiskItem~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) GetId() *string {
	return s.Id
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) GetName() *string {
	return s.Name
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) SetId(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem {
	s.Id = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) SetName(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem {
	s.Name = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailBuildRiskItem) Validate() error {
	return dara.Validate(s)
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile struct {
	// The information about the malicious sample.
	Item []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) GetItem() []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem {
	return s.Item
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) SetItem(v []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile {
	s.Item = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) SetRiskLevel(v []*string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile {
	s.RiskLevel = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFile) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem struct {
	// The ID of the malicious sample.
	//
	// example:
	//
	// 1811
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the malicious sample.
	//
	// example:
	//
	// abnormal binary file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) GetId() *string {
	return s.Id
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) GetName() *string {
	return s.Name
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) SetId(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem {
	s.Id = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) SetName(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem {
	s.Name = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailMaliciousFileItem) Validate() error {
	return dara.Validate(s)
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile struct {
	// The configuration of sensitive file.
	Item []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) GetItem() []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem {
	return s.Item
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) SetItem(v []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile {
	s.Item = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) SetRiskLevel(v []*string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile {
	s.RiskLevel = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFile) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem struct {
	// The ID of the sensitive files.
	//
	// >  You can call the [GetSensitiveDefineRuleConfig](~~GetSensitiveDefineRuleConfig~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// key
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the sensitive files.
	//
	// >  You can call the [GetSensitiveDefineRuleConfig](~~GetSensitiveDefineRuleConfig~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) GetId() *string {
	return s.Id
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) GetName() *string {
	return s.Name
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) SetId(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem {
	s.Id = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) SetName(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem {
	s.Name = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailSensitiveFileItem) Validate() error {
	return dara.Validate(s)
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul struct {
	// The information about the vulnerability.
	Item []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// Risk type of vulnerability.
	RiskClass []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass `json:"RiskClass,omitempty" xml:"RiskClass,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) GetItem() []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem {
	return s.Item
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) GetRiskClass() []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass {
	return s.RiskClass
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) SetItem(v []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul {
	s.Item = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) SetRiskClass(v []*GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul {
	s.RiskClass = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) SetRiskLevel(v []*string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul {
	s.RiskLevel = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVul) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskClass != nil {
		for _, item := range s.RiskClass {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem struct {
	// The ID of the vulnerability.
	//
	// example:
	//
	// AVD-2023-1680169
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// ezOffice evoInterfaceServlet Info Leak
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) GetId() *string {
	return s.Id
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) GetName() *string {
	return s.Name
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) SetId(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem {
	s.Id = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) SetName(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem {
	s.Name = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulItem) Validate() error {
	return dara.Validate(s)
}

type GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass struct {
	// The ID of the vulnerability types. Valid values:
	//
	// 	- **cve**: system vulnerability
	//
	// 	- **app**: application vulnerability
	//
	// example:
	//
	// cve
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability. Valid values:
	//
	// 	- **system vulnerability**
	//
	// 	- **application vulnerability**
	//
	// example:
	//
	// system vulnerability
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) GetId() *string {
	return s.Id
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) GetName() *string {
	return s.Name
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) SetId(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass {
	s.Id = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) SetName(v string) *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass {
	s.Name = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataAlarmDetailVulRiskClass) Validate() error {
	return dara.Validate(s)
}

type GetOpaStrategyDetailNewResponseBodyDataScopes struct {
	// The rule instance ID of the cluster.
	//
	// example:
	//
	// ack-0
	AckPolicyInstanceId *string `json:"AckPolicyInstanceId,omitempty" xml:"AckPolicyInstanceId,omitempty"`
	// Indicates whether all namespaces are included. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c1fdb5fd8d**7163
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
}

func (s GetOpaStrategyDetailNewResponseBodyDataScopes) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponseBodyDataScopes) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) GetAckPolicyInstanceId() *string {
	return s.AckPolicyInstanceId
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) SetAckPolicyInstanceId(v string) *GetOpaStrategyDetailNewResponseBodyDataScopes {
	s.AckPolicyInstanceId = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) SetAllNamespace(v int32) *GetOpaStrategyDetailNewResponseBodyDataScopes {
	s.AllNamespace = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) SetClusterId(v string) *GetOpaStrategyDetailNewResponseBodyDataScopes {
	s.ClusterId = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) SetNamespaceList(v []*string) *GetOpaStrategyDetailNewResponseBodyDataScopes {
	s.NamespaceList = v
	return s
}

func (s *GetOpaStrategyDetailNewResponseBodyDataScopes) Validate() error {
	return dara.Validate(s)
}
