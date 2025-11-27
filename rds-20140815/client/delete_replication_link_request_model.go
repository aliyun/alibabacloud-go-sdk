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
	// The ID of the DR instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1trqb4p1xd****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to delete the data synchronization link between the DR instance and the primary instance and promote the DR instance to the primary instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
