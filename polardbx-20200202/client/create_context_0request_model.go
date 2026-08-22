// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContext0Request interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateContext0Request
	GetDBInstanceName() *string
	SetOpenSearchInstanceName(v string) *CreateContext0Request
	GetOpenSearchInstanceName() *string
	SetRegionId(v string) *CreateContext0Request
	GetRegionId() *string
}

type CreateContext0Request struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the PolarDB-X Search instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxs-********
	OpenSearchInstanceName *string `json:"OpenSearchInstanceName,omitempty" xml:"OpenSearchInstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateContext0Request) String() string {
	return dara.Prettify(s)
}

func (s CreateContext0Request) GoString() string {
	return s.String()
}

func (s *CreateContext0Request) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateContext0Request) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *CreateContext0Request) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateContext0Request) SetDBInstanceName(v string) *CreateContext0Request {
	s.DBInstanceName = &v
	return s
}

func (s *CreateContext0Request) SetOpenSearchInstanceName(v string) *CreateContext0Request {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *CreateContext0Request) SetRegionId(v string) *CreateContext0Request {
	s.RegionId = &v
	return s
}

func (s *CreateContext0Request) Validate() error {
	return dara.Validate(s)
}
