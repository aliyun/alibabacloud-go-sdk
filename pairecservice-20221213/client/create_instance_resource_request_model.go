// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateInstanceResourceRequest
	GetCategory() *string
	SetGroup(v string) *CreateInstanceResourceRequest
	GetGroup() *string
	SetType(v string) *CreateInstanceResourceRequest
	GetType() *string
	SetUri(v string) *CreateInstanceResourceRequest
	GetUri() *string
}

type CreateInstanceResourceRequest struct {
	// The resource category. Valid values:
	//
	// - DataManagement
	//
	// - Engine
	//
	// - Monitor
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The resource type. Valid values:
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
	// - Graph Compute
	//
	// - ApsaraDB for Redis
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The resource URI.
	//
	// This parameter is required.
	//
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateInstanceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceResourceRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceResourceRequest) GetGroup() *string {
	return s.Group
}

func (s *CreateInstanceResourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateInstanceResourceRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateInstanceResourceRequest) SetCategory(v string) *CreateInstanceResourceRequest {
	s.Category = &v
	return s
}

func (s *CreateInstanceResourceRequest) SetGroup(v string) *CreateInstanceResourceRequest {
	s.Group = &v
	return s
}

func (s *CreateInstanceResourceRequest) SetType(v string) *CreateInstanceResourceRequest {
	s.Type = &v
	return s
}

func (s *CreateInstanceResourceRequest) SetUri(v string) *CreateInstanceResourceRequest {
	s.Uri = &v
	return s
}

func (s *CreateInstanceResourceRequest) Validate() error {
	return dara.Validate(s)
}
