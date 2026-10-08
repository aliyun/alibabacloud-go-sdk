// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReplicationLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteReplicationLinkRequest
	GetDBInstanceId() *string
	SetPromoteToMaster(v bool) *DeleteReplicationLinkRequest
	GetPromoteToMaster() *bool
	SetResourceOwnerId(v int64) *DeleteReplicationLinkRequest
	GetResourceOwnerId() *int64
}

type DeleteReplicationLinkRequest struct {
	// The instance ID of the disaster recovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-2zecuz9tolf******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to delete the data synchronization link between the primary instance and the disaster recovery instance and promote the disaster recovery instance to a primary instance. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	PromoteToMaster *bool  `json:"PromoteToMaster,omitempty" xml:"PromoteToMaster,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteReplicationLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReplicationLinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteReplicationLinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteReplicationLinkRequest) GetPromoteToMaster() *bool {
	return s.PromoteToMaster
}

func (s *DeleteReplicationLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteReplicationLinkRequest) SetDBInstanceId(v string) *DeleteReplicationLinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteReplicationLinkRequest) SetPromoteToMaster(v bool) *DeleteReplicationLinkRequest {
	s.PromoteToMaster = &v
	return s
}

func (s *DeleteReplicationLinkRequest) SetResourceOwnerId(v int64) *DeleteReplicationLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteReplicationLinkRequest) Validate() error {
	return dara.Validate(s)
}
