// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationSlotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeReplicationSlotsRequest
	GetDBClusterId() *string
	SetSubGroupName(v string) *DescribeReplicationSlotsRequest
	GetSubGroupName() *string
}

type DescribeReplicationSlotsRequest struct {
	// The ID of the PolarDB cluster to query. This operation is applicable only to PostgreSQL clusters. If you specify the ID of a cluster that uses another database engine, the InvalidDBType.Malformed error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-test123456
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the child node of the distributed instance, which is used to filter replication slots.
	//
	// example:
	//
	// testGroup
	SubGroupName *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
}

func (s DescribeReplicationSlotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationSlotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicationSlotsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeReplicationSlotsRequest) GetSubGroupName() *string {
	return s.SubGroupName
}

func (s *DescribeReplicationSlotsRequest) SetDBClusterId(v string) *DescribeReplicationSlotsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeReplicationSlotsRequest) SetSubGroupName(v string) *DescribeReplicationSlotsRequest {
	s.SubGroupName = &v
	return s
}

func (s *DescribeReplicationSlotsRequest) Validate() error {
	return dara.Validate(s)
}
