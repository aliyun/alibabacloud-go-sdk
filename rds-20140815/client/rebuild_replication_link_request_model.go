// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildReplicationLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RebuildReplicationLinkRequest
	GetDBInstanceId() *string
}

type RebuildReplicationLinkRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1trqb4p1xd****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RebuildReplicationLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebuildReplicationLinkRequest) GoString() string {
	return s.String()
}

func (s *RebuildReplicationLinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RebuildReplicationLinkRequest) SetDBInstanceId(v string) *RebuildReplicationLinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RebuildReplicationLinkRequest) Validate() error {
	return dara.Validate(s)
}
