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
	// The type of the collector. Valid values:
	//
	// - fileBeat
	//
	// - metricBeat
	//
	// - heartBeat
	//
	// - auditBeat.
	//
	// This parameter is required.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// The version of the collector. The available versions vary based on the type of machine on which the collector is deployed. Valid values:
	//
	// - ECS: 6.8.5_with_community
	//
	// - ACK: 6.8.13_with_community.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// The type of machine on which the collector is deployed. If you do not specify this parameter, all types are returned. Valid values:
	//
	// - ECS: Elastic Compute Service (ECS) instance
	//
	// - ACK: Container Service for Kubernetes (ACK) cluster.
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
