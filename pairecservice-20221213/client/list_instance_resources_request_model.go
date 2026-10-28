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
	// The resource category. Valid values:
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
	// The resource group.
	//
	// If the resource category is DataManagement, valid values are:
	//
	// - storage
	//
	// - modelpipeline
	//
	// - datastorage
	//
	// - modeltrain
	//
	// If the resource category is Engine, valid values are:
	//
	// - feature
	//
	// - predict
	//
	// - recall
	//
	// - recengine
	//
	// If the resource category is Monitor, valid values are:
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
	// The resource type filter. Valid values:
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
	// - Kafka
	//
	// - Flinkvvp
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
	// - GraphCompute
	//
	// - Redis
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
