// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchReplicationLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SwitchReplicationLinkRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *SwitchReplicationLinkRequest
	GetOwnerId() *int64
	SetTargetInstanceName(v string) *SwitchReplicationLinkRequest
	GetTargetInstanceName() *string
	SetTargetInstanceRegionId(v string) *SwitchReplicationLinkRequest
	GetTargetInstanceRegionId() *string
}

type SwitchReplicationLinkRequest struct {
	// The ID of the source or primary instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2zecuz9tolf******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the destination DR instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4neh0q12v1******
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// The ID of the region in which the destination DR instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	TargetInstanceRegionId *string `json:"TargetInstanceRegionId,omitempty" xml:"TargetInstanceRegionId,omitempty"`
}

func (s SwitchReplicationLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchReplicationLinkRequest) GoString() string {
	return s.String()
}

func (s *SwitchReplicationLinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchReplicationLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchReplicationLinkRequest) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *SwitchReplicationLinkRequest) GetTargetInstanceRegionId() *string {
	return s.TargetInstanceRegionId
}

func (s *SwitchReplicationLinkRequest) SetDBInstanceId(v string) *SwitchReplicationLinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchReplicationLinkRequest) SetOwnerId(v int64) *SwitchReplicationLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchReplicationLinkRequest) SetTargetInstanceName(v string) *SwitchReplicationLinkRequest {
	s.TargetInstanceName = &v
	return s
}

func (s *SwitchReplicationLinkRequest) SetTargetInstanceRegionId(v string) *SwitchReplicationLinkRequest {
	s.TargetInstanceRegionId = &v
	return s
}

func (s *SwitchReplicationLinkRequest) Validate() error {
	return dara.Validate(s)
}
