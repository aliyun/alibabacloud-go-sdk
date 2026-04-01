// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceEndpointShrinkRequest
	GetClientToken() *string
	SetDBInstanceEndpointDescription(v string) *ModifyDBInstanceEndpointShrinkRequest
	GetDBInstanceEndpointDescription() *string
	SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointShrinkRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *ModifyDBInstanceEndpointShrinkRequest
	GetDBInstanceId() *string
	SetNodeItemsShrink(v string) *ModifyDBInstanceEndpointShrinkRequest
	GetNodeItemsShrink() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceEndpointShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceEndpointShrinkRequest struct {
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
	NodeItemsShrink *string `json:"NodeItems,omitempty" xml:"NodeItems,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceEndpointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceEndpointShrinkRequest) GetDBInstanceEndpointDescription() *string {
	return s.DBInstanceEndpointDescription
}

func (s *ModifyDBInstanceEndpointShrinkRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *ModifyDBInstanceEndpointShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceEndpointShrinkRequest) GetNodeItemsShrink() *string {
	return s.NodeItemsShrink
}

func (s *ModifyDBInstanceEndpointShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceEndpointShrinkRequest) SetClientToken(v string) *ModifyDBInstanceEndpointShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceEndpointShrinkRequest) SetDBInstanceEndpointDescription(v string) *ModifyDBInstanceEndpointShrinkRequest {
	s.DBInstanceEndpointDescription = &v
	return s
}

func (s *ModifyDBInstanceEndpointShrinkRequest) SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointShrinkRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *ModifyDBInstanceEndpointShrinkRequest) SetDBInstanceId(v string) *ModifyDBInstanceEndpointShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceEndpointShrinkRequest) SetNodeItemsShrink(v string) *ModifyDBInstanceEndpointShrinkRequest {
	s.NodeItemsShrink = &v
	return s
}

func (s *ModifyDBInstanceEndpointShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceEndpointShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceEndpointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
