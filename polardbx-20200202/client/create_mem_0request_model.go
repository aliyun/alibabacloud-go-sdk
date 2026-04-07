// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMem0Request interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateMem0Request
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateMem0Request
	GetRegionId() *string
}

type CreateMem0Request struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMem0Request) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0Request) GoString() string {
	return s.String()
}

func (s *CreateMem0Request) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateMem0Request) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMem0Request) SetDBInstanceName(v string) *CreateMem0Request {
	s.DBInstanceName = &v
	return s
}

func (s *CreateMem0Request) SetRegionId(v string) *CreateMem0Request {
	s.RegionId = &v
	return s
}

func (s *CreateMem0Request) Validate() error {
	return dara.Validate(s)
}
