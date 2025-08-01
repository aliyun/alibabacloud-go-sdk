// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEurekaInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListEurekaInstancesRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ListEurekaInstancesRequest
	GetClusterId() *string
	SetPageNum(v int32) *ListEurekaInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListEurekaInstancesRequest
	GetPageSize() *int32
	SetRequestPars(v string) *ListEurekaInstancesRequest
	GetRequestPars() *string
	SetServiceName(v string) *ListEurekaInstancesRequest
	GetServiceName() *string
}

type ListEurekaInstancesRequest struct {
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
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListEurekaInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListEurekaInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListEurekaInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListEurekaInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEurekaInstancesRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListEurekaInstancesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListEurekaInstancesRequest) SetAcceptLanguage(v string) *ListEurekaInstancesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetClusterId(v string) *ListEurekaInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetPageNum(v int32) *ListEurekaInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetPageSize(v int32) *ListEurekaInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetRequestPars(v string) *ListEurekaInstancesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetServiceName(v string) *ListEurekaInstancesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListEurekaInstancesRequest) Validate() error {
	return dara.Validate(s)
}
