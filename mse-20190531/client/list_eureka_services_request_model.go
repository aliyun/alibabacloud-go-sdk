// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEurekaServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListEurekaServicesRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ListEurekaServicesRequest
	GetClusterId() *string
	SetPageNum(v int32) *ListEurekaServicesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListEurekaServicesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListEurekaServicesRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListEurekaServicesRequest
	GetRequestPars() *string
}

type ListEurekaServicesRequest struct {
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
	// This parameter is required.
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
	// The region ID.
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

func (s ListEurekaServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaServicesRequest) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListEurekaServicesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListEurekaServicesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListEurekaServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEurekaServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEurekaServicesRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListEurekaServicesRequest) SetAcceptLanguage(v string) *ListEurekaServicesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListEurekaServicesRequest) SetClusterId(v string) *ListEurekaServicesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEurekaServicesRequest) SetPageNum(v int32) *ListEurekaServicesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEurekaServicesRequest) SetPageSize(v int32) *ListEurekaServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEurekaServicesRequest) SetRegionId(v string) *ListEurekaServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEurekaServicesRequest) SetRequestPars(v string) *ListEurekaServicesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListEurekaServicesRequest) Validate() error {
	return dara.Validate(s)
}
