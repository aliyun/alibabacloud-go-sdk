// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMultiZoneClusterNodeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyMultiZoneClusterNodeTypeRequest
	GetClusterId() *string
	SetCoreInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest
	GetCoreInstanceType() *string
	SetLogInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest
	GetLogInstanceType() *string
	SetMasterInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest
	GetMasterInstanceType() *string
}

type ModifyMultiZoneClusterNodeTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-dj45g7d6rbrd****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// hbase.sn1.2xlarge
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// example:
	//
	// hbase.sn1.8xlarge
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) GetLogInstanceType() *string {
	return s.LogInstanceType
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetClusterId(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetCoreInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetLogInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.LogInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetMasterInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) Validate() error {
	return dara.Validate(s)
}
