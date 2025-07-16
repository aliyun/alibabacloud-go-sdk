// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeUserEncryptionKeyListRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeUserEncryptionKeyListRequest
	GetRegionId() *string
}

type DescribeUserEncryptionKeyListRequest struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeUserEncryptionKeyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserEncryptionKeyListRequest) SetDBInstanceName(v string) *DescribeUserEncryptionKeyListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) Validate() error {
	return dara.Validate(s)
}
