// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBTablesRecoveryStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifyDBTablesRecoveryStateRequest
	GetCategory() *string
	SetClusterName(v string) *ModifyDBTablesRecoveryStateRequest
	GetClusterName() *string
	SetInstanceId(v string) *ModifyDBTablesRecoveryStateRequest
	GetInstanceId() *string
	SetRegionCode(v string) *ModifyDBTablesRecoveryStateRequest
	GetRegionCode() *string
	SetRetention(v string) *ModifyDBTablesRecoveryStateRequest
	GetRetention() *string
}

type ModifyDBTablesRecoveryStateRequest struct {
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode  *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	Retention   *string `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s ModifyDBTablesRecoveryStateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBTablesRecoveryStateRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBTablesRecoveryStateRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBTablesRecoveryStateRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyDBTablesRecoveryStateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDBTablesRecoveryStateRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ModifyDBTablesRecoveryStateRequest) GetRetention() *string {
	return s.Retention
}

func (s *ModifyDBTablesRecoveryStateRequest) SetCategory(v string) *ModifyDBTablesRecoveryStateRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetClusterName(v string) *ModifyDBTablesRecoveryStateRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetInstanceId(v string) *ModifyDBTablesRecoveryStateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetRegionCode(v string) *ModifyDBTablesRecoveryStateRequest {
	s.RegionCode = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetRetention(v string) *ModifyDBTablesRecoveryStateRequest {
	s.Retention = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) Validate() error {
	return dara.Validate(s)
}
