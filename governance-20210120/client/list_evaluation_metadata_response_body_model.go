// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadata) *ListEvaluationMetadataResponseBody
	GetEvaluationMetadata() []*ListEvaluationMetadataResponseBodyEvaluationMetadata
	SetRequestId(v string) *ListEvaluationMetadataResponseBody
	GetRequestId() *string
}

type ListEvaluationMetadataResponseBody struct {
	// The metadata of a governance maturity check.
	EvaluationMetadata []*ListEvaluationMetadataResponseBodyEvaluationMetadata `json:"EvaluationMetadata,omitempty" xml:"EvaluationMetadata,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 16B208DD-86BD-5E7D-AC93-FFD44B6FBDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEvaluationMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBody) GetEvaluationMetadata() []*ListEvaluationMetadataResponseBodyEvaluationMetadata {
	return s.EvaluationMetadata
}

func (s *ListEvaluationMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluationMetadataResponseBody) SetEvaluationMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadata) *ListEvaluationMetadataResponseBody {
	s.EvaluationMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBody) SetRequestId(v string) *ListEvaluationMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationMetadataResponseBody) Validate() error {
	if s.EvaluationMetadata != nil {
		for _, item := range s.EvaluationMetadata {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadata struct {
	// The metadata objects of a specific metadata type.
	Metadata []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Repeated"`
	// The type of the metadata. Valid values:
	//
	// 	- Metric: the check item
	//
	// example:
	//
	// Metric
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) GetMetadata() []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	return s.Metadata
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) GetType() *string {
	return s.Type
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) SetMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadata {
	s.Metadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) SetType(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadata {
	s.Type = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) Validate() error {
	if s.Metadata != nil {
		for _, item := range s.Metadata {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata struct {
	// The category of the check item.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// If you use an AccessKey pair of an Alibaba Cloud account, you have full permissions on the resources of the account. You cannot set limits on the account, such as setting limits on source IP addresses or access duration. If the AccessKey pair is leaked, resources within the account are exposed to high security risks. If your Alibaba Cloud account has an existing AccessKey pair, the check result is Non-compliant.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the check item.
	//
	// example:
	//
	// An AccessKey pair is enabled for the Alibaba Cloud account.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the metadata.
	//
	// example:
	//
	// pxgtda****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The governance level of the check item.
	//
	// example:
	//
	// High
	RecommendationLevel *string `json:"RecommendationLevel,omitempty" xml:"RecommendationLevel,omitempty"`
	// The metadata of the fixing task.
	RemediationMetadata *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata `json:"RemediationMetadata,omitempty" xml:"RemediationMetadata,omitempty" type:"Struct"`
	// The metadata of the checked resources.
	ResourceMetadata *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata `json:"ResourceMetadata,omitempty" xml:"ResourceMetadata,omitempty" type:"Struct"`
	// The scope of the check item. Valid values:
	//
	// 	- Account: the check item in a single-account governance maturity check
	//
	// 	- ResourceDirectory: the check item in a multi-account governance maturity check
	//
	// example:
	//
	// Account
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- Released: The check item is released.
	//
	// 	- Beta: The check item is pre-released.
	//
	// example:
	//
	// Released
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// example:
	//
	// ResourceUtilization
	TopicCode *string `json:"TopicCode,omitempty" xml:"TopicCode,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetCategory() *string {
	return s.Category
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetDescription() *string {
	return s.Description
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetId() *string {
	return s.Id
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetRecommendationLevel() *string {
	return s.RecommendationLevel
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetRemediationMetadata() *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata {
	return s.RemediationMetadata
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetResourceMetadata() *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata {
	return s.ResourceMetadata
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetScope() *string {
	return s.Scope
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetStage() *string {
	return s.Stage
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GetTopicCode() *string {
	return s.TopicCode
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetCategory(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Category = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetDescription(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Description = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetDisplayName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.DisplayName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetId(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Id = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetRecommendationLevel(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.RecommendationLevel = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetRemediationMetadata(v *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.RemediationMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetResourceMetadata(v *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.ResourceMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetScope(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Scope = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetStage(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Stage = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetTopicCode(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.TopicCode = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) Validate() error {
	if s.RemediationMetadata != nil {
		if err := s.RemediationMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceMetadata != nil {
		if err := s.ResourceMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata struct {
	// The fixing items.
	Remediation []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation `json:"Remediation,omitempty" xml:"Remediation,omitempty" type:"Repeated"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) GetRemediation() []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation {
	return s.Remediation
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) SetRemediation(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata {
	s.Remediation = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) Validate() error {
	if s.Remediation != nil {
		for _, item := range s.Remediation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation struct {
	// The fixing operations.
	Actions []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The type of the fixing method. Valid values:
	//
	// 	- Manual: manual fixing
	//
	// 	- QuickFix: quick fixing
	//
	// 	- Analysis: auxiliary decision-making
	//
	// example:
	//
	// Manual
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) GetActions() []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	return s.Actions
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) GetRemediationType() *string {
	return s.RemediationType
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) SetActions(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation {
	s.Actions = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) SetRemediationType(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation {
	s.RemediationType = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions struct {
	// The fixing method.
	//
	// >  This parameter is returned only if the value of `RemediationType` is `Analysis`.
	//
	// example:
	//
	// UnusedAccessKeyInRamUser
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// The fixing cost.
	//
	// example:
	//
	// You are not charged for this operation.
	CostDescription *string `json:"CostDescription,omitempty" xml:"CostDescription,omitempty"`
	// The description of the fixing item.
	//
	// >  This parameter is returned only if the value of `RemediationType` is `Analysis`.
	//
	// example:
	//
	// Console logon is enabled for the RAM user. The RAM user owns an AccessKey pair that is never used.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The content of the fixing items.
	Guidance []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance `json:"Guidance,omitempty" xml:"Guidance,omitempty" type:"Repeated"`
	// The usage notes of the fixing item.
	//
	// example:
	//
	// The BestPracticesForIdentityAndPermissions compliance package is enabled in Cloud Config to check the settings and usage of the AccessKey pair, Alibaba Cloud account, and RAM users.
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// The fixing suggestion.
	//
	// >  This parameter is returned only if the value of `RemediationType` is `Analysis`.
	//
	// example:
	//
	// Console logon is enabled for the RAM user and the RAM user owns an AccessKey pair, while the AccessKey pair has never been used by the RAM user. We recommend that you disable the AccessKey pair for 90 days. If no related issue occurs during this period, you can delete the AccessKey pair.
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GetClassification() *string {
	return s.Classification
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GetCostDescription() *string {
	return s.CostDescription
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GetDescription() *string {
	return s.Description
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GetGuidance() []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	return s.Guidance
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GetNotice() *string {
	return s.Notice
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetClassification(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Classification = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetCostDescription(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.CostDescription = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetDescription(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Description = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetGuidance(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Guidance = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetNotice(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Notice = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetSuggestion(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Suggestion = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) Validate() error {
	if s.Guidance != nil {
		for _, item := range s.Guidance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance struct {
	// The display name of the fixing button.
	//
	// example:
	//
	// Manual fixing
	ButtonName *string `json:"ButtonName,omitempty" xml:"ButtonName,omitempty"`
	// The navigation URL of the fixing button.
	//
	// example:
	//
	// https://ram.console.aliyun.com/users
	ButtonRef *string `json:"ButtonRef,omitempty" xml:"ButtonRef,omitempty"`
	// The fixing procedure.
	//
	// example:
	//
	// You must replace the AccessKey pair of your Alibaba Cloud account. To do so, perform the following steps:</br>1. Log on to the RAM console. In the left-side navigation pane, choose Identities > Users. On the Users page, click Create User.</br>2. On the Create User page, enter a logon name and select OpenAPI Access for the Access Mode parameter.</br>3. After the RAM user is created, save the AccessKey pair. Then, find the user that you created on the Users page and click Add Permissions in the Actions column. In the Grant Permission panel, find the AdministratorAccess policy and attach it to the RAM user.</br>4. In a program, replace the AccessKey pair of the Alibaba Cloud account with the AccessKey pair of the RAM user created in the previous step and check whether the program runs as expected in the test environment.</br>5. If the program runs as expected, publish the program to the production environment and disable the previous AccessKey pair of your Alibaba Cloud account. Then, check whether the program runs as expected.</br>6. If the program runs as expected, delete the disabled AccessKey pair after the specified period of time, such as 90 days.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The title of the fixing procedure.
	//
	// example:
	//
	// Scenario 3: AccessKey pair that is used within the last 90 days
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) GetButtonName() *string {
	return s.ButtonName
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) GetButtonRef() *string {
	return s.ButtonRef
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) GetContent() *string {
	return s.Content
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) GetTitle() *string {
	return s.Title
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetButtonName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.ButtonName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetButtonRef(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.ButtonRef = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetContent(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.Content = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetTitle(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.Title = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) Validate() error {
	return dara.Validate(s)
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata struct {
	// The metadata of the resource properties.
	ResourcePropertyMetadata []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata `json:"ResourcePropertyMetadata,omitempty" xml:"ResourcePropertyMetadata,omitempty" type:"Repeated"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) GetResourcePropertyMetadata() []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	return s.ResourcePropertyMetadata
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) SetResourcePropertyMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata {
	s.ResourcePropertyMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) Validate() error {
	if s.ResourcePropertyMetadata != nil {
		for _, item := range s.ResourcePropertyMetadata {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata struct {
	// The display name of the resource property.
	//
	// example:
	//
	// AccessKey Pair Last Used At
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the resource property.
	//
	// example:
	//
	// AkLastUsedTime
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The type of the resource property.
	//
	// example:
	//
	// String
	PropertyType *string `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) GetPropertyName() *string {
	return s.PropertyName
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) GetPropertyType() *string {
	return s.PropertyType
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) SetDisplayName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	s.DisplayName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) SetPropertyName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	s.PropertyName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) SetPropertyType(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	s.PropertyType = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) Validate() error {
	return dara.Validate(s)
}
