// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecuritySuggestionListResponseBody
	GetCode() *string
	SetData(v *GetSecuritySuggestionListResponseBodyData) *GetSecuritySuggestionListResponseBody
	GetData() *GetSecuritySuggestionListResponseBodyData
	SetMessage(v string) *GetSecuritySuggestionListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecuritySuggestionListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSecuritySuggestionListResponseBody
	GetSuccess() *bool
}

type GetSecuritySuggestionListResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSecuritySuggestionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 855FCC89-0B13-5FC0-AAD2-120878081C1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecuritySuggestionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecuritySuggestionListResponseBody) GetData() *GetSecuritySuggestionListResponseBodyData {
	return s.Data
}

func (s *GetSecuritySuggestionListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecuritySuggestionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecuritySuggestionListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecuritySuggestionListResponseBody) SetCode(v string) *GetSecuritySuggestionListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBody) SetData(v *GetSecuritySuggestionListResponseBodyData) *GetSecuritySuggestionListResponseBody {
	s.Data = v
	return s
}

func (s *GetSecuritySuggestionListResponseBody) SetMessage(v string) *GetSecuritySuggestionListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBody) SetRequestId(v string) *GetSecuritySuggestionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBody) SetSuccess(v bool) *GetSecuritySuggestionListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecuritySuggestionListResponseBodyData struct {
	ConfigRuleList []*GetSecuritySuggestionListResponseBodyDataConfigRuleList `json:"ConfigRuleList,omitempty" xml:"ConfigRuleList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 51
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSecuritySuggestionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponseBodyData) GetConfigRuleList() []*GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	return s.ConfigRuleList
}

func (s *GetSecuritySuggestionListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSecuritySuggestionListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSecuritySuggestionListResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSecuritySuggestionListResponseBodyData) SetConfigRuleList(v []*GetSecuritySuggestionListResponseBodyDataConfigRuleList) *GetSecuritySuggestionListResponseBodyData {
	s.ConfigRuleList = v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyData) SetPageNumber(v int32) *GetSecuritySuggestionListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyData) SetPageSize(v int32) *GetSecuritySuggestionListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyData) SetTotalCount(v int64) *GetSecuritySuggestionListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyData) Validate() error {
	if s.ConfigRuleList != nil {
		for _, item := range s.ConfigRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSecuritySuggestionListResponseBodyDataConfigRuleList struct {
	// example:
	//
	// 1625772519123804
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// OOS
	AutomationType *string `json:"AutomationType,omitempty" xml:"AutomationType,omitempty"`
	// example:
	//
	// {count=1, complianceType=NON_COMPLIANT}
	Compliance       *string                                                                  `json:"Compliance,omitempty" xml:"Compliance,omitempty"`
	ComplianceObject *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject `json:"ComplianceObject,omitempty" xml:"ComplianceObject,omitempty" type:"Struct"`
	// example:
	//
	// acs:config::100931896542****:rule/cr-fdc8626622af00f9****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// example:
	//
	// cr-bqa2f25bc5ce00af6323
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// example:
	//
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// example:
	//
	// ACTIVE
	ConfigRuleState *string                                                          `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	CreateBy        *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	// example:
	//
	// The description of the rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ACS::EIP::EipAddress
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// eip-bandwidth-limit
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// example:
	//
	// ALIYUN
	SourceOwner *string                                                        `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	Tags        []*GetSecuritySuggestionListResponseBodyDataConfigRuleListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleList) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetAutomationType() *string {
	return s.AutomationType
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetCompliance() *string {
	return s.Compliance
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetComplianceObject() *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject {
	return s.ComplianceObject
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetCreateBy() *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy {
	return s.CreateBy
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetDescription() *string {
	return s.Description
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) GetTags() []*GetSecuritySuggestionListResponseBodyDataConfigRuleListTags {
	return s.Tags
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetAccountId(v int64) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.AccountId = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetAutomationType(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.AutomationType = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetCompliance(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.Compliance = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetComplianceObject(v *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.ComplianceObject = v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetConfigRuleArn(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.ConfigRuleArn = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetConfigRuleId(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.ConfigRuleId = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetConfigRuleName(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetConfigRuleState(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.ConfigRuleState = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetCreateBy(v *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.CreateBy = v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetDescription(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.Description = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetResourceTypesScope(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetRiskLevel(v int32) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.RiskLevel = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetSourceIdentifier(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.SourceIdentifier = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetSourceOwner(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.SourceOwner = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) SetTags(v []*GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) *GetSecuritySuggestionListResponseBodyDataConfigRuleList {
	s.Tags = v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleList) Validate() error {
	if s.ComplianceObject != nil {
		if err := s.ComplianceObject.Validate(); err != nil {
			return err
		}
	}
	if s.CreateBy != nil {
		if err := s.CreateBy.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject struct {
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) GetCount() *int32 {
	return s.Count
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) SetComplianceType(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject {
	s.ComplianceType = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) SetCount(v int32) *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject {
	s.Count = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListComplianceObject) Validate() error {
	return dara.Validate(s)
}

type GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy struct {
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// example:
	//
	// The name of the compliance package.
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) SetCompliancePackId(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) SetCompliancePackName(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListCreateBy) Validate() error {
	return dara.Validate(s)
}

type GetSecuritySuggestionListResponseBodyDataConfigRuleListTags struct {
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) GetKey() *string {
	return s.Key
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) GetValue() *string {
	return s.Value
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) SetKey(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags {
	s.Key = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) SetValue(v string) *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags {
	s.Value = &v
	return s
}

func (s *GetSecuritySuggestionListResponseBodyDataConfigRuleListTags) Validate() error {
	return dara.Validate(s)
}
