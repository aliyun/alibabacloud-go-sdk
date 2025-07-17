// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstanceByTagAndResourceGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdRequest
	GetResourceGroupId() *string
	SetTag(v []*ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) *ListPrometheusInstanceByTagAndResourceGroupIdRequest
	GetTag() []*ListPrometheusInstanceByTagAndResourceGroupIdRequestTag
}

type ListPrometheusInstanceByTagAndResourceGroupIdRequest struct {
	// The region ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListPrometheusInstanceByTagAndResourceGroupIdRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) GetTag() []*ListPrometheusInstanceByTagAndResourceGroupIdRequestTag {
	return s.Tag
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) SetRegionId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) SetResourceGroupId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) SetTag(v []*ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) *ListPrometheusInstanceByTagAndResourceGroupIdRequest {
	s.Tag = v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequest) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusInstanceByTagAndResourceGroupIdRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// fpx-tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// fvt-tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) SetKey(v string) *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag {
	s.Key = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) SetValue(v string) *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag {
	s.Value = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdRequestTag) Validate() error {
	return dara.Validate(s)
}
