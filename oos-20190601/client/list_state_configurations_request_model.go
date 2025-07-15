// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStateConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListStateConfigurationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListStateConfigurationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListStateConfigurationsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListStateConfigurationsRequest
	GetResourceGroupId() *string
	SetStateConfigurationIds(v string) *ListStateConfigurationsRequest
	GetStateConfigurationIds() *string
	SetTags(v map[string]interface{}) *ListStateConfigurationsRequest
	GetTags() map[string]interface{}
	SetTemplateName(v string) *ListStateConfigurationsRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *ListStateConfigurationsRequest
	GetTemplateVersion() *string
}

type ListStateConfigurationsRequest struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AHJKH-AHKJHDJK-AKHDIOWJL
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the desired-state configuration.
	//
	// example:
	//
	// ["sc-asfgdhj12345"]
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template. If you do not specify this parameter, the latest version of the template is used.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListStateConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStateConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStateConfigurationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStateConfigurationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStateConfigurationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStateConfigurationsRequest) GetStateConfigurationIds() *string {
	return s.StateConfigurationIds
}

func (s *ListStateConfigurationsRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListStateConfigurationsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListStateConfigurationsRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListStateConfigurationsRequest) SetMaxResults(v int32) *ListStateConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetNextToken(v string) *ListStateConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetRegionId(v string) *ListStateConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetResourceGroupId(v string) *ListStateConfigurationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetStateConfigurationIds(v string) *ListStateConfigurationsRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTags(v map[string]interface{}) *ListStateConfigurationsRequest {
	s.Tags = v
	return s
}

func (s *ListStateConfigurationsRequest) SetTemplateName(v string) *ListStateConfigurationsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTemplateVersion(v string) *ListStateConfigurationsRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
