// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListInstanceCountRequest
	GetAcceptLanguage() *string
	SetClusterType(v string) *ListInstanceCountRequest
	GetClusterType() *string
	SetMseVersion(v string) *ListInstanceCountRequest
	GetMseVersion() *string
	SetRegionId(v string) *ListInstanceCountRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListInstanceCountRequest
	GetRequestPars() *string
}

type ListInstanceCountRequest struct {
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
	// The type of the instance. Valid values: ZooKeeper and Nacos-Ans.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The edition type of the instance. Valid values:
	//
	// 	- `mse_dev`: Developer Edition
	//
	// 	- `mse_pro`: Professional Edition
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The ID of the region where the instance resides. Examples:
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-zhangjiakou: China (Zhangjiakou)
	//
	// 	- cn-shenzhen: China (Shenzhen)
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

func (s ListInstanceCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCountRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceCountRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListInstanceCountRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListInstanceCountRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *ListInstanceCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceCountRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListInstanceCountRequest) SetAcceptLanguage(v string) *ListInstanceCountRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListInstanceCountRequest) SetClusterType(v string) *ListInstanceCountRequest {
	s.ClusterType = &v
	return s
}

func (s *ListInstanceCountRequest) SetMseVersion(v string) *ListInstanceCountRequest {
	s.MseVersion = &v
	return s
}

func (s *ListInstanceCountRequest) SetRegionId(v string) *ListInstanceCountRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceCountRequest) SetRequestPars(v string) *ListInstanceCountRequest {
	s.RequestPars = &v
	return s
}

func (s *ListInstanceCountRequest) Validate() error {
	return dara.Validate(s)
}
