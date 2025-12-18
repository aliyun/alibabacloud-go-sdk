// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPreManagedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManagedRules(v []*ListPreManagedRulesResponseBodyManagedRules) *ListPreManagedRulesResponseBody
	GetManagedRules() []*ListPreManagedRulesResponseBodyManagedRules
	SetPageNumber(v int64) *ListPreManagedRulesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPreManagedRulesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListPreManagedRulesResponseBody
	GetRequestId() *string
}

type ListPreManagedRulesResponseBody struct {
	// The evaluation rules.
	ManagedRules []*ListPreManagedRulesResponseBodyManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A2A9F1BE-0712-1B26-9899-D82F7DA8476C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPreManagedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPreManagedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreManagedRulesResponseBody) GetManagedRules() []*ListPreManagedRulesResponseBodyManagedRules {
	return s.ManagedRules
}

func (s *ListPreManagedRulesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPreManagedRulesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPreManagedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPreManagedRulesResponseBody) SetManagedRules(v []*ListPreManagedRulesResponseBodyManagedRules) *ListPreManagedRulesResponseBody {
	s.ManagedRules = v
	return s
}

func (s *ListPreManagedRulesResponseBody) SetPageNumber(v int64) *ListPreManagedRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPreManagedRulesResponseBody) SetPageSize(v int64) *ListPreManagedRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPreManagedRulesResponseBody) SetRequestId(v string) *ListPreManagedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPreManagedRulesResponseBody) Validate() error {
	if s.ManagedRules != nil {
		for _, item := range s.ManagedRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPreManagedRulesResponseBodyManagedRules struct {
	// The details of the required input parameters of the rule.
	//
	// example:
	//
	// {}
	CompulsoryInputParameterDetails map[string]interface{} `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// ram-user-ak-used-expired-check
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// Example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the topic that describes how the evaluation rule remediates the incompliant configurations.
	//
	// example:
	//
	// https://example.aliyundoc.com
	HelpUrls *string `json:"HelpUrls,omitempty" xml:"HelpUrls,omitempty"`
	// The identifier of the rule.
	//
	// example:
	//
	// ecs-instance-deletion-protection-enabled
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The details of the optional input parameters of the rule.
	//
	// example:
	//
	// {}
	OptionalInputParameterDetails map[string]interface{} `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	// The type of resource.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListPreManagedRulesResponseBodyManagedRules) String() string {
	return dara.Prettify(s)
}

func (s ListPreManagedRulesResponseBodyManagedRules) GoString() string {
	return s.String()
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetCompulsoryInputParameterDetails() map[string]interface{} {
	return s.CompulsoryInputParameterDetails
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetDescription() *string {
	return s.Description
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetHelpUrls() *string {
	return s.HelpUrls
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetOptionalInputParameterDetails() map[string]interface{} {
	return s.OptionalInputParameterDetails
}

func (s *ListPreManagedRulesResponseBodyManagedRules) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetCompulsoryInputParameterDetails(v map[string]interface{}) *ListPreManagedRulesResponseBodyManagedRules {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetConfigRuleName(v string) *ListPreManagedRulesResponseBodyManagedRules {
	s.ConfigRuleName = &v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetDescription(v string) *ListPreManagedRulesResponseBodyManagedRules {
	s.Description = &v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetHelpUrls(v string) *ListPreManagedRulesResponseBodyManagedRules {
	s.HelpUrls = &v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetIdentifier(v string) *ListPreManagedRulesResponseBodyManagedRules {
	s.Identifier = &v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetOptionalInputParameterDetails(v map[string]interface{}) *ListPreManagedRulesResponseBodyManagedRules {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) SetResourceType(v string) *ListPreManagedRulesResponseBodyManagedRules {
	s.ResourceType = &v
	return s
}

func (s *ListPreManagedRulesResponseBodyManagedRules) Validate() error {
	return dara.Validate(s)
}
