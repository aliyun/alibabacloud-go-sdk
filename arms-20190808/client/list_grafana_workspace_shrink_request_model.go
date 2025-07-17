// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *ListGrafanaWorkspaceShrinkRequest
	GetAliyunLang() *string
	SetRegionId(v string) *ListGrafanaWorkspaceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListGrafanaWorkspaceShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *ListGrafanaWorkspaceShrinkRequest
	GetTagsShrink() *string
}

type ListGrafanaWorkspaceShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListGrafanaWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *ListGrafanaWorkspaceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGrafanaWorkspaceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGrafanaWorkspaceShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListGrafanaWorkspaceShrinkRequest) SetAliyunLang(v string) *ListGrafanaWorkspaceShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *ListGrafanaWorkspaceShrinkRequest) SetRegionId(v string) *ListGrafanaWorkspaceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListGrafanaWorkspaceShrinkRequest) SetResourceGroupId(v string) *ListGrafanaWorkspaceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGrafanaWorkspaceShrinkRequest) SetTagsShrink(v string) *ListGrafanaWorkspaceShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListGrafanaWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
