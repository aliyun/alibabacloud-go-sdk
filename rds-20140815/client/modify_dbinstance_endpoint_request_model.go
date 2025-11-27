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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The user-defined description of the endpoint.
	//
	// example:
	//
	// for readonly business
	DBInstanceEndpointDescription *string `json:"DBInstanceEndpointDescription,omitempty" xml:"DBInstanceEndpointDescription,omitempty"`
	// The endpoint ID of the instance. You can call the DescribeDBInstanceEndpoints operation to query the endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-****
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The information about the endpoint.
	//
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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node ID.
	//
	// You can query the node ID by using the following methods:
	//
	// 	- Log on the ApsaraDB RDS console, go to the instance details page, and then view the ID of the node in the instance topology in the lower part of the instance details page.
	//
	// 	- Call the DescribeDBInstanceAttribute operation to query the node ID.
	//
	// example:
	//
	// rn-xxxx-****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The weight of the node. Read requests are distributed based on the weight.
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
