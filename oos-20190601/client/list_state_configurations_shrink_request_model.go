// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStateConfigurationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListStateConfigurationsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListStateConfigurationsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListStateConfigurationsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListStateConfigurationsShrinkRequest
	GetResourceGroupId() *string
	SetStateConfigurationIds(v string) *ListStateConfigurationsShrinkRequest
	GetStateConfigurationIds() *string
	SetTagsShrink(v string) *ListStateConfigurationsShrinkRequest
	GetTagsShrink() *string
	SetTemplateName(v string) *ListStateConfigurationsShrinkRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *ListStateConfigurationsShrinkRequest
	GetTemplateVersion() *string
}

type ListStateConfigurationsShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s ListStateConfigurationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStateConfigurationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStateConfigurationsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStateConfigurationsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStateConfigurationsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStateConfigurationsShrinkRequest) GetStateConfigurationIds() *string {
	return s.StateConfigurationIds
}

func (s *ListStateConfigurationsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListStateConfigurationsShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListStateConfigurationsShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListStateConfigurationsShrinkRequest) SetMaxResults(v int32) *ListStateConfigurationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetNextToken(v string) *ListStateConfigurationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetRegionId(v string) *ListStateConfigurationsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetResourceGroupId(v string) *ListStateConfigurationsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetStateConfigurationIds(v string) *ListStateConfigurationsShrinkRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTagsShrink(v string) *ListStateConfigurationsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTemplateName(v string) *ListStateConfigurationsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTemplateVersion(v string) *ListStateConfigurationsShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
