// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *ListGrafanaWorkspaceRequest
	GetAliyunLang() *string
	SetRegionId(v string) *ListGrafanaWorkspaceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListGrafanaWorkspaceRequest
	GetResourceGroupId() *string
	SetTags(v []*ListGrafanaWorkspaceRequestTags) *ListGrafanaWorkspaceRequest
	GetTags() []*ListGrafanaWorkspaceRequestTags
}

type ListGrafanaWorkspaceRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The region ID. Default value: cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Prometheus instance belongs.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListGrafanaWorkspaceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListGrafanaWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *ListGrafanaWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGrafanaWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGrafanaWorkspaceRequest) GetTags() []*ListGrafanaWorkspaceRequestTags {
	return s.Tags
}

func (s *ListGrafanaWorkspaceRequest) SetAliyunLang(v string) *ListGrafanaWorkspaceRequest {
	s.AliyunLang = &v
	return s
}

func (s *ListGrafanaWorkspaceRequest) SetRegionId(v string) *ListGrafanaWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *ListGrafanaWorkspaceRequest) SetResourceGroupId(v string) *ListGrafanaWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGrafanaWorkspaceRequest) SetTags(v []*ListGrafanaWorkspaceRequestTags) *ListGrafanaWorkspaceRequest {
	s.Tags = v
	return s
}

func (s *ListGrafanaWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

type ListGrafanaWorkspaceRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListGrafanaWorkspaceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceRequestTags) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListGrafanaWorkspaceRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListGrafanaWorkspaceRequestTags) SetKey(v string) *ListGrafanaWorkspaceRequestTags {
	s.Key = &v
	return s
}

func (s *ListGrafanaWorkspaceRequestTags) SetValue(v string) *ListGrafanaWorkspaceRequestTags {
	s.Value = &v
	return s
}

func (s *ListGrafanaWorkspaceRequestTags) Validate() error {
	return dara.Validate(s)
}
