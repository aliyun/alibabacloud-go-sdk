// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationMetricDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListEvaluationMetricDetailsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEvaluationMetricDetailsResponseBody
	GetRequestId() *string
	SetResources(v []*ListEvaluationMetricDetailsResponseBodyResources) *ListEvaluationMetricDetailsResponseBody
	GetResources() []*ListEvaluationMetricDetailsResponseBodyResources
}

type ListEvaluationMetricDetailsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAGEaXR18y1rqykZHIqRuBejOqED4S3Xne33c7zbn****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC9BD94C-D20C-4D27-88D4-89E8D75C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the non-compliant resources.
	Resources []*ListEvaluationMetricDetailsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListEvaluationMetricDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluationMetricDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluationMetricDetailsResponseBody) GetResources() []*ListEvaluationMetricDetailsResponseBodyResources {
	return s.Resources
}

func (s *ListEvaluationMetricDetailsResponseBody) SetNextToken(v string) *ListEvaluationMetricDetailsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBody) SetRequestId(v string) *ListEvaluationMetricDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBody) SetResources(v []*ListEvaluationMetricDetailsResponseBodyResources) *ListEvaluationMetricDetailsResponseBody {
	s.Resources = v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetricDetailsResponseBodyResources struct {
	// The compliance status of the resource. Valid values:
	//
	// 	- NonCompliant: non-compliant.
	//
	// 	- Excluded: ignored.
	//
	// 	- PendingExclusion: to be ignored.
	//
	// 	- PendingInclusion: to be unignored.
	//
	// example:
	//
	// NonCompliant
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The check results further analyzed by auxiliary decision-making.
	//
	// >  This parameter is returned only when the check item supports the auxiliary decision-making feature.
	//
	// example:
	//
	// RecentUnloginRamUser
	ResourceClassification *string `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// 26435103783237****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 176618589410****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The attributes of the resource.
	ResourceProperties []*ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty" type:"Repeated"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListEvaluationMetricDetailsResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetResourceClassification() *string {
	return s.ResourceClassification
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetResourceProperties() []*ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties {
	return s.ResourceProperties
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetComplianceType(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ComplianceType = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetRegionId(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceClassification(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceClassification = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceId(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceName(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceOwnerId(v int64) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceProperties(v []*ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceProperties = v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceType(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) Validate() error {
	if s.ResourceProperties != nil {
		for _, item := range s.ResourceProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties struct {
	// The name of the resource attribute.
	//
	// example:
	//
	// DisplayName
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The value of the resource attribute.
	//
	// example:
	//
	// example
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) GetPropertyName() *string {
	return s.PropertyName
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) SetPropertyName(v string) *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties {
	s.PropertyName = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) SetPropertyValue(v string) *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties {
	s.PropertyValue = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) Validate() error {
	return dara.Validate(s)
}
