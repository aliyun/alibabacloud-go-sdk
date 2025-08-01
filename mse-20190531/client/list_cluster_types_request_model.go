// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListClusterTypesRequest
	GetAcceptLanguage() *string
	SetConnectType(v string) *ListClusterTypesRequest
	GetConnectType() *string
	SetMseVersion(v string) *ListClusterTypesRequest
	GetMseVersion() *string
	SetRegionId(v string) *ListClusterTypesRequest
	GetRegionId() *string
}

type ListClusterTypesRequest struct {
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
	// The network type. Valid values:
	//
	// 	- slb
	//
	// 	- eni
	//
	// example:
	//
	// slb
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The edition of the MSE instance that you want to purchase.
	//
	// 	- mse_pro: Professional Edition
	//
	// 	- mse_dev: Developer Edition
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The ID of the region in which the instance resides. The region is supported by Microservices Engine (MSE).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClusterTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterTypesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListClusterTypesRequest) GetConnectType() *string {
	return s.ConnectType
}

func (s *ListClusterTypesRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *ListClusterTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClusterTypesRequest) SetAcceptLanguage(v string) *ListClusterTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListClusterTypesRequest) SetConnectType(v string) *ListClusterTypesRequest {
	s.ConnectType = &v
	return s
}

func (s *ListClusterTypesRequest) SetMseVersion(v string) *ListClusterTypesRequest {
	s.MseVersion = &v
	return s
}

func (s *ListClusterTypesRequest) SetRegionId(v string) *ListClusterTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterTypesRequest) Validate() error {
	return dara.Validate(s)
}
