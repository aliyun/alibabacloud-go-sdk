// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *ListApiTemplatesRequest
	GetApiName() *string
	SetMaxResults(v int32) *ListApiTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiTemplatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListApiTemplatesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListApiTemplatesRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *ListApiTemplatesRequest
	GetTemplateId() *string
	SetTemplateIds(v []*string) *ListApiTemplatesRequest
	GetTemplateIds() []*string
	SetTemplateName(v string) *ListApiTemplatesRequest
	GetTemplateName() *string
}

type ListApiTemplatesRequest struct {
	// 接口名。
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateCluster
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 区域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID。
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 集群模板id。
	//
	// example:
	//
	// at-41b4c6a0fc63****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 集群模板id列表。
	//
	// example:
	//
	// ["AT-****"]
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
	// 集群模板名字。
	//
	// example:
	//
	// DATALAKE模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListApiTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListApiTemplatesRequest) GetApiName() *string {
	return s.ApiName
}

func (s *ListApiTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApiTemplatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApiTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListApiTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *ListApiTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListApiTemplatesRequest) SetApiName(v string) *ListApiTemplatesRequest {
	s.ApiName = &v
	return s
}

func (s *ListApiTemplatesRequest) SetMaxResults(v int32) *ListApiTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiTemplatesRequest) SetNextToken(v string) *ListApiTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiTemplatesRequest) SetRegionId(v string) *ListApiTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListApiTemplatesRequest) SetResourceGroupId(v string) *ListApiTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApiTemplatesRequest) SetTemplateId(v string) *ListApiTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListApiTemplatesRequest) SetTemplateIds(v []*string) *ListApiTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *ListApiTemplatesRequest) SetTemplateName(v string) *ListApiTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListApiTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
