// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListClustersRequest
	GetAcceptLanguage() *string
	SetClusterAliasName(v string) *ListClustersRequest
	GetClusterAliasName() *string
	SetKeyId(v string) *ListClustersRequest
	GetKeyId() *string
	SetPageNum(v int32) *ListClustersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListClustersRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListClustersRequest
	GetRequestPars() *string
	SetResourceGroupId(v string) *ListClustersRequest
	GetResourceGroupId() *string
	SetTag(v []*ListClustersRequestTag) *ListClustersRequest
	GetTag() []*ListClustersRequestTag
	SetVpcId(v string) *ListClustersRequest
	GetVpcId() *string
}

type ListClustersRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The alias of the instance. Fuzzy match is supported.
	//
	// example:
	//
	// cluster
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	KeyId            *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the instance resides. The region is supported by MSE.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The ID of the resource group. For more information about resource groups, see the topic "View basic information of a resource group."
	//
	// example:
	//
	// rg-acfmxbzafebvvfa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags. A maximum number of 20 tags are supported.
	Tag   []*ListClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VpcId *string                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListClustersRequest) GetClusterAliasName() *string {
	return s.ClusterAliasName
}

func (s *ListClustersRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ListClustersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClustersRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersRequest) GetTag() []*ListClustersRequestTag {
	return s.Tag
}

func (s *ListClustersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClustersRequest) SetAcceptLanguage(v string) *ListClustersRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListClustersRequest) SetClusterAliasName(v string) *ListClustersRequest {
	s.ClusterAliasName = &v
	return s
}

func (s *ListClustersRequest) SetKeyId(v string) *ListClustersRequest {
	s.KeyId = &v
	return s
}

func (s *ListClustersRequest) SetPageNum(v int32) *ListClustersRequest {
	s.PageNum = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListClustersRequest) SetRequestPars(v string) *ListClustersRequest {
	s.RequestPars = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersRequest) SetTag(v []*ListClustersRequestTag) *ListClustersRequest {
	s.Tag = v
	return s
}

func (s *ListClustersRequest) SetVpcId(v string) *ListClustersRequest {
	s.VpcId = &v
	return s
}

func (s *ListClustersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// prd
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequestTag) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListClustersRequestTag) SetKey(v string) *ListClustersRequestTag {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTag) SetValue(v string) *ListClustersRequestTag {
	s.Value = &v
	return s
}

func (s *ListClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
