// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListClusterVersionsRequest
	GetAcceptLanguage() *string
	SetClusterType(v string) *ListClusterVersionsRequest
	GetClusterType() *string
	SetMseVersion(v string) *ListClusterVersionsRequest
	GetMseVersion() *string
}

type ListClusterVersionsRequest struct {
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
	// The type of the instance. Valid values: ZooKeeper, Nacos-Ans, and Eureka.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- `mse_dev`: Developer Edition.
	//
	// 	- `mse_pro`: Professional Edition. This is the default value.
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
}

func (s ListClusterVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListClusterVersionsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterVersionsRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *ListClusterVersionsRequest) SetAcceptLanguage(v string) *ListClusterVersionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListClusterVersionsRequest) SetClusterType(v string) *ListClusterVersionsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListClusterVersionsRequest) SetMseVersion(v string) *ListClusterVersionsRequest {
	s.MseVersion = &v
	return s
}

func (s *ListClusterVersionsRequest) Validate() error {
	return dara.Validate(s)
}
