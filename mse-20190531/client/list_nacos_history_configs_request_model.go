// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosHistoryConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListNacosHistoryConfigsRequest
	GetAcceptLanguage() *string
	SetDataId(v string) *ListNacosHistoryConfigsRequest
	GetDataId() *string
	SetGroup(v string) *ListNacosHistoryConfigsRequest
	GetGroup() *string
	SetInstanceId(v string) *ListNacosHistoryConfigsRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListNacosHistoryConfigsRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *ListNacosHistoryConfigsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListNacosHistoryConfigsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListNacosHistoryConfigsRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListNacosHistoryConfigsRequest
	GetRequestPars() *string
}

type ListNacosHistoryConfigsRequest struct {
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
	// The ID of the data.
	//
	// example:
	//
	// user-prod.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The name of the configuration group.
	//
	// example:
	//
	// fc-dev-cluster-1
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-2r42e3bk20n
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// fc0f6e40-****-946b-45e3af313707
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
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
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {\\\\"appGroup\\\\":\\\\"aliyun-phecda-service-staging\\\\",\\\\"appName\\\\":\\\\"aliyun-phecda-service\\\\",\\\\"appStage\\\\":\\\\"DAILY\\\\",\\\\"appUnit\\\\":\\\\"center\\\\",\\\\"bucId\\\\":\\\\"250858\\\\",\\\\"bucName\\\\":\\\\"Alibaba Cloud\\\\",\\\\"provider\\\\":\\\\"aliyun\\\\"}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListNacosHistoryConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNacosHistoryConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListNacosHistoryConfigsRequest) GetDataId() *string {
	return s.DataId
}

func (s *ListNacosHistoryConfigsRequest) GetGroup() *string {
	return s.Group
}

func (s *ListNacosHistoryConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNacosHistoryConfigsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNacosHistoryConfigsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListNacosHistoryConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNacosHistoryConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNacosHistoryConfigsRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListNacosHistoryConfigsRequest) SetAcceptLanguage(v string) *ListNacosHistoryConfigsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetDataId(v string) *ListNacosHistoryConfigsRequest {
	s.DataId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetGroup(v string) *ListNacosHistoryConfigsRequest {
	s.Group = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetInstanceId(v string) *ListNacosHistoryConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetNamespaceId(v string) *ListNacosHistoryConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetPageNum(v int32) *ListNacosHistoryConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetPageSize(v int32) *ListNacosHistoryConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetRegionId(v string) *ListNacosHistoryConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetRequestPars(v string) *ListNacosHistoryConfigsRequest {
	s.RequestPars = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) Validate() error {
	return dara.Validate(s)
}
