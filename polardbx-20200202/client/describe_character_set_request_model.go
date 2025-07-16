// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCharacterSetRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeCharacterSetRequest
	GetRegionId() *string
}

type DescribeCharacterSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCharacterSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCharacterSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCharacterSetRequest) SetDBInstanceName(v string) *DescribeCharacterSetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCharacterSetRequest) SetRegionId(v string) *DescribeCharacterSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCharacterSetRequest) Validate() error {
	return dara.Validate(s)
}
