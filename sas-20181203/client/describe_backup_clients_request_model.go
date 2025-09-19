// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSupportRegionId(v string) *DescribeBackupClientsRequest
	GetSupportRegionId() *string
}

type DescribeBackupClientsRequest struct {
	// The region in which the anti-ransomware feature is supported.
	//
	// > You can call the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the regions in which the anti-ransomware feature is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SupportRegionId *string `json:"SupportRegionId,omitempty" xml:"SupportRegionId,omitempty"`
}

func (s DescribeBackupClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsRequest) GetSupportRegionId() *string {
	return s.SupportRegionId
}

func (s *DescribeBackupClientsRequest) SetSupportRegionId(v string) *DescribeBackupClientsRequest {
	s.SupportRegionId = &v
	return s
}

func (s *DescribeBackupClientsRequest) Validate() error {
	return dara.Validate(s)
}
