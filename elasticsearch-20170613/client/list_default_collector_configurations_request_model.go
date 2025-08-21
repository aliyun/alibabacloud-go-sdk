// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDefaultCollectorConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResType(v string) *ListDefaultCollectorConfigurationsRequest
	GetResType() *string
	SetResVersion(v string) *ListDefaultCollectorConfigurationsRequest
	GetResVersion() *string
	SetSourceType(v string) *ListDefaultCollectorConfigurationsRequest
	GetSourceType() *string
}

type ListDefaultCollectorConfigurationsRequest struct {
	// The shipper type. Valid values:
	//
	// 	- fileBeat
	//
	// 	- metricBeat
	//
	// 	- heartBeat
	//
	// 	- auditBeat
	//
	// This parameter is required.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// The shipper version. The shipper version varies based on the type of the machine on which the shipper is deployed. Valid values:
	//
	// 	- ECS: 6.8.5_with_community
	//
	// 	- ACK: 6.8.13_with_community
	//
	// This parameter is required.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// The type of the machine on which the shipper is deployed. If you do not configure this parameter, the default configuration files of shippers deployed on all types of machines are returned. Valid values:
	//
	// 	- ECS: ECS instance
	//
	// 	- ACK: ACK cluster
	//
	// example:
	//
	// ECS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListDefaultCollectorConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsRequest) GetResType() *string {
	return s.ResType
}

func (s *ListDefaultCollectorConfigurationsRequest) GetResVersion() *string {
	return s.ResVersion
}

func (s *ListDefaultCollectorConfigurationsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListDefaultCollectorConfigurationsRequest) SetResType(v string) *ListDefaultCollectorConfigurationsRequest {
	s.ResType = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsRequest) SetResVersion(v string) *ListDefaultCollectorConfigurationsRequest {
	s.ResVersion = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsRequest) SetSourceType(v string) *ListDefaultCollectorConfigurationsRequest {
	s.SourceType = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
