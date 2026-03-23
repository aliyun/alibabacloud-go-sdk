// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceEndpointRequest
	GetClientToken() *string
	SetDBInstanceEndpointDescription(v string) *ModifyDBInstanceEndpointRequest
	GetDBInstanceEndpointDescription() *string
	SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *ModifyDBInstanceEndpointRequest
	GetDBInstanceId() *string
	SetNodeItems(v []*ModifyDBInstanceEndpointRequestNodeItems) *ModifyDBInstanceEndpointRequest
	GetNodeItems() []*ModifyDBInstanceEndpointRequestNodeItems
	SetResourceOwnerId(v int64) *ModifyDBInstanceEndpointRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceEndpointRequest struct {
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceEndpointDescription *string `json:"DBInstanceEndpointDescription,omitempty" xml:"DBInstanceEndpointDescription,omitempty"`
	// This parameter is required.
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// if can be null:
	// true
	NodeItems       []*ModifyDBInstanceEndpointRequestNodeItems `json:"NodeItems,omitempty" xml:"NodeItems,omitempty" type:"Repeated"`
	ResourceOwnerId *int64                                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceEndpointRequest) GetDBInstanceEndpointDescription() *string {
	return s.DBInstanceEndpointDescription
}

func (s *ModifyDBInstanceEndpointRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *ModifyDBInstanceEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceEndpointRequest) GetNodeItems() []*ModifyDBInstanceEndpointRequestNodeItems {
	return s.NodeItems
}

func (s *ModifyDBInstanceEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceEndpointRequest) SetClientToken(v string) *ModifyDBInstanceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequest) SetDBInstanceEndpointDescription(v string) *ModifyDBInstanceEndpointRequest {
	s.DBInstanceEndpointDescription = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequest) SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequest) SetDBInstanceId(v string) *ModifyDBInstanceEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequest) SetNodeItems(v []*ModifyDBInstanceEndpointRequestNodeItems) *ModifyDBInstanceEndpointRequest {
	s.NodeItems = v
	return s
}

func (s *ModifyDBInstanceEndpointRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequest) Validate() error {
	if s.NodeItems != nil {
		for _, item := range s.NodeItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBInstanceEndpointRequestNodeItems struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Weight       *int64  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyDBInstanceEndpointRequestNodeItems) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointRequestNodeItems) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) GetWeight() *int64 {
	return s.Weight
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) SetDBInstanceId(v string) *ModifyDBInstanceEndpointRequestNodeItems {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) SetNodeId(v string) *ModifyDBInstanceEndpointRequestNodeItems {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) SetWeight(v int64) *ModifyDBInstanceEndpointRequestNodeItems {
	s.Weight = &v
	return s
}

func (s *ModifyDBInstanceEndpointRequestNodeItems) Validate() error {
	return dara.Validate(s)
}
