// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListInstanceResourcesRequest
	GetCategory() *string
	SetGroup(v string) *ListInstanceResourcesRequest
	GetGroup() *string
	SetType(v string) *ListInstanceResourcesRequest
	GetType() *string
}

type ListInstanceResourcesRequest struct {
	// The category of the resource. Valid values:
	//
	// - DataManagement
	//
	// - Engine
	//
	// - Monitor
	//
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The group of the resource.
	//
	// If `Category` is `DataManagement`, valid values are:
	//
	// - storage
	//
	// - modelpipeline
	//
	// - datastorage
	//
	// - modeltrain
	//
	// If `Category` is `Engine`, valid values are:
	//
	// - feature
	//
	// - predict
	//
	// - recall
	//
	// - recengine
	//
	// If `Category` is `Monitor`, valid values are:
	//
	// - logs
	//
	// - logsback
	//
	// - coldstart
	//
	// - deploy
	//
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The type of the resource. If specified, only resources of this type are returned.
	//
	// - Hologres
	//
	// - EAS
	//
	// - BE
	//
	// - Rec
	//
	// - Platform
	//
	// - SLS
	//
	// - DataHub
	//
	// - ApsaraMQ for Kafka
	//
	// - Realtime Compute for Apache Flink
	//
	// - ACR
	//
	// - OSS
	//
	// - DataWorks
	//
	// - PAI
	//
	// - MaxCompute
	//
	// - Graph Compute Service
	//
	// - ApsaraDB for Redis
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstanceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListInstanceResourcesRequest) GetGroup() *string {
	return s.Group
}

func (s *ListInstanceResourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListInstanceResourcesRequest) SetCategory(v string) *ListInstanceResourcesRequest {
	s.Category = &v
	return s
}

func (s *ListInstanceResourcesRequest) SetGroup(v string) *ListInstanceResourcesRequest {
	s.Group = &v
	return s
}

func (s *ListInstanceResourcesRequest) SetType(v string) *ListInstanceResourcesRequest {
	s.Type = &v
	return s
}

func (s *ListInstanceResourcesRequest) Validate() error {
	return dara.Validate(s)
}
