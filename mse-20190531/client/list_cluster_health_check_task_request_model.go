// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterHealthCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListClusterHealthCheckTaskRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *ListClusterHealthCheckTaskRequest
	GetInstanceId() *string
	SetPageNum(v int32) *ListClusterHealthCheckTaskRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListClusterHealthCheckTaskRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListClusterHealthCheckTaskRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListClusterHealthCheckTaskRequest
	GetRequestPars() *string
}

type ListClusterHealthCheckTaskRequest struct {
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
	// The ID of the instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-7pp2o4wfx01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the instance resides.
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
}

func (s ListClusterHealthCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHealthCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *ListClusterHealthCheckTaskRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListClusterHealthCheckTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClusterHealthCheckTaskRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListClusterHealthCheckTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterHealthCheckTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClusterHealthCheckTaskRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListClusterHealthCheckTaskRequest) SetAcceptLanguage(v string) *ListClusterHealthCheckTaskRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListClusterHealthCheckTaskRequest) SetInstanceId(v string) *ListClusterHealthCheckTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListClusterHealthCheckTaskRequest) SetPageNum(v int32) *ListClusterHealthCheckTaskRequest {
	s.PageNum = &v
	return s
}

func (s *ListClusterHealthCheckTaskRequest) SetPageSize(v int32) *ListClusterHealthCheckTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterHealthCheckTaskRequest) SetRegionId(v string) *ListClusterHealthCheckTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterHealthCheckTaskRequest) SetRequestPars(v string) *ListClusterHealthCheckTaskRequest {
	s.RequestPars = &v
	return s
}

func (s *ListClusterHealthCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
