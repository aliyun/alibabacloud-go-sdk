// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMem0SecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeMem0SecurityIpsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeMem0SecurityIpsRequest
	GetRegionId() *string
}

type DescribeMem0SecurityIpsRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMem0SecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0SecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMem0SecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeMem0SecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMem0SecurityIpsRequest) SetDBInstanceName(v string) *DescribeMem0SecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeMem0SecurityIpsRequest) SetRegionId(v string) *DescribeMem0SecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMem0SecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
