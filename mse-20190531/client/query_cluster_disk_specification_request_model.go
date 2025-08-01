// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterDiskSpecificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryClusterDiskSpecificationRequest
	GetAcceptLanguage() *string
	SetClusterType(v string) *QueryClusterDiskSpecificationRequest
	GetClusterType() *string
}

type QueryClusterDiskSpecificationRequest struct {
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
	// ZooKeeper
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s QueryClusterDiskSpecificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDiskSpecificationRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryClusterDiskSpecificationRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *QueryClusterDiskSpecificationRequest) SetAcceptLanguage(v string) *QueryClusterDiskSpecificationRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryClusterDiskSpecificationRequest) SetClusterType(v string) *QueryClusterDiskSpecificationRequest {
	s.ClusterType = &v
	return s
}

func (s *QueryClusterDiskSpecificationRequest) Validate() error {
	return dara.Validate(s)
}
