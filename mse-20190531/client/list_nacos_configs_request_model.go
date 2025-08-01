// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListNacosConfigsRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ListNacosConfigsRequest
	GetAppName() *string
	SetDataId(v string) *ListNacosConfigsRequest
	GetDataId() *string
	SetGroup(v string) *ListNacosConfigsRequest
	GetGroup() *string
	SetInstanceId(v string) *ListNacosConfigsRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListNacosConfigsRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *ListNacosConfigsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListNacosConfigsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListNacosConfigsRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListNacosConfigsRequest
	GetRequestPars() *string
	SetTags(v string) *ListNacosConfigsRequest
	GetTags() *string
}

type ListNacosConfigsRequest struct {
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
	// The name of the application.
	//
	// example:
	//
	// fpx-pds-pns
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the data.
	//
	// example:
	//
	// zeekr-*
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The name of the group. Default value: `default`
	//
	// example:
	//
	// crm
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-7mz2fj****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// fad732a7-ff1a-4f21-8126-4edd4****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the instance resides. The region is supported by Microservices Engine (MSE).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters. The JSON format is supported.
	//
	// example:
	//
	// {\\\\"appGroup\\\\":\\\\"sm_zk_asi_na610\\\\",\\\\"appName\\\\":\\\\"sm-zk\\\\",\\\\"appStage\\\\":\\\\"PUBLISH\\\\",\\\\"appUnit\\\\":\\\\"center\\\\",\\\\"bucId\\\\":\\\\"193554\\\\",\\\\"bucName\\\\":\\\\"Alibaba Mobile Business Group-UC\\\\",\\\\"provider\\\\":\\\\"aliyun\\\\"}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The tags.
	//
	// example:
	//
	// billing
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListNacosConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNacosConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListNacosConfigsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListNacosConfigsRequest) GetDataId() *string {
	return s.DataId
}

func (s *ListNacosConfigsRequest) GetGroup() *string {
	return s.Group
}

func (s *ListNacosConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNacosConfigsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNacosConfigsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListNacosConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNacosConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNacosConfigsRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListNacosConfigsRequest) GetTags() *string {
	return s.Tags
}

func (s *ListNacosConfigsRequest) SetAcceptLanguage(v string) *ListNacosConfigsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListNacosConfigsRequest) SetAppName(v string) *ListNacosConfigsRequest {
	s.AppName = &v
	return s
}

func (s *ListNacosConfigsRequest) SetDataId(v string) *ListNacosConfigsRequest {
	s.DataId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetGroup(v string) *ListNacosConfigsRequest {
	s.Group = &v
	return s
}

func (s *ListNacosConfigsRequest) SetInstanceId(v string) *ListNacosConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetNamespaceId(v string) *ListNacosConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetPageNum(v int32) *ListNacosConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *ListNacosConfigsRequest) SetPageSize(v int32) *ListNacosConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacosConfigsRequest) SetRegionId(v string) *ListNacosConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetRequestPars(v string) *ListNacosConfigsRequest {
	s.RequestPars = &v
	return s
}

func (s *ListNacosConfigsRequest) SetTags(v string) *ListNacosConfigsRequest {
	s.Tags = &v
	return s
}

func (s *ListNacosConfigsRequest) Validate() error {
	return dara.Validate(s)
}
