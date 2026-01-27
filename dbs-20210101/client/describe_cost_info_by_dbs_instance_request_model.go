// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostInfoByDbsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeCostInfoByDbsInstanceRequest
	GetBackupPlanId() *string
	SetRegionCode(v string) *DescribeCostInfoByDbsInstanceRequest
	GetRegionCode() *string
}

type DescribeCostInfoByDbsInstanceRequest struct {
	// example:
	//
	// dbsr1l3ro21****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeCostInfoByDbsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostInfoByDbsInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostInfoByDbsInstanceRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeCostInfoByDbsInstanceRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeCostInfoByDbsInstanceRequest) SetBackupPlanId(v string) *DescribeCostInfoByDbsInstanceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceRequest) SetRegionCode(v string) *DescribeCostInfoByDbsInstanceRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
