// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0Request interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteMem0Request
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteMem0Request
	GetRegionId() *string
}

type DeleteMem0Request struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hz1fds
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMem0Request) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0Request) GoString() string {
	return s.String()
}

func (s *DeleteMem0Request) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteMem0Request) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMem0Request) SetDBInstanceName(v string) *DeleteMem0Request {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteMem0Request) SetRegionId(v string) *DeleteMem0Request {
	s.RegionId = &v
	return s
}

func (s *DeleteMem0Request) Validate() error {
	return dara.Validate(s)
}
